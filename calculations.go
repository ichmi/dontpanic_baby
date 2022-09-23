package main

import (
	"errors"
	"strconv"
)

func IsValidDoubleOperators(i int, op []int) bool {
	if (op[i] == '-' && op[i+1] == '+') || (op[i] == '+' && op[i+1] == '-') {
		return true
	}
	if (op[i] == '*' && op[i+1] == '+') || (op[i] == '*' && op[i+1] == '-') {
		return true
	}
	if (op[i] == '/' && op[i+1] == '+') || (op[i] == '/' && op[i+1] == '-') {
		return true
	}
	return false
}

func CheckNumberOfOperations(operators []int) int {
	n := 0
	i := 0
	if (operators[0] == '-' && operators[1] == '+') || (operators[0] == '+' && operators[1] == '-') {
		i = 2
	} else if operators[0] == '-' || operators[0] == '+' {
		i++
	}
	for i < len(operators) {
		if IsValidDoubleOperators(i, operators) {
			i++
			continue
		}
		if operators[i] != -1 {
			n++
		}
		i++
	}
	return n
}
func GetInitialValue(numbers []int, operators []int) int {
	buff := ""
	sign := 1
	i := 0
	if operators[0] == '-' && operators[1] == '+' {
		sign = -1
		i = 2
	} else if operators[0] == '-' {
		sign = -1
		i++
	} else if operators[0] == '+' {
		i++
	}
	for ; i < len(numbers) && numbers[i] != -1; i++ {
		buff += strconv.Itoa(numbers[i])
	}
	v, _ := strconv.Atoi(buff)
	return v * sign
}

func GetValue(numbers []int) int {
	buff := ""
	for i := 0; i < len(numbers) && numbers[i] != -1; i++ {
		buff += strconv.Itoa(numbers[i])
	}
	v, _ := strconv.Atoi(buff)
	return v
}

func CalculateSingle(numbers []int, operators []int) (int, error) {
	i := 0
	if operators[0] == '+' {
		i++
	}
	x := 0
	if numbers[i] > 9 {
		x = numbers[i]
	} else {
		x = GetInitialValue(numbers[i:], operators[i:])
	}
	if x < 0 {
		i++
	}
	if operators[0] == '-' && operators[1] == '+' {
		i++
	}
	for i < len(numbers) && numbers[i] != -1 {
		i++
	}
	var y int
	if i < len(numbers) {
		y = GetInitialValue(numbers[i+1:], operators[i+1:])
	} else {
		return x, nil
	}
	switch operators[i] {
	case '+':
		return x + y, nil
	case '-':
		return x - y, nil
	case '*':
		return x * y, nil
	case '/':
		if y == 0 {
			return 0, errors.New("division by zero")
		}
		return x / y, nil
	}
	return x, nil
}

func Calculate(numbers []int, operators []int) (int, error) {
	n := CheckNumberOfOperations(operators)
	result := 0
	switch n {
	case 0:
		result = GetValue(numbers)
	case 1:
		result, _ = CalculateSingle(numbers, operators)
	case 2:
		// if IsPrecedence(operators) {
		// 	i := GetPrecedenceIndex(operators)
		// }
		result, _ = CalculateTwice(numbers, operators)
	}
	return result, nil
}

func IsPrecedence(operators []int) bool {
	isSumSub := false
	isMulDiv := false
	i := 0
	if (operators[0] == '+' && operators[1] == '-') || (operators[0] == '-' && operators[1] == '+') {
		i = 2
	} else if operators[0] == '+' || operators[0] == '-' {
		i++
	}
	for i < len(operators) {
		if operators[i] == '+' || operators[i] == '-' {
			isSumSub = true
			break
		}
		i++
	}
	if isSumSub {
		for i < len(operators) {
			if operators[i] == '*' || operators[i] == '/' {
				isMulDiv = true
				break
			}
			i++
		}
	}
	return isMulDiv
}

func GetPrecedenceIndex(operators []int) int {
	sumSubIndex := 0
	mulDivIndex := 0
	i := 0
	if (operators[0] == '+' && operators[1] == '-') || (operators[0] == '-' && operators[i] == '+') {
		i++
	}
	if operators[0] == '+' || operators[0] == '-' {
		i++
	}
	for i < len(operators) {
		if (operators[i] == '+' || operators[i] == '-') && operators[i-1] != '*' && operators[i-1] != '/' {
			sumSubIndex = i
		} else if operators[i] == '*' || operators[i] == '/' {
			mulDivIndex = i
		}
		i++
	}
	if mulDivIndex > sumSubIndex {
		return mulDivIndex
	}
	return 0
}

func CalculateTwice(numbers []int, operators []int) (int, error) {
	i := 0
	if operators[0] == '+' {
		i++
	}
	x := GetInitialValue(numbers[i:], operators[i:])
	if x < 0 {
		i++
	}
	if operators[0] == '-' && operators[1] == '+' {
		i++
	}
	for i < len(numbers) && numbers[i] != -1 {
		i++
	}

	var y int
	if i < len(numbers) {
		y = GetInitialValue(numbers[i+1:], operators[i+1:])
	} else {
		return x, nil
	}
	if y > 9 {
		switch operators[i] {
		case '+':
			numbers[i+2] = x + y
		case '-':
			numbers[i+2] = x - y
		case '*':
			numbers[i+2] = x * y
		case '/':
			if y == 0 {
				return 0, errors.New("division by zero")
			}
			numbers[i+1] = x / y
		}
	} else {
		switch operators[i] {
		case '+':
			numbers[i+1] = x + y
		case '-':
			numbers[i+1] = x - y
		case '*':
			numbers[i+1] = x * y
		case '/':
			if y == 0 {
				return 0, errors.New("division by zero")
			}
			numbers[i+1] = x / y
		}
	}
	for i < len(numbers) && numbers[i+1] != -1 {
		i++
	}
	v, _ := CalculateSingle(numbers[i:], operators[i:])
	return v, nil
}