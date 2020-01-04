package minesweeper

import (
	"testing"
)

func TestBoom(t *testing.T) {
	InjectNumber(t, []int{0, 1, 1, 0, 1, 2, 1, 0, 1}, 4)
	InjectNumber(t, []int{0, 1, 1, 0, 0, 0, 0}, 1)
	InjectNumber(t, []int{1, 0, 0, 0, 0, 0, 1, 2, 1, 0}, 3)
	InjectNumber(t, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 0)
}

func InjectNumber(t *testing.T, arr []int, expected int) {
	if Sweeping(arr) != expected {
		t.Errorf("Array %d must Expected %d", arr, expected)
	}
}
