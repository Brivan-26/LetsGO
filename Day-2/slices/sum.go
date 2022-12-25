package slices

func Sum(numbers []int) (result int){
	for _, number := range numbers {
		result += number
	}

	return
}

func SumAll(numbersToSum ...[]int) (result []int) {
	for _, slice := range numbersToSum {
		result = append(result, Sum(slice))
	}
	return
}

func SumAllTails(numbersToSumTail ...[]int) (result []int) {
	for _, slice := range numbersToSumTail {
		if(len(slice) == 0) {
			result = append(result, 0)
			continue
		}
		result = append(result, Sum(slice[1:]))
	}
	return
}