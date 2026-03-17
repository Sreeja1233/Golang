package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{2, 3, 5},
		{6, 7, 13},
	}
	for _, tt := range tests {
		result := add(tt.a, tt.b)

		if result != tt.expected {
			t.Errorf("Expected %d but got %d", tt.expected, result)
		}
	}
}
