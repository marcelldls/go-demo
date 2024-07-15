package main

import "testing"

func Test_isEven(t *testing.T) {
	result := isEven(3)
	if result != "3 is odd" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Foo")
	}
}
