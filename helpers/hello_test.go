package helpers

import "testing"

func TestHello(t *testing.T) {
	result := Hello("Tiyan")
	if result != "Hello Tiyan" {
		// Unit test failed
		panic("Result is not Hello Tiyan")
	}
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
