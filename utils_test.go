package main

import (
	"testing"
)

var (
	files = []string{
		"/tmp/test1.txt",
		"/tmp/test2.txt",
		"/tmp/test3.txt",
	}
)

func TestFileHash(t *testing.T) {
	for _, f := range files {
		_, err := fileHash(f)
		if err != nil {
			t.Error(err)
		}
	}
}
