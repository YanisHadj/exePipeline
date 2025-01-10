package main

import "testing"

func TestMain(t *testing.T) {
    result := "Hello, Go CI/CD!"
    if result != "Hello, Go CI/CD!" {
        t.Errorf("Expected 'Hello, Go CI/CD!', but got '%s'", result)
    }
}
