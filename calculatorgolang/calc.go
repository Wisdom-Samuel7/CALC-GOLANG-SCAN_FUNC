package main

import (
	"fmt"
	"strconv"
)

func main() {
	var operand_1 int
	var operand_2 int
	var operation string

	//FIRST OPERAND
	fmt.Print("Input 1 : ")
	fmt.Scan(&operand_1)

	//OPERATION
	fmt.Print("What operation : ")
	fmt.Scan(&operation)

	//SECOND OPERAND
	fmt.Print("Input 2 : ")
	fmt.Scan(&operand_2)

	if operation == "+" {
		fmt.Println("ANSWER :" + strconv.Itoa(operand_1+operand_2))
	} else if operation == "-" {
		fmt.Println("ANSWER :" + strconv.Itoa(operand_1-operand_2))
	} else if operation == "*" {
		fmt.Println("ANSWER :" + strconv.Itoa(operand_1*operand_2))
	} else if operation == "/" {
		fmt.Println("ANSWER :" + strconv.Itoa(operand_1/operand_2))
	}

}
