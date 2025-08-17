package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TableTest struct {
	name string
	request string
	want string
}

// go test -v -bench=.
// go test -v -run=TestNothing -bench=.
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Taka")
	}
}

func BenchmarkHelloWorldAlfianTaka(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Alfian Taka")
	}
}

// go test -v -run=TestNothing -bench=BenchmarkHelloWorldSub/Alfian\ Taka
func BenchmarkHelloWorldSub(b *testing.B) {
	b.Run("Taka", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Taka")
		}
	})

	b.Run("Alfian Taka", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Alfian Taka")
		}
	})
}

func TestMain(m *testing.M) {
	fmt.Println("BEFORE unit test")
	m.Run()
	fmt.Println("AFTER unit test")
}

func TestTableHelloWorld(t *testing.T) {
	var tests = []TableTest{
		{"Alfian", "Alfian", "Hello, Alfian!"},
		{"Taka", "Taka", "Hello, Taka!"},
		{"Moe", "Moe", "Hello, Moe!"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.want, result)
		})
	}
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