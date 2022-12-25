package slices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	got := Sum(nums)
	want := 15
	
	if want != got {
		t.Errorf("Got: %d, Exptected: %d, from array: %v", got, want, nums)
	}
}

func ExampleSum() {
	sum := Sum([]int{2,3})
	fmt.Println(sum)
	
	// Output: 5
}

func TestSumAll(t *testing.T) {
	slice1 := []int{2, 3}
	slice2 := []int{5, 5}

	got := SumAll(slice1, slice2)
	want := []int{5, 10}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got: %v, Expected: %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got: %v, Expected: %v", got, want)
		}
	}

	t.Run("Make sum of tails that contain non empty slices", func(t *testing.T) {
		got := SumAllTails([]int{2,2}, []int{2,3}, []int{1,1,1,5,7})
		want := []int{2,3,14}
	
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got: %v, Expected: %v", got, want)
		}
	})
	t.Run("Make sum of tails that contain empty slices safely", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{2,3}, []int{1,1,1,5,7})
		want := []int{0,3,14}
		
		checkSums(t, got, want)
		
	})
}