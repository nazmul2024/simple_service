package main

import "testing"

func TestRun(t *testing.T) {
	expected := "CI/CD tools"
	got := run()
	if expected != got {
		t.Fatalf("expected %v got %v", expected, got)
	}
}
