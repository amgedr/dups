package main

import "testing"

func TestExists(t *testing.T) {
	ext := Extensions{}
	ext.Set("jpg")

	if !ext.Exists("/tmp/test.jpg") {
		t.Fatal("jpg check should be true")
	}

	if ext.Exists("/tmp/test.png") {
		t.Fatal("png check should be false")
	}
}
