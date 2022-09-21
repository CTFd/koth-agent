# koth-agent

KoTH Server Agent for use with CTFd Enterprise

API docs: https://docs.ctfd.io/docs/custom-challenges/king-of-the-hill/redoc

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

## Example Application

An example application that shows a basic idea of how to use this agent with a challenge is provided in the example folder. It can be started with `docker-compose up`.

---
# KoTH Agent Server

- What your application does,
- Why you used the technologies you used

The KoTH Agent Server is an open-source agent to use with CTFd Enterprise, King of the Hill (KoTH) challenges.

It is setup to run run alongside the KoTH Challenge Type and its target server/application. 

The agent monitors the target server/application for the current "King of the Hill" and simultaneously listens and responds to HTTP requests.

To learn more about King of the Hill challenges, [check out its documentation right here](https://docs.ctfd.io/docs/custom-challenges/king-of-the-hill).

# How to install and run the agent





# How to use the agent







---
  

by default, the agent is configured to look at the `owner.txt` file. You can view the code by examining the serve.sh file in the example applicaton. 

koth agent on same server and tell it to read the `owner.txt` file. Something like `./agent -file owner.txt`. 

```shell title="/example/serve.sh"
#!/bin/sh
agent -file /opt/app/owner.txt &
python /opt/app/app.py



API docs: https://docs.ctfd.io/docs/custom-challenges/king-of-the-hill/redoc

```