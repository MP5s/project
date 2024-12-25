package application

import (
	"calc_service/pkg/calculator"
	"fmt"
)

func ProcessExpression(expression string) (string, error) {
	if !calculator.IsValidExpression(expression) {
		return "", fmt.Errorf("expression is not valid")
	}
	return calculator.Calculate(expression)
}
