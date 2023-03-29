package main

import "testing"

var tests = []struct {
	x, y, expected int
}{
	{1, 2, 3},
	{2, 3, 5},
	{3, 4, 7},
}

func TestSum(t *testing.T) {
	for _, tt := range tests {
		if sum(tt.x, tt.y) != tt.expected {
			t.Errorf("Expected %d, got %d", tt.expected, sum(tt.x, tt.y))
		}
	}
}
