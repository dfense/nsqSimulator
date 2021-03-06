# nsqSimulator
Provides the following independent executable programs below. It will always connect to NSQ on localhost:4150 TCP.

### A Producer (server)
creates a topic and populates it with desired number of messages/speed

### A Consumer (client)
subscribes to a topic and reads messages from it, at desired rate and total messages to read

## Pre-requisites
- installed the nsq suite of proudcts https://nsq.io/deployment/installing.html
- golang (version 1.12+ is required)

```
## build instructions  
<create new nsqUtil dir somewhere>  
<cd nsqUtil> 
# export GOBIN=$PWD  

# git clone https://github.com/dfense/nsqSimulator.git $GOPATH/src/github.com/dfense/nsqSimulator
# cd src/github.com/dfense/nsqSimulator
# go install ./...
```

*run program to receive usage format*
```
$ bin/server --help
usage: server --topic=TOPIC --ms=MS [<flags>]

Flags:
      --help         Show context-sensitive help (also try --help-long and --help-man).
  -t, --topic=TOPIC  name of topic.
  -m, --ms=MS        millisecond to wait between sending messages.
      --count=1      total number of messages to send
      --version      Show application version.

$ bin/client --help
usage: client --channel=CHANNEL --topic=TOPIC [<flags>]

Flags:
      --help             Show context-sensitive help (also try --help-long and --help-man).
  -c, --channel=CHANNEL  name of channel.
  -t, --topic=TOPIC      name of topic.
      --ttlPackets=1     total number of messages to read
      --timeBetween=1    number of millis between message reads
      --stayConnected    stay connected after ttlPacketsRead
      --version          Show application version.

```

## Update
09/19/2019 - added python script to clean up big queues (depth > 10)

