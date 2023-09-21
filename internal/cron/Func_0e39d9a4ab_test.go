package main

import "testing"

func AddNumbers(a, b int) int {
    return a + b
}

func TestAddNumbers(t *testing.T) {
    result := AddNumbers(2, 3)
    if result != 5 {
        t.Errorf("AddNumbers was incorrect, got: %d, want: %d.", result, 5)
    }
}
