package main

import "math/big"

func factorial(n int64) *big.Int {
	fact := new(big.Int)
	fact.MulRange(1, n)
	return fact
}

/*func main(){
	fmt.Println(factorial(5000))
}*/
