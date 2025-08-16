package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Alfian")

	if result != "Hello, Alfian!" {
		panic("Result is not 'Hello, Alfian!'")
	}
}

func TestHelloWorldTaka(t *testing.T) {
	result := HelloWorld("Taka")

	if result != "Hello, Taka!" {
		panic("Result is not 'Hello, Taka!'")
	}
}