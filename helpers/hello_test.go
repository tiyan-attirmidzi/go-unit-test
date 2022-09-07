package helpers

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Tests struct {
	Name     string
	Request  string
	Expected string
}

func BenchmarkHelloTable(b *testing.B) {
	data := []Tests{
		{
			Name:     "Tiyan",
			Request:  "Tiyan",
			Expected: "Hello Tiyan",
		},
		{
			Name:     "Attirmidzi",
			Request:  "Attirmidzi",
			Expected: "Hello Attirmidzi",
		},
		{
			Name:     "TiyanAttirmidzi",
			Request:  "Tiyan Attirmidzi",
			Expected: "Hello Tiyan Attirmidzi",
		},
	}

	for _, benchmark := range data {
		b.Run(benchmark.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Hello(benchmark.Name)
			}
		})
	}
}

func BenchmarkHelloSub(b *testing.B) {
	b.Run("Tiyan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Hello("Tiyan")
		}
	})
	b.Run("Attirmidzi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Hello("Attirmidzi")
		}
	})
}

func BenchmarkHelloAttirmidzi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hello("Attirmidzi")
	}
}

func BenchmarkHelloTiyan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hello("Tiyan")
	}
}

func TestHelloWithTable(t *testing.T) {
	tests := []Tests{
		{
			Name:     "Tiyan",
			Request:  "Tiyan",
			Expected: "Hello Tiyan",
		},
		{
			Name:     "Attirmidzi",
			Request:  "Attirmidzi",
			Expected: "Hello Attirmidzi",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := Hello(test.Request)
			require.Equal(t, test.Expected, result, fmt.Sprintf("Result must be '%s'", test.Expected))
		})
	}

}

func TestSubTest(t *testing.T) {
	t.Run("Tiyan", func(t *testing.T) {
		result := Hello("Tiyan")
		require.Equal(t, "Hello Tiyan", result, "Result must be 'Hello Tiyan'")
	})
	t.Run("Attirmidzi", func(t *testing.T) {
		result := Hello("Attirmidzi")
		assert.Equal(t, "Hello Attirmidzi", result, "Result must be 'Hello Attirmidzi'")
	})
}

func TestMain(m *testing.M) {
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	fmt.Println("AFTER UNIT TEST")
}

func TestHelloWithSkipTest(t *testing.T) {
	// skip test
	if runtime.GOOS == "darwin" {
		t.Skip("Can't Run on MacOS")
	}
	result := Hello("Tiyan Attirmidzi")
	assert.Equal(t, "Hello Tiyan Attirmidzi", result, "Result must be 'Hello Tiyan Attirmidzi'")
}

func TestHelloWithAssert(t *testing.T) {
	result := Hello("Attirmidzi")
	assert.Equal(t, "Hello Attirmidzi", result, "Result must be 'Hello Attirmidzi'")
	fmt.Println("Test TestHelloWithAssert is DONE")
}

func TestHelloWithRequire(t *testing.T) {
	result := Hello("Tiyan")
	require.Equal(t, "Hello Tiyan", result, "Result must be 'Hello Tiyan'")
	fmt.Println("Test TestHelloWithRequire is DONE")
}

func TestHelloAttirmidzi(t *testing.T) {
	result := Hello("Attirmidzi")
	if result != "Hello Attirmidzi" {
		// Unit test failed
		// panic("Result is not Hello Attirmidzi")
		t.Fatal("Result must be 'Hello Attirmidzi'")
	}
	fmt.Println("TestHelloAttirmidzi DONE")
}

func TestHelloTiyan(t *testing.T) {
	result := Hello("Tiyan")
	if result != "Hello Tiyan" {
		// Unit test failed
		// panic("Result is not Hello Tiyan")
		t.Fatal("Result must be 'Hello Tiyan'")
	}
	fmt.Println("TestHelloTiyan DONE")
}
