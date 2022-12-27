package methods_interfaces

import "testing"

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		
		if got != want {
			t.Errorf("Got: %g, Exptected: %g", got, want)
		}
	}

	t.Run("rectangles: ", func(t *testing.T) {
		rectangle := Rectangle{13.0, 2.0}
		checkArea(t, rectangle, 26.0)
	})
	t.Run("Circles: ", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}



// Table test driven

func TestArea2(t *testing.T) {
	areaTests := []struct {
		name string
		shape Shape
		want float64
	}{
		{name: "Rectangle",shape: Rectangle{13.0, 2.0}, want: 26.0},
		{name: "Circle",shape: Circle{10}, want: 314.1592653589793},
		{name: "Triangle",shape: Triangle{6.0, 6.0}, want: 18.0},
	}

	for _, testCase := range areaTests {
		t.Run(testCase.name, func(t *testing.T) {
			got := testCase.shape.Area()
			if got != testCase.want {
				t.Errorf("Got: %g, Exptected: %g", got, testCase.want)
			}
		})
	}
}
