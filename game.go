package main

import (
	"errors"
	"strconv"
)

const SENT = 2147483647

func IsOperator(s string) bool {
	if s == "+" || s == "-" || s == "*" || s == "/" {
		return true
	}
	return false
}

func GetIntVal(s string) (int, error) {
	if !IsOperator(s) {
		intVal, err := strconv.Atoi(s)
		if err != nil {
			return -1, err
		}
		return intVal, nil
	}
	return -1, nil
}

func GetSliceOfOperators(ss []string) []int {
	var result []int
	for _, s := range ss {
		_, sign := strconv.Atoi(s)
		if sign != nil {
			switch {
			case s[0] == '+':
				result = append(result, '+')
			case s[0] == '-':
				result = append(result, '-')
			case s[0] == '*':
				result = append(result, '*')
			case s[0] == '/':
				result = append(result, '/')
			}
		} else {
			result = append(result, -1)
		}
	}
	return result
}

func GetSliceOfNumbers(ss []string) []int {
	var result []int
	for _, s := range ss {
		v, op := strconv.Atoi(s)
		if op != nil {
			result = append(result, -1)
		} else {
			result = append(result, v)
		}
	}
	return result
}

func InitGameStructure(ss []string) ([]int, []int, error) {
	for _, s := range ss {
		if !IsValidArg(int(s[0])) {
			return nil, nil, errors.New("invalid argument")
		}
	}
	return GetSliceOfNumbers(ss), GetSliceOfOperators(ss), nil
}

func IsSliceNumberedOnly(numbers []int) bool {
	for _, n := range numbers {
		if n == -1 {
			return false
		}
	}
	return true
}

func CheckForOperatorPrecedence(operators []int) bool {
	for _, n := range operators {
		if n == '*' || n == '/' {
			return true
		}
	}
	return false
}

func GetOperatorPrecedenceIndex(operators []int) int {
	for i, n := range operators {
		if n == '*' || n == '/' {
			return i
		}
	}
	return -1
}

func GetOperatorIndex(operators []int) int {
	for i, n := range operators {
		if n == '+' || n == '-' {
			return i
		}
	}
	return -1
}

func GetIndex(operators []int, precedence bool) int {
	if precedence {
		return GetOperatorPrecedenceIndex(operators)
	}
	return GetOperatorIndex(operators)
}

func GetValuesAfter(numbers []int, operators []int) int {
	var index int
	buff := ""
	if IsSliceNumberedOnly(numbers) {
		for i := 0; i < len(numbers); i++ {
			buff += strconv.Itoa(numbers[i])
		}
		v, _ := strconv.Atoi(buff)
		return v
	}
	if CheckForOperatorPrecedence(operators) {
		index = GetIndex(operators, true)
	} else {
		index = GetIndex(operators, false)
	}
	for _, n := range numbers[index+1:] {
		buff += strconv.Itoa(n)
	}
	v, _ := strconv.Atoi(buff)
	return v
}

func GetValuesBefore(numbers []int, operators []int) int {
	var index int
	buff := ""
	if IsSliceNumberedOnly(numbers) {
		for i := 0; i < len(numbers); i++ {
			buff += strconv.Itoa(numbers[i])
		}
		v, _ := strconv.Atoi(buff)
		return v
	}
	if CheckForOperatorPrecedence(operators) {
		index = GetIndex(operators, true)
	} else {
		index = GetIndex(operators, false)
	}
	for i := 0; i < index; i++ {
		buff += strconv.Itoa(numbers[i])
	}
	v, _ := strconv.Atoi(buff)
	return v
}

func DividedByZero(n int, op int) error {
	if op == '/' && n == 0 {
		return errors.New("cant divide by zero")
	}
	return nil
}

// func CalculateOnce(numbers []int, operators []int) (int, error) {
// 	var index int
// 	if CheckForOperatorPrecedence(operators) {
// 		index = GetIndex(operators, true)
// 	} else {
// 		index = GetIndex(operators, false)
// 	}
// 	before := GetValuesBefore(numbers, operators)
// 	after := GetValuesAfter(numbers, operators)
// 	if index != SENT {
// 		if err := DividedByZero(after, operators[index]); err != nil {
// 			return 0, err
// 		}
// 		switch operators[index] {
// 		case '+':
// 			return before + after, nil
// 		case '-':
// 			return before - after, nil
// 		case '*':
// 			return before * after, nil
// 		case '/':
// 			return before / after, nil
// 		}
// 	}
// 	return before, nil
// }

func CheckNumberOfCalculationsAndPrecedence(numbers []int, operators []int) (bool, int) {
	count := 0
	for i, v := range operators {
		if operators[0] == '+' || operators[0] == '-' {
			continue
		}
		if v != -1 {
			if i < len(operators) && v != operators[i+1] {
				count++
			}
		}
	}
	return CheckForOperatorPrecedence(operators), count
}

