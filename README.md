# KoTH Agent Server

The KoTH Agent Server is an open-source agent to use with CTFd Enterprise, King of the Hill (KoTH) challenges.

It is setup to run alongside the KoTH Challenge Type and its target server/application. 

The agent monitors the target server/application for the current "King of the Hill" and simultaneously listens and responds to HTTP requests.

To learn more about King of The Hill challenges, [check out its documentation right here](https://docs.ctfd.io/docs/custom-challenges/king-of-the-hill).


## File structure

This repository is contains the KoTH Agent Server source code and binaries as well as and example web application to show the interaction between the agent and a target server. 

The `dist` folder contains compiled agents for different operating systems.
The `src` folder contains the source code for the agent.
The `/example` folder contains the files for an example web application used to demonstrate how the agent interacts with other applications, [as shown here](#example-application)

## How to use the agent

You can use the executables found in the `/dist` folder, or you can modify and recompile the agent's source code into an executable file using `go build`. 

You can then run the agent using its available [options](#agent-cli-usage).

For example, running the following code below, with the options indicated, tells the agent to monitor the `owner.txt` file (assuming that the `owner.txt` file is present in the current working directory, and contains the text "example"). 

In addition we specify an API key to prevent unauthorized users from accessing the agent. 

```
./agent -file owner.txt -apikey 123
Listening on 0.0.0.0:31337
Running without encryption
```

We can then access the the `/status` endpoint using cURL or your browser to see the current "owner" of the target application.

```
curl http://localhost:31337/status --header "authorization:123"
{"success":true,"data":{"identifier":"example"}}
```

For more information about the agent's API, you can refer to this article: https://docs.ctfd.io/docs/custom-challenges/king-of-the-hill/redoc


### Agent CLI Usage
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
        provide a command to run when asked for a healthcheck (default "true")
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
        provide a command to run when asked for an owner
  -port string
        port number to listen on (default "31337")
```

### Example Application

To get a sense of how the agent works, an example application is provided in the repository in the `/example` folder. It is a simple web application, built with [Flask](https://flask.palletsprojects.com/), that serves as the agent's target application for it to monitor. It takes in the user's identifier or any text, and writes it to a file called `owner.txt`.

Run the web application together with the agent using `docker-compose up` in the root directory of the repository.

Once the Docker instance is running, you can interact with the agent and example web application.

The web application can be accessed in `http://[server]:5000/`. And the agent can be accessed from two endpoints: `/status` and `/healthcheck`. For example, `http://[server]:31337/status` and `http://[server]:31337/healthcheck`.

Try entering text on the web application's input and submit it.

Then, send a request to the agent via the `/status` endpoint. The agent responds in JSON format, where, the identifier key's value would be the text submitted from the web application.
