package main

import (
	"flag"
	"github.com/robertkrimen/otto"
)

var buffers Buffers

var JS = otto.New()

func main() {
	if err := parseArgs(); err != nil {
		panic(err)
	}
	JS.Run(`
	console.log("Welcome to Janitor!")
	console.log("We have " + buffers.length + " buffers")
	`)
}

func parseArgs() error {
	flag.Parse()
	var err error
	if buffers, err = BuffersFromFiles(flag.Args()...); err != nil {
		return err
	}
	buffers.Sync(JS)
	return nil
}
