# nsqSample 

- requirements: installed the nsq suite of proudcts
https://nsq.io/deployment/installing.html

## build instructions
create new nsqDir
cd nsqDir
export GOPATH=$PWD

go get github.com/dfense/nsqSample
cd github.com/dfense/nsqSample
dep ensure -update

cd $GOPATH
go install ./...


*run program to receive usage format*
$ bin/server --help
usage: server --ms=MS [<flags>]

Flags:
      --help     Show context-sensitive help (also try --help-long and --help-man).
  -m, --ms=MS    millisecond to wait between sending messages.
      --count=1  total number of messages to send
      --version  Show application version.

$ bin/server --help
usage: server --ms=MS [<flags>]

Flags:
      --help     Show context-sensitive help (also try --help-long and --help-man).
  -m, --ms=MS    millisecond to wait between sending messages.
      --count=1  total number of messages to send
      --version  Show application version.