func IsAvaiableInSolution(b rune, solution string) bool {
	for _, c := range solution {
		if b == c {
			return true
		}
	}
	return false
}

func IsInRightSpot(try rune, solution rune) bool {
	return try == solution
}

func GetHints(try string, solution string) string {
	var hints [6]byte
	for i, c := range try {
		if IsAvaiableInSolution(c, solution) {
			if IsInRightSpot(c, rune(solution[i])) {
				hints[i] = 'C'
			} else {
				hints[i] = 'T'
			}
		} else {
			hints[i] = 'X'
		}
	}
	return string(hints[:])
}

// func _CalculateTwice(result *int, numbers []int, operators []int) {
// 	sign := 1
// 	i := 0
// 	if operators[i] == '+' {
// 		i++
// 	} else if operators[i] == '-' {
// 		sign = -1
// 	}
// 	x := GetValue(numbers[i:])
// 	for operators[i] == -1 {
// 		i++
// 	}
// 	y := GetValue(numbers[i+1:])

// 	*result = 0
// 	switch operators[i] {
// 	case '+':
// 		*result = (x + y) * sign
// 	case '-':
// 		*result = (x - y) * sign
// 	}
// 	fmt.Println(*result) // -54

// 	fmt.Println(i)            // 1
// 	fmt.Println(numbers[i])   // -1
// 	fmt.Println(operators[i]) // 45
// 	if numbers[i] != -1 {
// 		for operators[i] == -1 {
// 			i++
// 		}
// 		i++
// 	}

// 	y = GetValue(numbers[i+2:])
// 	switch operators[i+1] {
// 	case '+':
// 		*result += y
// 	case '-':
// 		*result -= y
// 	}
// 	fmt.Println(*result)
// }

func GetSign(operators []int) int {
	if operators[0] == '-' {
		return -1
	}
	return 1
}

// func _SumSubFirst(op, i, x, y, sign int) int {
// 	switch op {
// 	case '+':
// 		x += y
// 	case '-':
// 		x -= y
// 	}
// 	if sign == -1 {
// 		x = -x
// 	}
// 	return x
// }

// - 1
// func CalculateTwice(numbers []int, operators []int) int {
// 	i := 0
// 	sign := 1
// 	if sign = GetSign(operators); sign == -1 {
// 		i++
// 	}
// 	x := GetValue(numbers[i:])
// 	for numbers[i] != -1 {
// 		i++
// 	}
// 	y := GetValue(numbers[i+1:])
// 	x = _SumSubFirst(operators[i], i, x, y, sign)
// 	i++
// 	for numbers[i] != -1 {
// 		i++
// 	}
// 	y = GetValue(numbers[i+1:])
// 	result := _SumSubFirst(operators[i], i, x, y, 1)
// 	fmt.Println(result)
// 	return result
// }

// 1 + (2 * 40)
// 1 + 80
// 81
// func Calculate2(numbers []int, operators []int) (int, error) {
// 	y, index, err := CalculatePrecedenceFirst(numbers, operators)
// 	if err != nil {
// 		return 0, err
// 	}
// 	x := GetValue(numbers[:index])
// 	i := 0
// 	for operators[i] == -1 {
// 		i++
// 	}
// 	result := 0
// 	switch operators[i] {
// 	case '+':
// 		result = x + y
// 	case '-':
// 		result = x - y
// 	}
// 	return result, nil
// }

// func Calculate3(numbers []int, operators []int) (int, error) {
// 	x := GetValue(numbers, operators)
// 	i := 0
// 	for operators[i] == -1 {
// 		i++
// 	}
// 	y := GetValue(numbers[i+1:])

// 	result := 0
// 	switch operators[i] {
// 	case '*':
// 		result = x * y
// 	case '/':
// 		if y == 0 {
// 			return 0, errors.New("cant divide by zero")
// 		}
// 		result = x / y
// 	}
// 	return result, nil
// }

// func CalculatePrecedenceFirst(numbers []int, operators []int) (int, int, error) {
// 	var index int
// 	if CheckForOperatorPrecedence(operators) {
// 		index = GetIndex(operators, true)
// 	} else {
// 		index = GetIndex(operators, false)
// 	}
// 	n := GetSliceOfNumbers(SiToSs(numbers[index-1:]))
// 	o := GetSliceOfNumbers(SiToSs(operators[index-1:]))
// 	v, err := CalculateOnce(n, o)
// 	return v, index, err

// }

func SiToSs(numbers []int) []string {
	size := 0
	for i := 0; i < len(numbers); i++ {
		size++
	}
	buff := make([]string, size)
	for i, v := range numbers {
		buff[i] = strconv.Itoa(v)
	}
	return buff
}
