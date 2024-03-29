package arrayslices

// here numbers is actually a slice
// arrays have their size embedded in the type
// [5]int is a different type from [4]int
func Sum(numbers []int) int {
	result := 0

	// _ is the blank identifier
	// helps us avoid the `unused variable` errors
	// index, value
	for _, val := range numbers {
		result += val
	}

	return result
}

func SumAll(args ...[]int) []int {
	result := []int{}

	for _, list := range args {
		result = append(result, Sum(list))
	}

	return result
}

// sum of all the elements except the first one (head)
func SumTail(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	tail := numbers[1:]

	return Sum(tail)
}
