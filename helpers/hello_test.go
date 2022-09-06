package helpers

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	fmt.Println("AFTER UNIT TEST")
}

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
