package main

import "fmt"

func main() {

	var number1, number2 float64

	var operator string

	fmt.Printf("Enter a 1st number:")
	fmt.Scanln(&number1)

	fmt.Printf("Enter a 2nd number:")
	fmt.Scanln(&number2)

	fmt.Printf("Enter a operator (+ - * /)")
	fmt.Scanln(&operator)

	switch operator {

	case "+":
		fmt.Printf("%f %s %f=%f", number1, operator, number2, number1+number2)

	case "-":
		fmt.Printf("%f %s %f=%f", number1, operator, number2, number1-number2)

	case "*":
		fmt.Printf("%f %s %f=%f", number1, operator, number2, number1*number2)

	case "/":
		if number2 == 0.0 {
			fmt.Printf("Error,divisible by 0")
		} else {
			fmt.Printf("%f %s %f=%f", number1, operator, number2, number1/number2)
		}
	default:
		fmt.Println("Invalid operator")

	}
}
