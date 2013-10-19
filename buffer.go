package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Buffer struct {
	lines []string
}

func BufferFromFile(file string) (Buffer, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return Buffer{}, err
	}
	lines := strings.Split(string(data), "\n")
	return Buffer{lines}, nil
}

func BuffersFromFiles(files []string) ([]Buffer, error) {
	buffers := make([]Buffer, 0)
	if len(files) > 0 {
		for _, file := range files {
			if buffer, err := BufferFromFile(file); err != nil {
				return nil, err
			} else {
				buffers = append(buffers, buffer)
			}
		}
	} else {
		buffers = append(buffers, Buffer{})
	}
	return buffers, nil
}

func (b Buffer) Line(line int) (string, error) {
	if line < 1 || line >= len(b.lines) {
		return "", fmt.Errorf("Line index out of range")
	}
	return b.lines[line-1], nil
}

func (b Buffer) Count() int {
	return len(b.lines)
}
