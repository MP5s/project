package calculator

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Calculate(expression string) (string, error) {
	result, err := eval(expression)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%f", result), nil
}

func eval(expression string) (float64, error) {
	expression = strings.ReplaceAll(expression, " ", "")

	if !IsValidExpression(expression) {
		return 0, fmt.Errorf("invalid expression")
	}

	postfix, err := divideByPriority(expression)
	if err != nil {
		return 0, err
	}

	return calculateByPriority(postfix)
}

func divideByPriority(expression string) ([]string, error) {
	var output []string
	var stack []string

	precedence := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
		"(": 0,
	}

	for _, token := range tokenize(expression) {
		if isNumber(token) {
			output = append(output, token)
		} else if token == "(" {
			stack = append(stack, token)
		} else if token == ")" {
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				return nil, fmt.Errorf("mismatched parentheses")
			}
			stack = stack[:len(stack)-1]
		} else if isOperator(token) {
			for len(stack) > 0 && precedence[stack[len(stack)-1]] >= precedence[token] {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		} else {
			return nil, fmt.Errorf("expression is not valid")
		}
	}

	for len(stack) > 0 {
		if stack[len(stack)-1] == "(" {
			return nil, fmt.Errorf("mismatched parentheses")
		}
		output = append(output, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return output, nil
}

func calculateByPriority(postfix []string) (float64, error) {
	var stack []float64

	for _, token := range postfix {
		if isNumber(token) {
			num, _ := strconv.ParseFloat(token, 64)
			stack = append(stack, num)
		} else if isOperator(token) {
			if len(stack) < 2 {
				return 0, fmt.Errorf("expression is not valid")
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			switch token {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				if b == 0 {
					return 0, fmt.Errorf("division by zero")
				}
				stack = append(stack, a/b)
			}
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("invalid expression")
	}

	return stack[0], nil
}

func tokenize(expression string) []string {
	var tokens []string
	var currentToken strings.Builder

	for _, char := range expression {
		if isOperator(string(char)) || string(char) == "(" || string(char) == ")" {
			if currentToken.Len() > 0 {
				tokens = append(tokens, currentToken.String())
				currentToken.Reset()
			}
			tokens = append(tokens, string(char))
		} else {
			currentToken.WriteRune(char)
		}
	}

	if currentToken.Len() > 0 {
		tokens = append(tokens, currentToken.String())
	}

	return tokens
}

func isNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func isOperator(s string) bool {
	return s == "+" || s == "-" || s == "*" || s == "/"
}

func IsValidExpression(expression string) bool {
	valid := regexp.MustCompile(`^[0-9+\-*/(). ]+$`).MatchString(expression)
	if !valid {
		return false
	}

	balance := 0
	for _, char := range expression {
		if char == '(' {
			balance++
		} else if char == ')' {
			balance--
			if balance < 0 {
				return false
			}
		}
	}
	return balance == 0
}
