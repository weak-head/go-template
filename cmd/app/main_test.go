package app

import "testing"

func TestHelloModule(t *testing.T) {
	want := "Ahoy, world!"
	if got := HelloModule(); got != want {
		t.Errorf("HelloModule() = %q, want %q", got, want)
	}
}
