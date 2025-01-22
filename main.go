package main

import (
	"strconv"

	"github.com/alexflint/go-arg"
)

var args struct {
	Input   string `arg:"positional,env:PWD" help:"path input to be used on the mocked API"`
	Listen  string `arg:"-l,--listen" default:"127.0.0.1" help:"address used to listen for HTTP requests"`
	Port    int    `arg:"-p,--port" default:"8080" help:"port used to listen for HTTP requests"`
	Verbose bool   `arg:"-v,--verbose" help:"verbosity level"`
}

func main() {
	arg.MustParse(&args)

	address := args.Listen + ":" + strconv.Itoa(args.Port)

	HandlesRequests(address, args.Input)
}
