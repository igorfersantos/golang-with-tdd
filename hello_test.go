package hello

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Igor")
	want := "Hello, Igor"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
