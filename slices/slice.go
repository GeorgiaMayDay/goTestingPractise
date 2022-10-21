package slices

func SumSlices(arr []int) int {
	sum := 0
	for _, number := range arr {
		sum += number
	}
	return sum
}

func SliceSumCreator(numbersToSum ...[]int) []int {
	var finalArr []int
	for _, arr := range numbersToSum {
		finalArr = append(finalArr, SumSlices(arr))
	}
	return finalArr
}
