package slices

import "testing"

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
