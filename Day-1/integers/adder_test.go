package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	got := Add(1, 3)
	want :=4
	
	if got != want {
		t.Errorf("Expected: %d, got: %d", want, got)
	}
}

func ExampleAdd() {
	sum := Add (1, 3)
	fmt.Println(sum)
	// Output: 4
}