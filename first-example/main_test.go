package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printSomenthing(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)

	go printSomething("rtotest", &wg)

	wg.Wait()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "rtotest") {
		t.Error("Expected to find rtotest, but it is not there")
	}

}
