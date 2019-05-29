package math

import "testing"

func TestSumArray(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	want := 10
	got := Sum(nums...)
	if got != want {
		t.Errorf("Sum() = %q, want %q", got, want)
	}
}

func TestSumVariadic(t *testing.T) {
	want := 10
	got := Sum(1, 2, 3, 4)
	if got != want {
		t.Errorf("Sum() = %q, want %q", got, want)
	}
}
