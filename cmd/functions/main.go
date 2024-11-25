package main

import (
	"fmt" 
	"errors"
)

func main() {
	printMe("Hello World")
	var result, remainder, err = intDivision(10, 5)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else if remainder == 0 {
	// 	fmt.Printf("Result: %v", result)
	// } else {
	// 	fmt.Printf("Result: %v, remainder: %v", result, remainder)
	// }
	switch {
		case err!=nil:
			fmt.Println(err.Error())
		case remainder==0:
			fmt.Printf("Result: %v", result)
		default:
			fmt.Printf("Result: %v, remainder: %v", result, remainder)
	}

	switch remainder {
		case 0:
			fmt.Println("The division was exact")
		case 1, 2:
			fmt.Println("The division was close!")
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator ==0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}

	result := numerator/denominator
	remainder := numerator%denominator
	return result, remainder, err
}
