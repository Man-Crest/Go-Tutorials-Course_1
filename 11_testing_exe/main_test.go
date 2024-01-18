package main

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	total := Add(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestSkip(t *testing.T) {
	if len(os.Getenv("GOPATH")) != 0 {
		t.Skip("skipping all because gopath is not there")
	}

	fmt.Println("gopath is set :", os.Getenv("GOPATH"))
}

func TestCleanup(t *testing.T) {
	if len(os.Getenv("GOPATH")) != 0 {
		t.Cleanup(func() {
			log.Fatal("error")
		})
	}
}
