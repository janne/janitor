package main

import (
	"testing"
	"github.com/robertkrimen/otto"
)

func TestBuffersFromFiles(t *testing.T) {
	buffers, _ := BuffersFromFiles("fixtures/fixture.txt")
	buffer := buffers[0]
	if buffer.Name != "fixtures/fixture.txt" {
		t.Errorf("Was supposed to be 'fixtures/fixture.txt' but was '%v'", buffer.Name)
	}
}

func TestBufferLine(t *testing.T) {
	buffers, _ := BuffersFromFiles("fixtures/fixture.txt")
	buffer := buffers[0]
	if line, err := buffer.Line(1); err != nil || line != "Hello" {
		t.Errorf("Was supposed to be %v but was %v, error: %v", "Hello", line, err)
	}
}

func TestBuffersSync(t *testing.T) {
	JS := otto.New()
	buffers, _ := BuffersFromFiles("fixtures/fixture.txt")
	buffers.Sync(JS)
	value, _ := JS.Run(`buffers.length`)
	if value, _ := value.ToInteger(); value != 1 {
		t.Errorf("Expected 1, got %v\n", value)
	}
	value, _ = JS.Run(`buffers[0].Name`)
	if value, _ := value.ToString(); value != "fixtures/fixture.txt" {
		t.Errorf("Expected 'fixtures/fixture.txt', got '%v'\n", value)
	}
}
