package main

import (
	"errors"
	"fmt"
)

func calculate(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("division by zero is not allowed")
		}
		return a / b, nil
	default:
		return 0, errors.New("invalid operator")
	}
}

func main() {
	var a, b float64
	var op string

	fmt.Print("Enter calculation (e.g. 10 / 2): ")
	fmt.Scan(&a, &op, &b)

	result, err := calculate(a, b, op)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result:", result)
}
