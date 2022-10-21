package slices

import (
	"reflect"
	"testing"
)

func TestSliceSum(t *testing.T) {
	t.Run("collection of any size of numbers", func(t *testing.T) {
		var finalSum int
		exampleSlice := []int{1, 2, 4}
		finalSum = SumSlices(exampleSlice)

		want := 1 + 2 + 4
		if finalSum != want {
			t.Errorf("got %d want %d given, %v", finalSum, want, exampleSlice)
		}

	})
}

func TestSliceSumCreator(t *testing.T) {
	exampleSlice1 := []int{1, 2, 4}
	exampleSlice2 := []int{5, 3, 4}
	exampleSlice3 := []int{7, 1, 8}
	got := SliceSumCreator(exampleSlice1, exampleSlice2, exampleSlice3)
	want := []int{1 + 2 + 4, 5 + 3 + 4, 7 + 1 + 8}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
