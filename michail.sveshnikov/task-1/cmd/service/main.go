package main

import "fmt"

func main() {
	var operand1, operand2 int
	_, err := fmt.Scan(&operand1)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	_, err = fmt.Scan(&operand2)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	var operator string
	_, err = fmt.Scan(&operator)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch operator {
	case "+":
		fmt.Println(operand1 + operand2)
	case "-":
		fmt.Println(operand1 - operand2)
	case "*":
		fmt.Println(operand1 * operand2)
	case "/":
		if operand2 == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(operand1 / operand2)
	default:
		fmt.Println("Invalid operation")
	}
}
