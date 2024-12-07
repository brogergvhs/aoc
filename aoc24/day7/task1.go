package main

func evaluateExpression(operands []int, operators []string) int {
	result := operands[0]
	for i, operator := range operators {
		switch operator {
		case "+":
			result += operands[i+1]
		case "*":
			result *= operands[i+1]
		}
	}
	return result
}
