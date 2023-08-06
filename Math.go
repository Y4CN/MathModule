package Math

import "errors"

func Subtraction(numbers ...int) int {
	var x int
	for _, value := range numbers {
		x = x - value
	}
	return x
}

func Sum(numbers ...int) int {
	var x int
	for _, value := range numbers {
		x = x + value
	}
	return x
}

func Multiplication(numbers ...int) (int, error) {

	var x int
	for _, value := range numbers {
		if value == 0 {
			return 0, errors.New("zero value")
		}
		x = x * value
	}
	return x, nil
}

func Division(numbers ...int) (int, error) {

	var x int
	for _, value := range numbers {
		if value == 0 {
			return 0, errors.New("zero value")
		}
		x = x * value
	}
	return x, nil
}
