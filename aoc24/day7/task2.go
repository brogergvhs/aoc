package main

import (
	"fmt"
	"strconv"
)

func evaluateExpression2(operands []int, operators []string) int {
	result := operands[0]
	for i, operator := range operators {
		switch operator {
		case "+":
			result += operands[i+1]
		case "*":
			result *= operands[i+1]
		case "||":
			result = concatenateNumbers(result, operands[i+1])
		}
	}
	return result
}

func concatenateNumbers(a, b int) int {
	concat, err := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
	if err != nil {
		fmt.Println("Error concatenating numbers:", err)
		return 0
	}
	return concat
}

