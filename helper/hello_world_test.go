package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("BEFORE unit test")
	m.Run()
	fmt.Println("AFTER unit test")
}

func TestHelloWorldAlfian(t *testing.T) {
	result := HelloWorld("Alfian")

	if result != "Hello, Alfian!" {
		t.Error("Result is not 'Hello, Alfian!'")
	}

	fmt.Println("TestHelloWorldAlfian, done!")
}

func TestHelloWorldTaka(t *testing.T) {
	result := HelloWorld("Taka")

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

func TestHelloWorldSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Unit test doesn't work on MacOS.")
	}
	
	result := HelloWorld("Taka")
	require.Equal(t, "Hello, Taka!", result, "Result is not 'Hello, Taka!'")

	fmt.Println("TestHelloWorldRequire executed!")
}

func TestMySubTest(t *testing.T) {
	// go test ./... -run=TestMySubTest/Taka
	// go test ./... -run=/Taka
	t.Run("Taka", func(t *testing.T) {
		result := HelloWorld("Taka")
		require.Equal(t, "Hello, Taka!", result, "Result is not 'Hello, Taka!'")
	})
	// go test ./... -run=TestMySubTest/Alfian
	t.Run("Alfian", func(t *testing.T) {
		result := HelloWorld("Alfian")
		require.Equal(t, "Hello, Alfian!", result, "Result is not 'Hello, Alfian!'")
	})
}