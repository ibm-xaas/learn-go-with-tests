package main

//Sum ... Sum(numbers [5]int) int {
func Sum(numbers []int) int {
	sum := 0

	//for i := 0; i < 5; i++ {
	//	sum += numbers[i]
	//}

	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll ...
func SumAll(numbersToSum ...[]int) []int {
	/*
		lengthOfNumbers := len(numbersToSum)
		sums = make([]int, lengthOfNumbers)

		for i, numbers := range numbersToSum {
			sums[i] = Sum(numbers)
		}
	*/
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// SumAllTails ...
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) > 0 {
			sums = append(sums, Sum(numbers[1:]))
		} else {
			sums = append(sums, 0)
		}
	}

	return sums
}

func main() {

}
