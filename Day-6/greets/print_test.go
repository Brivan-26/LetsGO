package main 
import (
	"testing"
	"bytes"
)
func TestGreeting(t *testing.T) {
	
	buffer := bytes.Buffer{}
	Greet(&buffer, "Abdessamed")

	got := buffer.String()
	want := "Hello, Abdessamed"

	if got != want {
		t.Errorf("Got %q, Expected: %q", got, want)
	}
}