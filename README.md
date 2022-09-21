# KoTH Agent Server

The KoTH Agent Server is an open-source agent to use with CTFd Enterprise, King of the Hill (KoTH) challenges.

It is setup to run alongside the KoTH Challenge Type and its target server/application. 

The agent monitors the target server/application for the current "King of the Hill" and simultaneously listens and responds to HTTP requests.

To learn more about King of The Hill challenges, [check out its documentation right here](https://docs.ctfd.io/docs/custom-challenges/king-of-the-hill).


## File structure

This repository is setup to run an example web application along with the KoTH Agent Server. Essentially, the only file reponsible for running the agent are its executables located in the `/dist` folder, which are all compiled for different OS's (operating systems).

The agent is built with the [Go](https://go.dev/) programming language, and its source code is located in `/src/main.go`.

The `/example` folder contains the files for the example web application used to demonstrate how the agent interacts with it, [as shown here](#example-application)

## How to use the agent

You can use the executables found in the `/dist` folder, or you can modify and recompile the agent's source code into an executable file using `go build`. This compiles it into a file named, `src.exe`. After such, rename the recompiled file to `agent.exe`.

You can then run the agent using its available options below:

```
❯ ./agent -h
Usage of ./agent:
  -apikey string
        API Key to authenticate with
  -certfile string
        SSL certificate file
  -certstring string
        SSL cert as a string
  -file string
        text file to watch for server ownership changes (default "owner.txt")
  -health-cmd string
        command to run when asked for a healthcheck (default "true")
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
  -owner-cmd string
        command to run when asked for an owner
  -port string
        port number to listen on (default "31337")
```

### Example Application

To get a sense of how the agent works, an example application is provided in the repository in the `/example` folder. 

This simple web application, built with [Flask](https://flask.palletsprojects.com/en/2.2.x/), serves as the agent's target application for it to monitor. It is a website that takes in the user's identifier or any text, and writes it to a file called `owner.txt`.

It can be started with `docker-compose up`.

Once the Docker instance is running, you can interact with the agent and example web application.

The web application can be accessed in `http://<server ip address>:5000/`. And the agent can be accessed from two endpoints: `/status` and `/healthcheck`. For example, `http://<server ip address>:31337/status` and `http://<server ip address>:31337/healthcheck`.

For more information about the agent's API, you can refer to this article: https://docs.ctfd.io/docs/custom-challenges/king-of-the-hill/redoc

Try entering a text on the web application's input and submit it.

Then, send a request to the agent via the `/status` endpoint. The agent responds in JSON format, where, the identifier key's value would be the text submitted from the web application.