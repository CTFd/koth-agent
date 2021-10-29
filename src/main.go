package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"

	"github.com/gorilla/handlers"
)

// StatusCheckData : string used to represent the current owner of the server
type StatusCheckData struct {
	Identifier string `json:"identifier"`
}

// StatusCheckResponse : structure used for JSON response on /status
type StatusCheckResponse struct {
	Success bool            `json:"success"`
	Data    StatusCheckData `json:"data"`
}

// HealthCheckData : structure used for JSON responses to represent health check data
type HealthCheckData struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
	Status int    `json:"status"`
}

// HealthCheckResponse : structure used for JSON responses for /healthcheck
type HealthCheckResponse struct {
	Success bool            `json:"success"`
	Data    HealthCheckData `json:"data"`
}

// AllowedOrigins : List of CIDR ranges that are allowed to access the agent
var AllowedOrigins []net.IPNet

// AuthenticationToken : Authentication token used to access the agent if specified
var AuthenticationToken string

// TargetFile : File that the agent will read for the server owner
var TargetFile string

// HealthCheckCommand : Command that the agent will run when the /healthcheck endpoint is hit
var HealthCheckCommand string

var host string
var port string
var file string
var cmd string
var origin string
var keystring string
var certstring string
var keyfile string
var certfile string
var apikey string

func runCommand(name string, args ...string) (stdout string, stderr string, exitCode int) {
	// https://stackoverflow.com/a/40770011
	var outbuf, errbuf bytes.Buffer
	cmd := exec.Command(name, args...)
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	err := cmd.Run()
	stdout = outbuf.String()
	stderr = errbuf.String()

	if err != nil {
		// try to get the exit code
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			exitCode = ws.ExitStatus()
		} else {
			exitCode = 1
			if stderr == "" {
				stderr = err.Error()
			}
		}
	} else {
		// success, exitCode should be 0 if go is ok
		ws := cmd.ProcessState.Sys().(syscall.WaitStatus)
		exitCode = ws.ExitStatus()
	}
	return
}

func authorizeRequestKey(req *http.Request) bool {
	if len(AuthenticationToken) > 0 {
		header := req.Header.Get("Authorization")

		if strings.Contains(header, " ") {
			auth := strings.SplitN(header, " ", 2)
			if len(auth) != 2 || auth[0] != "Token" {
				return false
			}
			if auth[1] != AuthenticationToken {
				return false
			}
		} else {
			auth := header
			if auth != AuthenticationToken {
				return false
			}
		}
	}
	return true
}

func authorizeRequestIP(req *http.Request) bool {
	host, _, _ := net.SplitHostPort(req.RemoteAddr)
	ip := net.ParseIP(host)
	valid := false
	for _, origin := range AllowedOrigins {
		if origin.Contains(ip) == true {
			valid = true
		}
	}
	return valid
}

// status godoc
// @Summary Show the current owner of the server that the agent is currently running on
// @Security AuthenticationToken
// @Accept  json
// @Produce  json
// @Success 200 {object} StatusCheckResponse
// @Success 401 "Request did not provide a valid authentication token"
// @Success 403 "Request did not come from an IP within the whitelisted IP ranges"
// @Router /status [get]
func status(w http.ResponseWriter, req *http.Request) {
	validIP := authorizeRequestIP(req)
	validKey := authorizeRequestKey(req)

	if validKey == false {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if validIP == false {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	f, _ := ioutil.ReadFile(TargetFile)
	identifier := strings.TrimSpace(string(f))
	resp := StatusCheckResponse{Success: true, Data: StatusCheckData{Identifier: identifier}}
	j, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

// healthcheck godoc
// @Summary Show the current status of the server by running the stored command
// @Security AuthenticationToken
// @Accept  json
// @Produce  json
// @Success 200 {object} HealthCheckResponse
// @Success 401 "Request did not provide a valid authentication token"
// @Success 403 "Request did not come from an IP within the whitelisted IP ranges"
// @Router /healthcheck [get]
func healthcheck(w http.ResponseWriter, req *http.Request) {
	validIP := authorizeRequestIP(req)
	validKey := authorizeRequestKey(req)

	if validKey == false {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if validIP == false {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	stdout, stderr, exitCode := runCommand(HealthCheckCommand)
	success := (len(stderr) == 0) && (exitCode == 0)

	resp := HealthCheckResponse{Success: success, Data: HealthCheckData{Stdout: stdout, Stderr: stderr, Status: exitCode}}
	j, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

// @title CTFd King of the Hill Agent
// @version 1.0
// @description This agent implements a small HTTP interface for scoring servers (i.e. CTFd Enterprise) to poll during a King of the Hill CTF.
// @license.name Apache 2.0

// @securityDefinitions.apikey AuthenticationToken
// @in header
// @name Authorization
func main() {
	if host == "" {
		flag.StringVar(&host, "host", "0.0.0.0", "host address to listen on")
	}

	if port == "" {
		flag.StringVar(&port, "port", "31337", "port number to listen on")
	}

	if file == "" {
		flag.StringVar(&file, "file", "owner.txt", "text file to watch for server ownership changes")
	}

	if cmd == "" {
		flag.StringVar(&cmd, "cmd", "true", "command to run when asked for a healthcheck")
	}

	if origin == "" {
		flag.StringVar(&origin, "origin", "0.0.0.0/0,::/0", "CIDR ranges to allow connections from. IPv4 and IPv6 networks must be specified seperately")
	}

	if keystring == "" {
		flag.StringVar(&keystring, "keystring", "", "SSL key as a string")
	}

	if certstring == "" {
		flag.StringVar(&certstring, "certstring", "", "SSL cert as a string")
	}

	if keyfile == "" {
		flag.StringVar(&keyfile, "keyfile", "", "SSL key file")
	}

	if certfile == "" {
		flag.StringVar(&certfile, "certfile", "", "SSL certificate file")
	}

	if apikey == "" {
		flag.StringVar(&apikey, "apikey", "", "API Key to authenticate with")
	}

	help := flag.Bool("help", false, "print help text")
	flag.Parse()

	var Help = *help
	if Help {
		flag.PrintDefaults()
		return
	}

	origins := strings.Split(origin, ",")
	for _, o := range origins {
		_, ipnet, _ := net.ParseCIDR(o)
		AllowedOrigins = append(AllowedOrigins, *ipnet)
	}

	AuthenticationToken = apikey
	TargetFile = file
	HealthCheckCommand = cmd
	rawPort, _ := strconv.Atoi(port)

	addr := fmt.Sprintf("%s:%d", host, rawPort)
	fmt.Println("Listening on " + addr)

	http.HandleFunc("/status", status)
	http.HandleFunc("/healthcheck", healthcheck)

	if len(keyfile) > 0 && len(certfile) > 0 {
		fmt.Println("Running with encryption certificates from filesystem")
		http.ListenAndServeTLS(addr, certfile, keyfile, handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
	} else if len(keystring) > 0 && len(certstring) > 0 {
		fmt.Println("Running with pinned encryption certificates")
		cert, err := tls.X509KeyPair([]byte(certstring), []byte(keystring))
		if err != nil {
			log.Fatal(err)
		}

		tlsConfig := &tls.Config{
			Certificates: []tls.Certificate{cert},
		}
		server := http.Server{
			Addr:      addr,
			Handler:   handlers.LoggingHandler(os.Stdout, http.DefaultServeMux),
			TLSConfig: tlsConfig,
		}
		server.ListenAndServeTLS("", "")
	} else {
		fmt.Println("Running without encryption")
		http.ListenAndServe(addr, handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
	}
}
