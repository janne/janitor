package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"github.com/robertkrimen/otto"
)

type Buffer struct {
	Num int
	Name string
	Lines []string
}

type Buffers []*Buffer

func BuffersFromFiles(files ...string) (Buffers, error) {
	buffers := Buffers{}
	for num, file := range files {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, err
		}
		lines := strings.Split(string(data), "\n")
		buffers = append(buffers, &Buffer{num, file, lines})
	}
	return buffers, nil
}

func (b Buffer) Line(line int) (string, error) {
	if line < 1 || line >= len(b.Lines) {
		return "", fmt.Errorf("Line index out of range")
	}
	return b.Lines[line-1], nil
}

func (b Buffers) Sync(o *otto.Otto) error {
	if v, err := o.ToValue(b); err != nil {
		return err
	} else {
		o.Set("buffers", v)
	}
	return nil
}
