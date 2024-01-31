# Go-ssip

## Installation
Clone this repo with `git clone git@github.com:vposloncec/go-ssip.git`

* [**Go**](https://golang.org/cmd/go/) cli tool should be installed (run `go` to check).

  **If you have `go` you can skip this part.**

  The tool can be found and installed using your operating system's package manager or just
  position yourself inside this repo and run:
  ```shell
  make install-go
  # refresh your shell env variables with:
  source ~/.(your_shell_rc_file)

  ```
  This script should install the tool and set needed env variables for all go-based tools to work.

#### Install using
```shell
make
```

## Build without install
```shell
make build
```
this will create an executable in `./bin` folder


## Documentation
`go-ssip help`
```
Go-ssip creates a number of nodes that represent IoT devices
connected together in a P2P network. The nodes can have various attributes
they may or may not be public to other nodes. Propagation is done using gossip algorithm
with various parameters.

Usage:
  go-ssip [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  run         Start simulation

Flags:
  -l, --adjacency         Print adjacency list
      --config string     config file (default is $HOME/.go-ssip.yaml)
  -c, --connections int   number of connections each node has to others (default 3)
  -h, --help              help for go-ssip
  -n, --nodes int         number of nodes to spawn (default 10)
  -v, --verbose           Debug output (verbose)
      --version           version for go-ssip

Use "go-ssip [command] --help" for more information about a command.
```

## Running tests

Running performance test. This increases number of Nodes and connections in
simulation until simulation lasts for at least 10 seconds

-v is included for a more verbose test, exclude otherwise.

```
go test -bench=. -run=Bench -test.benchtime=10s -v ./...
```

Running all other tests:

```
go test -v ./...
```

Add `-count=1` flag to disable test caching


