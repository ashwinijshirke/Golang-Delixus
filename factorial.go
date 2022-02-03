package question1

import "errors"

func getFactorial(n int) (result int, err error) {
	if n < 0 {
		err = errors.New("cannot find factorial of negative number")
		return
	}

	result = 1
	if n == 0 || n == 1 {
		return
	}
	for i := 2; i <= n; i++ {
		result *= i
	}
	return
}
