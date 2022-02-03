package question1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*

//Problem statement: Test a function that finds factorial of any number

//Test cases
1. find factorial of a number which is positive
2. find factorial of a number which is negative
3. find factorial of large number
4. imperfect number (float)
5. find factorial of zero

*/

//1. test file name should end with _test.go
//2. test function name should start with Test
//3. test function should accept single parameter t *testing.T
func TestGetFactorial(t *testing.T) {
	t.Run("Positive - find factorial of a number which is positive", func(t *testing.T) {
		n := 3
		result, _ := getFactorial(n)
		/*
		   if result != 7 {
		   t.Error("Result is not as expected")
		   }
		*/
		assert.Equal(t, 6, result)
	})

	t.Run("Positive - find factorial of zero", func(t *testing.T) {
		n := 0
		result, _ := getFactorial(n)
		/*
		   if result != 7 {
		   t.Error("Result is not as expected")
		   }
		*/
		assert.Equal(t, 1, result)
	})

	t.Run("Negative - find factorial of a number which is negative", func(t *testing.T) {
		n := -3
		_, err := getFactorial(n)
		assert.NotNil(t, err)
	})

}
