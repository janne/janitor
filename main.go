package main

import (
	"flag"
)

var buffers []Buffer

func init() {
}

func main() {
	if err := parseArgs(); err != nil {
		panic(err)
	}
}

func parseArgs() error {
	flag.Parse()
	if bufs, err := BuffersFromFiles(flag.Args()); err != nil {
		return err
	} else {
		buffers = append(buffers, bufs...)
	}
	return nil
}
