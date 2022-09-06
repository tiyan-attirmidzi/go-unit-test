package helpers

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHello(t *testing.T) {
	result := Hello("Tiyan")
	if result != "Hello Tiyan" {
		// Unit test failed
		panic("Result is not Hello Tiyan")
	}
}

func TestHelloWithAssert(t *testing.T) {
	// skip test
	if runtime.GOOS == "darwin" {
		t.Skip("Unit Test tidak dapat berjalan pada MacOS")
	}
	result := Hello("Attirmidzi")
	assert.Equal(t, "Hello Attirmidzi", result, "Result must be 'Hello Attirmidzi'")
	fmt.Println("Test TestHelloWithAssert is DONE")
}

func TestHelloWithRequire(t *testing.T) {
	result := Hello("Tiyan")
	require.Equal(t, "Hello Tiyan", result, "Result must be 'Hello Tiyan'")
	fmt.Println("Test TestHelloWithRequire is DONE")
}

func TestHelloWithSkipTest(t *testing.T) {
	// skip test
	if runtime.GOOS == "darwin" {
		t.Skip("Can't Run on MacOS")
	}
	result := Hello("Attirmidzi")
	assert.Equal(t, "Hello Attirmidzi", result, "Result must be 'Hello Attirmidzi'")
}

// RUN IN TERMINAL

// run test all on package
// go test

// `-v` for see detail func test
// go test -v

// run test spesific func
// go test -run funcTestName

// run test all package (from root dir)
// go test ./...
