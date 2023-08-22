package helloworld

import "testing"

func TestHelloWorld(t *testing.T) {
	wanted := "Hello, World!"
	result := HelloWorld()

	if result != wanted {
		t.Errorf("HelloWorld() = %s; must be %s", result, wanted)
	}
}
