package main

import "fmt"

//example of if statment
func main() {
	var passWord int = 100

	if passWord == 100 {
		fmt.Println("OK!")
	}
}

//example of if-else statment
func main1() {
	fmt.Print("enter your name : ")
	var name string
	fmt.Scanln(&name)

	if name == "ashu" {
		fmt.Println("welcome", name)
	} else {
		fmt.Println("wrong name")
	}
}

//example of else-if statment
func main2() {
	fmt.Print("enter your percentage : ")
	var p int
	fmt.Scanln(&p)

	if p > 80 {
		fmt.Println("you got 1st class with distinction and A+ grade")
	} else if p > 60 && p < 80 {
		fmt.Println("you got 1st class and A grade ")
	} else if p > 45 && p < 60 {
		fmt.Println("you got 2nd class and B grade")
	} else if p < 45 {
		fmt.Println("you are fail")
	} else {
		fmt.Println("This is your academic progress")
	}

}

//example of nested if-else
func main3() {
	var percentage int = 90
	var rank string = "first"

	if percentage == 90 {
		if rank == "first" {
			fmt.Println("you got admission to IIT")
		} else {
			fmt.Println("you not eligible")
		}
	} else {
		fmt.Println("Try next time")
	}

}
