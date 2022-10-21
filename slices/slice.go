package slices

func SumSlices(arr []int) int {
	sum := 0
	for _, number := range arr {
		sum += number
	}
	return sum
}

func SliceSumCreator(numbersToSum ...[]int) []int {
	sum := 0
	lengthOfNumbers := len(numbersToSum)
	finalArr := make([]int, lengthOfNumbers)
	for i := 0; i < lengthOfNumbers; i++ {
		sum = SumSlices(numbersToSum[i])
		finalArr[i] = sum
	}
	return finalArr
}
