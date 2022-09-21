package src

import (
	"errors"
	"strconv"
)

func IsOperator(s string) bool {
	if s == "+" || s == "-" || s == "*" || s == "/" {
		return true
	}
	return false
}

func GetIntVal(s string) (int, error) {
	if IsOperator(s) == false {
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
		if IsValidArg(int(s[0])) == false {
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

func GetMultiPlaceValue(numbers []int, operators []int) int {
	var index int
	if CheckForOperatorPrecedence(operators) {
		index = GetIndex(operators, true)
	} else {
		index = GetIndex(operators, false)
	}
	buff := ""
	for _, n := range numbers[index+1:] {
		buff += strconv.Itoa(n)
	}
	v, _ := strconv.Atoi(buff)
	return v
}
