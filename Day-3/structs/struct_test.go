package structs

import "testing"

func TestPerimeter(t *testing.T) {

	got := Perimeter(Rectangle{10.0, 10.0})
	want := 40.0
	if want != got {
		t.Errorf("Got : %.2f, Expected: %.2f", got, want)
	}
}
