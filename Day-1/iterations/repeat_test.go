package iterations

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("d", 4)
	want := "dddd"

	if want != got {
		t.Errorf("Exptected: %q, got: %q", want, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i :=0; i< b.N; i++ {
		Repeat("b", 4)
	}
}

func ExampleRepeat() {
	got := Repeat("d", 4)
	fmt.Println(got)

	// Output: dddd
}