package main

import "testing"

func TestLoadEnv(t *testing.T) {
	err := LoadEnv()
	if err != nil {
		t.Errorf("Failed to load environment variables")
	}
}
