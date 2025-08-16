package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldAlfian(t *testing.T) {
	result := HelloWorld("Alfian")

	if result != "Hello, Alfian!" {
		t.Fail()
	}

	fmt.Println("TestHelloWorldAlfian, done!")
}

func TestHelloWorldTaka(t *testing.T) {
	result := HelloWorld("Alfian")

	if result != "Hello, Taka!" {
		t.FailNow()
	}

	fmt.Println("TestHelloWorldTaka, done!")
}