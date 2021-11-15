# koth-agent

KoTH Server Agent for use with CTFd Enterprise

```
❯ ./agent -h
Usage of ./agent:
  -apikey string
        API Key to authenticate with
  -certfile string
        SSL certificate file
  -certstring string
        SSL cert as a string
  -cmd string
        command to run when asked for a healthcheck (default "true")
  -file string
        text file to watch for server ownership changes (default "owner.txt")
  -help
        print help text
  -host string
        host address to listen on (default "0.0.0.0")
  -keyfile string
        SSL key file
  -keystring string
        SSL key as a string
  -origin string
        CIDR ranges to allow connections from. IPv4 and IPv6 networks must be specified seperately (default "0.0.0.0/0,::/0")
  -port string
        port number to listen on (default "31337")
```