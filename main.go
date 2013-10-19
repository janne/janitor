package main

import (
	"flag"
	"fmt"
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
	if len(flag.Args()) > 0 {
		for _, file := range flag.Args() {
			if buffer, err := BufferFromFile(file); err != nil {
				return err
			} else {
				buffers = append(buffers, buffer)
			}
		}
	} else {
		buffers = append(buffers, Buffer{})
	}
	return nil
}
