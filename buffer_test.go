package main

import "testing"

func TestBufferLine(t *testing.T) {
	buffer, _ := BufferFromFile("fixtures/fixture.txt")
	if line, err := buffer.Line(1); err != nil || line != "Hello" {
		t.Errorf("Was supposed to be %v but was %v, error: %v", "Hello", line, err)
	}
}

func TestBufferCount(t *testing.T) {
	buffer, _ := BufferFromFile("fixtures/fixture.txt")
	if count := buffer.Count(); count != 4 {
		t.Errorf("Was supposed to be 4 but was %v", count)
	}
}
