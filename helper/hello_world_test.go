package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Taka")

	assert.Equal(t, "Hello, Taka!", result, "Result is not 'Hello, Taka!'")

	fmt.Println("TestHelloWorldAssert executed!")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Taka")

	require.Equal(t, "Hello, Taka!", result, "Result is not 'Hello, Taka!'")

	fmt.Println("TestHelloWorldRequire executed!")
}