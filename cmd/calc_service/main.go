package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

// RequestBody представляет тело запроса
type RequestBody struct {
	Expression string `json:"expression"`
}

// ResponseBody представляет тело ответа
type ResponseBody struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

// calculate вычисляет арифметическое выражение
func calculate(expression string) (string, error) {
	result, err := eval(expression)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%f", result), nil
}

// eval - реализация вычисления выражения с учетом порядка операций
func eval(expression string) (float64, error) {
	// Удаляются пробелы
	expression = strings.ReplaceAll(expression, " ", "")

	postfix, err := divideByPriority(expression)
	if err != nil {
		return 0, err
	}

	return calculateByPriority(postfix)
}

// Приоритет операций
func divideByPriority(expression string) ([]string, error) {
	var output []string
	var stack []string

	precedence := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
	}

	for _, token := range tokenize(expression) {
		if isNumber(token) {
			output = append(output, token)
		} else if isOperator(token) {
			for len(stack) > 0 && precedence[stack[len(stack)-1]] >= precedence[token] {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		} else {
			return nil, fmt.Errorf("Expression is not valid")
		}
	}

	for len(stack) > 0 {
		output = append(output, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return output, nil
}

// calculateByPriority вычисляет результат с учетом приоритетов операций
func calculateByPriority(postfix []string) (float64, error) {
	var stack []float64

	for _, token := range postfix {
		if isNumber(token) {
			num, _ := strconv.ParseFloat(token, 64)
			stack = append(stack, num)
		} else if isOperator(token) {
			if len(stack) < 2 {
				return 0, fmt.Errorf("Expression is not valid")
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
					return 0, fmt.Errorf("Division by zero")
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
		if isOperator(string(char)) {
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

// isNumber проверяет, является ли строка числом
func isNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// isOperator проверяет, является ли символ оператором
func isOperator(s string) bool {
	return s == "+" || s == "-" || s == "*" || s == "/"
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error": "Method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var reqBody RequestBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}

	// Проверка на валидность выражения
	if !isValidExpression(reqBody.Expression) {
		http.Error(w, `{"error": "Expression is not valid"}`, http.StatusUnprocessableEntity)
		return
	}

	result, err := calculate(reqBody.Expression)
	if err != nil {
		http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	response := ResponseBody{Result: result}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// isValidExpression проверяет, что выражение состоит только из цифр и операторов
func isValidExpression(expression string) bool {
	return regexp.MustCompile(`^[0-9+\-*/(). ]+$`).MatchString(expression)
}

func main() {
	http.HandleFunc("/api/v1/calculate", calculateHandler)
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
