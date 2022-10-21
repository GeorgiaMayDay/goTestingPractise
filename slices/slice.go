package slices

func SumSlices(arr []int) int {
	sum := 0
	for _, number := range arr {
		sum += number
	}
	return sum
}
