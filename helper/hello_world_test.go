package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldAlfian(t *testing.T) {
	result := HelloWorld("Alfian")

	if result != "Hello, Alfian!" {
		t.Error("Result is not 'Hello, Alfian!'")
	}

	fmt.Println("TestHelloWorldAlfian, done!")
}

func TestHelloWorldTaka(t *testing.T) {
	result := HelloWorld("Alfian")

	if result != "Hello, Taka!" {
		t.Fatal("Result is not 'Hello, Taka!'")
	}

	fmt.Println("TestHelloWorldTaka, done!")
}