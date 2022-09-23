package main

import (
	"strings"
	"testing"
)

func TestGetIntVal_passing_operators_expect_minus_one_as_value(t *testing.T) {
	test := []string{"+", "-", "*", "/"}
	for _, s := range test {
		v, _ := GetIntVal(s)
		if v != -1 {
			t.Errorf("expected -1 but got %d", v)
		}
	}
}

func TestGetIntVal_passing_valid_numbers_str_expect_int_values(t *testing.T) {
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	test := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for i, s := range test {
		v, _ := GetIntVal(s)
		if v != expected[i] {
			t.Errorf("expected %d and got %d", expected[i], v)
		}
	}
}

func TestGetIntVal_passing_invalid_numbers_str_expect_error(t *testing.T) {
	test := []string{".", "a", "#", "^", ",", ")", ";", ":", "=", "x"}
	for i, s := range test {
		_, err := GetIntVal(s)
		if err == nil {
			t.Errorf("test %d: expect error. we didnt get one", i)
		}
	}
}

func TestIsOperator_passing_valid_operator_expect_true(t *testing.T) {
	operators := []string{"-", "+", "/", "*"}
	for i, s := range operators {
		op := IsOperator(s)
		if op != true {
			t.Errorf("test %d: expect true but got false", i)
		}
	}
}

func TestIsOperator_passing_invalid_operator_expect_false(t *testing.T) {
	operators := []string{">", "<", "=", ".", "x", "(", ")", "{", "}", "1", "2", "9", "0"}
	for i, s := range operators {
		op := IsOperator(s)
		if op != false {
			t.Errorf("test %d: expect false but got true", i)
		}
	}
}

func TestGetSliceOfOperators_passing_numbers_expect_zeros(t *testing.T) {
	var results = []struct {
		args []int
	}{
		{[]int{-1, -1, -1, -1, -1, -1}},
		{[]int{-1, -1, -1, -1, -1, -1}},
		{[]int{-1, -1, -1, -1, -1, -1}},
	}
	var tests = []struct {
		args []string
	}{
		{strings.Split("1 1 1 1 1 1", " ")},
		{strings.Split("0 1 2 3 4 5", " ")},
		{strings.Split("9 8 7 6 5 4", " ")},
	}
	for i, tt := range tests {
		operators := GetSliceOfOperators(tt.args)
		for j, v := range operators {
			if v != results[i].args[j] {
				t.Errorf("test %d: expect %d but got %d", i, results[i].args[j], v)
			}
		}
	}
}

func TestGetSliceOfOperators_passing_operators_and_numbers(t *testing.T) {
	var results = []struct {
		args []int
	}{
		{[]int{-1, '+', '+', -1, -1, -1}},
		{[]int{-1, -1, '*', -1, '-', -1}},
		{[]int{'/', -1, -1, -1, -1, -1}},
		{[]int{'/', -1, -1, -1, -1, '*'}},
		{[]int{-1, -1, -1, -1, -1, '*'}},
		{[]int{'+', -1, '-', '/', '/', -1}},
	}
	var tests = []struct {
		args []string
	}{
		{strings.Split("1 + + 1 1 1", " ")},
		{strings.Split("0 1 * 3 - 5", " ")},
		{strings.Split("/ 8 7 6 5 4", " ")},
		{strings.Split("/ 8 7 6 5 *", " ")},
		{strings.Split("0 8 7 6 5 *", " ")},
		{strings.Split("+ 8 - / / 3", " ")},
	}
	for i, tt := range tests {
		operators := GetSliceOfOperators(tt.args)
		for j, v := range operators {
			if v != results[i].args[j] {
				t.Errorf("test %d: expect %d but got %d", i, results[i].args[j], v)
			}
		}
	}
}

func TestGetSliceOfOperators_passing_operators_only(t *testing.T) {
	var results = []struct {
		args []int
	}{
		{[]int{'+', '-', '*', '/', '*', '+'}},
		{[]int{'-', '+', '+', '+', '/', '+'}},
		{[]int{'-', '-', '-', '-', '-', '-'}},
		{[]int{'-', '/', '*', '-', '-', '-'}},
		{[]int{'/', '+', '+', '+', '/', '/'}},
		{[]int{'-', '*', '/', '+', '*', '*'}},
	}
	var tests = []struct {
		args []string
	}{
		{strings.Split("+ - * / * +", " ")},
		{strings.Split("- + + + / +", " ")},
		{strings.Split("- - - - - -", " ")},
		{strings.Split("- / * - - -", " ")},
		{strings.Split("/ + + + / /", " ")},
		{strings.Split("- * / + * *", " ")},
	}
	for i, tt := range tests {
		operators := GetSliceOfOperators(tt.args)
		for j, v := range operators {
			if v != results[i].args[j] {
				t.Errorf("test %d: expect %d but got %d", i, results[i].args[j], v)
			}
		}
	}
}

func TestGetSliceOfNumbers_passing_numbers_only(t *testing.T) {
	var results = []struct {
		args []int
	}{
		{[]int{1, 0, 1, 0, 4, 0}},
		{[]int{9, 2, 3, 6, 0, 2}},
		{[]int{4, 2, 2, 7, 4, 9}},
	}
	var tests = []struct {
		args []string
	}{
		{strings.Split("1 0 1 0 4 0", " ")},
		{strings.Split("9 2 3 6 0 2", " ")},
		{strings.Split("4 2 2 7 4 9", " ")},
	}
	for i, tt := range tests {
		numbers := GetSliceOfNumbers(tt.args)
		for j, v := range numbers {
			if v != results[i].args[j] {
				t.Errorf("test %d: expect %d but got %d", i, results[i].args[j], v)
			}
		}
	}
}

func TestGetSliceOfNumbers_passing_operators_only(t *testing.T) {
	var results = []struct {
		args []int
	}{
		{[]int{-1, -1, -1, -1, -1, -1}},
		{[]int{-1, -1, -1, -1, -1, -1}},
		{[]int{-1, -1, -1, -1, -1, -1}},
		{[]int{-1, -1, -1, -1, -1, -1}},
		{[]int{-1, -1, -1, -1, -1, -1}},
	}
	var tests = []struct {
		args []string
	}{
		{strings.Split("+ + + + + +", " ")},
		{strings.Split("- - - - - -", " ")},
		{strings.Split("* * * * * *", " ")},
		{strings.Split("/ / / / / /", " ")},
		{strings.Split("/ * - + / *", " ")},
	}
	for i, tt := range tests {
		numbers := GetSliceOfNumbers(tt.args)
		for j, v := range numbers {
			if v != results[i].args[j] {
				t.Errorf("test %d: expect %d but got %d", i, results[i].args[j], v)
			}
		}
	}
}
func TestGetSliceOfNumbers_passing_numbers_and_operators(t *testing.T) {
	var results = []struct {
		args []int
	}{
		{[]int{1, -1, 1, -1, 4, 0}},
		{[]int{9, -1, 3, -1, 0, 2}},
		{[]int{0, 0, 2, -1, 4, 0}},
	}
	var tests = []struct {
		args []string
	}{
		{strings.Split("1 + 1 + 4 0", " ")},
		{strings.Split("9 - 3 / 0 2", " ")},
		{strings.Split("0 0 2 * 4 0", " ")},
	}
	for i, tt := range tests {
		numbers := GetSliceOfNumbers(tt.args)
		for j, v := range numbers {
			if v != results[i].args[j] {
				t.Errorf("test %d: expect %d but got %d", i, results[i].args[j], v)
			}
		}
	}
}

func TestInitGameStructure_passing_invalid_args_expect_error(t *testing.T) {
	var tests = []struct {
		args []string
	}{
		{[]string{" ", "0", "0", "0", "0", "0"}},
		{[]string{"0", "0", "0", "0", "0", " "}},
		{[]string{"+", "!", "4", "8", "9", "0"}},
		{[]string{"4", "2", "+", "_", "0", "0"}},
		{[]string{"2", "0", "*", "2", "+", "x"}},
		{[]string{"1", ".", "8", "/", "0", "4"}},
	}
	for i, tt := range tests {
		_, _, err := InitGameStructure(tt.args)
		if err == nil {
			t.Errorf("test %d: expect error. we didnt get one", i)
		}
	}
}

func TestInitGameStructure_passing_valid_args_expect_no_error(t *testing.T) {
	var tests = []struct {
		args []string
	}{
		{strings.Split("1 + 1 + 4 0", " ")},
		{strings.Split("9 - 3 / 0 2", " ")},
		{strings.Split("0 0 2 * 4 0", " ")},
		{strings.Split("+ + 1 + 4 0", " ")},
		{strings.Split("9 - - - 1 0", " ")},
		{strings.Split("+ - / - * +", " ")},
		{strings.Split("0 0 0 0 0 0", " ")},
		{strings.Split("1 2 4 - 4 2", " ")},
		{strings.Split("2 0 * 2 + 2", " ")},
		{strings.Split("1 6 8 / 0 4", " ")},
		{strings.Split("8 * 8 - 2 2", " ")},
	}
	for i, tt := range tests {
		_, _, err := InitGameStructure(tt.args)
		if err != nil {
			t.Errorf("test %d: got error `%s` where we didnt expect", i, err)
		}
	}
}

func TestInitGameStructure_passing_operators_only(t *testing.T) {
	var results = []struct {
		args []int
	}{
		{[]int{'+', '-', '*', '/', '*', '+'}},
		{[]int{'-', '+', '+', '+', '/', '+'}},
		{[]int{'-', '-', '-', '-', '-', '-'}},
		{[]int{'-', '/', '*', '-', '-', '-'}},
		{[]int{'/', '+', '+', '+', '/', '/'}},
		{[]int{'-', '*', '/', '+', '*', '*'}},
	}
	var tests = []struct {
		args []string
	}{
		{strings.Split("+ - * / * +", " ")},
		{strings.Split("- + + + / +", " ")},
		{strings.Split("- - - - - -", " ")},
		{strings.Split("- / * - - -", " ")},
		{strings.Split("/ + + + / /", " ")},
		{strings.Split("- * / + * *", " ")},
	}
	for i, tt := range tests {
		_, operators, _ := InitGameStructure(tt.args)
		for j, v := range operators {
			if v != results[i].args[j] {
				t.Errorf("test %d: expect %d but got %d", i, results[i].args[j], v)
			}
		}
	}
}

func TestInitGameStructure_passing_numbers_only(t *testing.T) {
	var results = []struct {
		args []int
	}{
		{[]int{1, 1, 1, 1, 1, 1}},
		{[]int{0, 1, 2, 3, 4, 5}},
		{[]int{9, 8, 7, 6, 5, 4}},
	}
	var tests = []struct {
		args []string
	}{
		{strings.Split("1 1 1 1 1 1", " ")},
		{strings.Split("0 1 2 3 4 5", " ")},
		{strings.Split("9 8 7 6 5 4", " ")},
	}
	for i, tt := range tests {
		numbers, _, _ := InitGameStructure(tt.args)
		for j, v := range numbers {
			if v != results[i].args[j] {
				t.Errorf("test %d: expect %d but got %d", i, results[i].args[j], v)
			}
		}
	}
}

func TestInitGameStructure_checking_number_slice(t *testing.T) {
	var results = []struct {
		args []int
	}{
		{[]int{1, -1, 1, -1, 4, 0}},
		{[]int{9, -1, 3, -1, 0, 2}},
		{[]int{0, 0, 2, -1, 4, 0}},
	}
	var tests = []struct {
		args []string
	}{
		{strings.Split("1 + 1 + 4 0", " ")},
		{strings.Split("9 - 3 / 0 2", " ")},
		{strings.Split("0 0 2 * 4 0", " ")},
	}
	for i, tt := range tests {
		numbers, _, _ := InitGameStructure(tt.args)
		for j, v := range numbers {
			if v != results[i].args[j] {
				t.Errorf("test %d: expect %d but got %d", i, results[i].args[j], v)
			}
		}
	}
}

func TestInitGameStructure_checking_operators_slice(t *testing.T) {
	var results = []struct {
		args []int
	}{
		{[]int{-1, '+', -1, '+', -1, -1}},
		{[]int{-1, '-', -1, '/', -1, -1}},
		{[]int{-1, -1, -1, '*', -1, -1}},
	}
	var tests = []struct {
		args []string
	}{
		{strings.Split("1 + 1 + 4 0", " ")},
		{strings.Split("9 - 3 / 0 2", " ")},
		{strings.Split("0 0 2 * 4 0", " ")},
	}
	for i, tt := range tests {
		_, operators, _ := InitGameStructure(tt.args)
		for j, v := range operators {
			if v != results[i].args[j] {
				t.Errorf("test %d: expect %d but got %d", i, results[i].args[j], v)
			}
		}
	}
}

func TestIsSliceNumberedOnly_passing_numbers_expect_true(t *testing.T) {
	expected := true
	test := strings.Split("0 0 0 0 0 0", " ")
	numbers, _, _ := InitGameStructure(test)
	got := IsSliceNumberedOnly(numbers)
	if got != expected {
		t.Errorf("expected %v but got %v", expected, got)
	}
}

func TestIsSliceNumberedOnly_passing_numbers_and_operators_expect_false(t *testing.T) {
	expected := false
	var tests = []struct {
		args []string
	}{
		{strings.Split("1 + 1 + 4 0", " ")},
		{strings.Split("1 + 1 + + 0", " ")},
		{strings.Split("- 1 1 3 + 0", " ")},
		{strings.Split("- 4 - + + 3", " ")},
		{strings.Split("* 4 - + + 3", " ")},
		{strings.Split("1 4 / 1 0 3", " ")},
		{strings.Split("1 4 1 1 0 +", " ")},
		{strings.Split("1 4 1 1 * 2", " ")},
	}
	for i, tt := range tests {
		numbers, _, _ := InitGameStructure(tt.args)
		got := IsSliceNumberedOnly(numbers)
		if got != expected {
			t.Errorf("test %d: expected %v but got %v", i, expected, got)
		}
	}
}

func TestCheckForOperatorPrecedence_passing_non_precedence_arg_expect_false(t *testing.T) {
	expected := false
	var tests = []struct {
		args []string
	}{
		{strings.Split("1 + 1 + 4 0", " ")},
		{strings.Split("1 + 1 + + 0", " ")},
		{strings.Split("- 1 1 3 + 0", " ")},
		{strings.Split("- 4 - + + 3", " ")},
	}
	for i, tt := range tests {
		_, operators, _ := InitGameStructure(tt.args)
		got := CheckForOperatorPrecedence(operators)
		if got != expected {
			t.Errorf("test %d: expected %v but got %v", i, expected, got)
		}
	}
}

func TestCheckForOperatorPrecedence_passing_precedent_arg_expect_true(t *testing.T) {
	expected := true
	var tests = []struct {
		args []string
	}{
		{strings.Split("2 + 2 0 * 2", " ")},
		{strings.Split("2 0 / 2 * 2", " ")},
		{strings.Split("2 0 / 2 - 2", " ")},
		{strings.Split("2 + 2 0 + /", " ")},
		{strings.Split("2 + 2 0 - *", " ")},
		{strings.Split("2 0 * 2 / 2", " ")},
		{strings.Split("* 0 / 2 + 2", " ")},
		{strings.Split("/ 0 0 2 - 2", " ")},
	}
	for i, tt := range tests {
		_, operators, _ := InitGameStructure(tt.args)
		got := CheckForOperatorPrecedence(operators)
		if got != expected {
			t.Errorf("test %d: expected %v but got %v", i, expected, got)
		}
	}
}

func TestGetOperatorPrecedenceIndex_passing_non_precedence_arg_expect_minus_one(t *testing.T) {
	expected := -1
	var tests = []struct {
		args []string
	}{
		{strings.Split("1 + 1 + 4 0", " ")},
		{strings.Split("1 + 1 2 4 +", " ")},
		{strings.Split("1 + 1 + + 0", " ")},
		{strings.Split("- 1 1 3 + 0", " ")},
		{strings.Split("- 4 - + + 3", " ")},
	}
	for i, tt := range tests {
		_, operators, _ := InitGameStructure(tt.args)
		got := GetOperatorPrecedenceIndex(operators)
		if got != expected {
			t.Errorf("test %d: expected %v but got %v", i, expected, got)
		}
	}
}

func TestGetOperatorPrecedenceIndex_passing_precedence_arg_expect_its_index(t *testing.T) {
	expected := []int{3, 0, 0, 5, 5, 1, 1}
	var tests = []struct {
		args []string
	}{
		{strings.Split("1 + 1 * 4 0", " ")},
		{strings.Split("/ + 1 2 4 +", " ")},
		{strings.Split("* + 1 + + 0", " ")},
		{strings.Split("- 1 1 3 + /", " ")},
		{strings.Split("- 4 - + + *", " ")},
		{strings.Split("9 / 3 + * 2", " ")},
		{strings.Split("9 * 2 / 2 0", " ")},
	}
	for i, tt := range tests {
		_, operators, _ := InitGameStructure(tt.args)
		got := GetOperatorPrecedenceIndex(operators)
		if got != expected[i] {
			t.Errorf("test %d: expected %v but got %v", i, expected[i], got)
		}
	}
}

func TestGetOperatorIndex_passing_non_plus_minus_operator_expect_minus_one(t *testing.T) {
	expected := -1
	var tests = []struct {
		args []string
	}{
		{strings.Split("2 * 2 * 2 0", " ")},
		{strings.Split("/ / 1 2 4 *", " ")},
		{strings.Split("* * 1 * / 0", " ")},
		{strings.Split("/ 1 1 3 2 /", " ")},
	}
	for i, tt := range tests {
		_, operators, _ := InitGameStructure(tt.args)
		got := GetOperatorIndex(operators)
		if got != expected {
			t.Errorf("test %d: expected %v but got %v", i, expected, got)
		}
	}
}

func TestGetOperatorIndex_passing_plus_and_minus_operators_expect_its_index(t *testing.T) {
	expected := []int{4, 1, 0, 0, 5, 5, 2, 3, 1, 0}
	var tests = []struct {
		args []string
	}{
		{strings.Split("0 * 4 0 + 5", " ")},
		{strings.Split("/ + 1 2 4 *", " ")},
		{strings.Split("- * 1 * / 0", " ")},
		{strings.Split("+ 1 1 3 2 /", " ")},
		{strings.Split("2 1 1 3 2 +", " ")},
		{strings.Split("* 1 1 3 2 -", " ")},
		{strings.Split("* 1 + 3 2 -", " ")},
		{strings.Split("* 1 / - 2 -", " ")},
		{strings.Split("* + + - + -", " ")},
		{strings.Split("- + + - + -", " ")},
	}
	for i, tt := range tests {
		_, operators, _ := InitGameStructure(tt.args)
		got := GetOperatorIndex(operators)
		if got != expected[i] {
			t.Errorf("test %d: expected %v but got %v", i, expected[i], got)
		}
	}
}

func TestGetIndex_passing_precedence_arg_with_true_flag_expect_its_index(t *testing.T) {
	expected := []int{4, 1, 0, 1, 5, 0}
	var tests = []struct {
		args []string
	}{
		{strings.Split("- 5 + 5 * 2", " ")},
		{strings.Split("0 * 4 0 + 5", " ")},
		{strings.Split("/ + 1 2 4 *", " ")},
		{strings.Split("- * 1 * / 0", " ")},
		{strings.Split("+ 1 1 3 2 /", " ")},
		{strings.Split("* 1 1 3 2 -", " ")},
	}
	for i, tt := range tests {
		_, operators, _ := InitGameStructure(tt.args)
		got := GetIndex(operators, true)
		if got != expected[i] {
			t.Errorf("test %d: expected %v but got %v", i, expected[i], got)
		}
	}
}

func TestGetIndex_passing_precedence_arg_with_false_flag_expect_index_of_non_precedence_operator(t *testing.T) {
	expected := []int{2, 0, 4, 1, 0, 5, 5, 0, 0, 1, 2}
	var tests = []struct {
		args []string
	}{
		{strings.Split("4 2 + 5 * 2", " ")},
		{strings.Split("- 5 + 5 * 2", " ")},
		{strings.Split("0 * 4 0 + 5", " ")},
		{strings.Split("/ + 1 2 4 *", " ")},
		{strings.Split("+ * 1 * / 0", " ")},
		{strings.Split("2 1 1 3 2 +", " ")},
		{strings.Split("* 1 1 3 2 -", " ")},
		{strings.Split("+ 1 1 3 2 -", " ")},
		{strings.Split("- 1 1 3 2 -", " ")},
		{strings.Split("2 - - - - -", " ")},
		{strings.Split("2 2 + - - -", " ")},
	}
	for i, tt := range tests {
		_, operators, _ := InitGameStructure(tt.args)
		got := GetIndex(operators, false)
		if got != expected[i] {
			t.Errorf("test %d: expected %v but got %v", i, expected[i], got)
		}
	}
}

func TestGetValuesAfter_passing_valid_args_expecting_its_numerical_value(t *testing.T) {
	expected := []int{50000, 5000, 500, 50, 5, 0}
	var tests = []struct {
		args []string
	}{
		{strings.Split("+ 5 0 0 0 0", " ")},
		{strings.Split("1 + 5 0 0 0", " ")},
		{strings.Split("1 0 + 5 0 0", " ")},
		{strings.Split("1 0 0 + 5 0", " ")},
		{strings.Split("1 0 0 0 + 5", " ")},
		{strings.Split("1 0 0 0 0 +", " ")},
	}
	for i, tt := range tests {
		numbers, operators, _ := InitGameStructure(tt.args)
		got := GetValuesAfter(numbers, operators)
		if got != expected[i] {
			t.Errorf("test %d: expected %v but got %v", i, expected[i], got)
		}
	}
}

func TestGetValuesAfter_passing_numerical_values_only(t *testing.T) {
	expected := []int{500000, 0, 42, 123456}
	var tests = []struct {
		args []string
	}{
		{strings.Split("5 0 0 0 0 0", " ")},
		{strings.Split("0 0 0 0 0 0", " ")},
		{strings.Split("0 0 0 0 4 2", " ")},
		{strings.Split("1 2 3 4 5 6", " ")},
	}
	for i, tt := range tests {
		numbers, operators, _ := InitGameStructure(tt.args)
		got := GetValuesAfter(numbers, operators)
		if got != expected[i] {
			t.Errorf("test %d: expected %v but got %v", i, expected[i], got)
		}
	}
}

func TestGetValuesBefore_passing_valid_args_expecting_its_numerical_value(t *testing.T) {
	expected := []int{50000, 5000, 500, 50, 5, 0}
	var tests = []struct {
		args []string
	}{
		{strings.Split("5 0 0 0 0 +", " ")},
		{strings.Split("5 0 0 0 + 1", " ")},
		{strings.Split("5 0 0 + 1 0", " ")},
		{strings.Split("5 0 + 1 0 0", " ")},
		{strings.Split("5 + 1 0 0 0", " ")},
		{strings.Split("+ 1 0 0 0 0", " ")},
	}
	for i, tt := range tests {
		numbers, operators, _ := InitGameStructure(tt.args)
		got := GetValuesBefore(numbers, operators)
		if got != expected[i] {
			t.Errorf("test %d: expected %v but got %v", i, expected[i], got)
		}
	}
}

func TestGetValuesBefore_passing_numerical_values_only(t *testing.T) {
	expected := []int{500000, 0, 42, 123456}
	var tests = []struct {
		args []string
	}{
		{strings.Split("5 0 0 0 0 0", " ")},
		{strings.Split("0 0 0 0 0 0", " ")},
		{strings.Split("0 0 0 0 4 2", " ")},
		{strings.Split("1 2 3 4 5 6", " ")},
	}
	for i, tt := range tests {
		numbers, operators, _ := InitGameStructure(tt.args)
		got := GetValuesBefore(numbers, operators)
		if got != expected[i] {
			t.Errorf("test %d: expected %v but got %v", i, expected[i], got)
		}
	}
}

// func TestCalculateOnce_passing_single_equations_expect_its_values(t *testing.T) {
// 	expected := []int{42, 101, 200, 1005, 123456, 100, 200, 5, 10000, 2000, 2, 4, 9, 0, 3000}
// 	var tests = []struct {
// 		args []string
// 	}{
// 		{strings.Split("0 0 5 0 - 8", " ")},
// 		{strings.Split("1 0 0 + 0 1", " ")},
// 		{strings.Split("1 0 1 + 9 9", " ")},
// 		{strings.Split("5 + 1 0 0 0", " ")},
// 		{strings.Split("1 2 3 4 5 6", " ")},
// 		{strings.Split("1 2 3 - 2 3", " ")},
// 		{strings.Split("1 0 0 0 / 5", " ")},
// 		{strings.Split("1 * 0 0 0 5", " ")},
// 		{strings.Split("2 * 5 0 0 0", " ")},
// 		{strings.Split("2 0 * 1 0 0", " ")},
// 		{strings.Split("1 0 / 0 0 5", " ")},
// 		{strings.Split("9 / 0 0 0 2", " ")},
// 		{strings.Split("9 / 0 0 0 1", " ")},
// 		{strings.Split("9 / 1 0 0 0", " ")},
// 		{strings.Split("9 0 0 0 / 3", " ")},
// 	}
// 	for i, tt := range tests {
// 		numbers, operators, _ := InitGameStructure(tt.args)
// 		got, _ := CalculateOnce(numbers, operators)
// 		if got != expected[i] {
// 			t.Errorf("test %d: expected %v but got %v", i, expected, got)
// 		}
// 	}
// }

// func TestCalculateOnce_passing_division_by_zero_expect_error(t *testing.T) {
// 	var tests = []struct {
// 		args []string
// 	}{
// 		{strings.Split("1 0 0 0 / 0", " ")},
// 		{strings.Split("1 / 0 0 0 0", " ")},
// 		{strings.Split("0 0 0 0 / 0", " ")},
// 		{strings.Split("0 / 0 0 0 0", " ")},
// 		{strings.Split("1 0 / 0 0 0", " ")},
// 		{strings.Split("9 / 0 0 0 0", " ")},
// 		{strings.Split("9 0 0 / 0 0", " ")},
// 	}
// 	for i, tt := range tests {
// 		numbers, operators, _ := InitGameStructure(tt.args)
// 		_, err := CalculateOnce(numbers, operators)
// 		if err == nil {
// 			t.Errorf("test %d: expected error. but we didnt get", i)
// 		}
// 	}
// }

func TestIsAvaiableInSolution_passing_NA_arg_expect_false(t *testing.T) {
	expected := false
	solution := "022+20"
	var tests = []struct {
		arg rune
	}{
		{'1'}, {'3'}, {'4'}, {'5'}, {'6'}, {'7'},
		{'8'}, {'9'}, {'-'}, {'*'}, {'/'},
	}
	for i, tt := range tests {
		got := IsAvaiableInSolution(tt.arg, solution)
		if got != expected {
			t.Errorf("test %d: expected %v but got %v", i, expected, got)
		}
	}
}

func TestIsAvaiableInSolution_passing_avaiable_arg_expect_false(t *testing.T) {
	expected := true
	solution := "022+20"
	var tests = []struct {
		arg rune
	}{
		{'0'}, {'2'}, {'+'},
	}
	for i, tt := range tests {
		got := IsAvaiableInSolution(tt.arg, solution)
		if got != expected {
			t.Errorf("test %d: expected %v but got %v", i, expected, got)
		}
	}
}

func TestIsInRightSpot_passing_arg_in_wrong_spot_expect_false(t *testing.T) {
	expected := false
	solution := "022+20"
	tests := []struct {
		arg rune
	}{
		{'2'},
		{'+'},
	}
	for i, tt := range tests {
		got := IsInRightSpot(tt.arg, rune(solution[0]))
		if got != expected {
			t.Errorf("test %d: expected %v but got %v", i, expected, got)
		}
	}
}

func TestIsInRightSpot_passing_arg_in_right_spot_expect_true(t *testing.T) {
	expected := true
	solution := "022+20"
	tests := []struct {
		arg rune
	}{
		{'+'},
	}
	for i, tt := range tests {
		got := IsInRightSpot(tt.arg, rune(solution[3]))
		if got != expected {
			t.Errorf("test %d: expected %v but got %v", i, expected, got)
		}
	}
}

func TestInitHint_passing_args_unavaiable_in_the_solution_expect_X_hints(t *testing.T) {
	solution := "222+20"
	expected := "XXXXXX"
	var tests = []struct {
		args string
	}{
		{"333-39"},
		{"3/3143"},
		{"8*5649"},
		{"7-35*9"},
	}
	for i, tt := range tests {
		got := GetHints(tt.args, solution)
		if got != expected {
			t.Errorf("test %d: expected %v but got %v", i, expected, got)
		}
	}
}

func TestInitHint_passing_args_avaiable_in_the_solution_but_in_the_wrong_spot_expect_T_hints(t *testing.T) {
	solution := "120/40"
	expected := []string{"XXXXXT", "XXXXTX", "XXXTXX", "XXTXXX", "XTXXXX", "TTTTTT", "TTXXXX", "TTTTXX"}
	var tests = []struct {
		args string
	}{
		{"333331"},
		{"333313"},
		{"333133"},
		{"331333"},
		{"313333"},
		{"20/104"},
		{"013*59"},
		{"4/1275"},
	}
	for i, tt := range tests {
		got := GetHints(tt.args, solution)
		if got != expected[i] {
			t.Errorf("test %d: expected %v but got %v", i, expected[i], got)
		}
	}
}

func TestInitHint_passing_args_avaiable_in_the_solution_that_is_in_the_right_spot_expect_C_hints(t *testing.T) {
	solution := "120/40"
	expected := []string{"XXXXXC", "XXXXCX", "XXXCXX", "XXCXXX", "XCXXXX", "CXXXXX"}
	var tests = []struct {
		args string
	}{
		{"333330"},
		{"333343"},
		{"333/33"},
		{"330333"},
		{"323333"},
		{"133333"},
	}
	for i, tt := range tests {
		got := GetHints(tt.args, solution)
		if got != expected[i] {
			t.Errorf("test %d: expected %v but got %v", i, expected[i], got)
		}
	}
}

func TestInitHint_passing_args_avaiable_in_the_solution_at_all_expect_TC_hints(t *testing.T) {
	solution := "120/40"
	expected := []string{"XXTXTC", "TCTXCC", "XXCXXC", "TTCCCX", "TCTTTC"}
	var tests = []struct {
		args string
	}{
		{"33/320"},
		{"/21340"},
		{"330-30"},
		{"/10/49"},
		{"42/010"},
	}
	for i, tt := range tests {
		got := GetHints(tt.args, solution)
		if got != expected[i] {
			t.Errorf("test %d: expected %v but got %v", i, expected[i], got)
		}
	}
}

func TestCheckNumberOfCalculationsAndPrecedence_passing_non_precedence_arg_expect_false(t *testing.T) {
	expected := false
	test := strings.Split("1 + 1 + 4 0", " ")
	numbers, operators, _ := InitGameStructure(test)
	got, _ := CheckNumberOfCalculationsAndPrecedence(numbers, operators)
	if got != expected {
		t.Errorf("expected %v but got %v", expected, got)
	}
}

func TestCheckNumberOfCalculationsAndPrecedence_passing_precedence_arg_expect_true(t *testing.T) {
	expected := true
	test := strings.Split("1 + 2 * 4 0", " ")
	numbers, operators, _ := InitGameStructure(test)
	got, _ := CheckNumberOfCalculationsAndPrecedence(numbers, operators)
	if got != expected {
		t.Errorf("expected %v but got %v", expected, got)
	}
}

func TestCheckNumberOfCalculationsAndPrecedence_passing_n_operators_expect_n_values(t *testing.T) {
	expected := []int{2, 1, 4, 0, 0, 0, 3, 1, 1, 1}
	var tests = []struct {
		args []string
	}{
		{strings.Split("1 + 1 + 4 0", " ")},
		{strings.Split("1 + + + 4 0", " ")},
		{strings.Split("1 + - * / 0", " ")},
		{strings.Split("1 0 0 0 0 0", " ")},
		{strings.Split("+ 0 0 0 0 0", " ")},
		{strings.Split("- 0 0 0 0 0", " ")},
		{strings.Split("1 - - + * 0", " ")},
		{strings.Split("1 2 0 0 * 0", " ")},
		{strings.Split("1 2 0 * * 0", " ")},
		{strings.Split("1 / / / / 0", " ")},
	}
	for i, tt := range tests {
		numbers, operators, _ := InitGameStructure(tt.args)
		_, got := CheckNumberOfCalculationsAndPrecedence(numbers, operators)
		if got != expected[i] {
			t.Errorf("test %d: expected %v but got %v", i, expected[i], got)
		}
	}
}

func TestSiToSs_passing_slice_of_ints_expect_slice_of_strings(t *testing.T) {
	expected := []struct {
		args []string
	}{
		{strings.Split("2 42 4 0", " ")},
		{strings.Split("1 43 43 0", " ")},
	}
	var tests = []struct {
		args []int
	}{
		{[]int{2, 42, 4, 0}},
		{[]int{1, 43, 43, 0}},
	}
	for i, tt := range tests {
		got := SiToSs(tt.args)
		for j, v := range expected[i].args {
			if got[j] != v {
				t.Errorf("test %d: expected %v but got %v", i, expected[i].args, got)
			}
		}
	}
}

// func TestCalculatePrecedenceFirst_passing_precedence_arg_expect_value_of_mult_div_only(t *testing.T) {
// 	expected := []int{80, 0, 8, 35, 0, 3}
// 	var tests = []struct {
// 		args []string
// 	}{
// 		{strings.Split("1 + 2 * 4 0", " ")},
// 		{strings.Split("1 + 2 0 * 0", " ")},
// 		{strings.Split("9 - 0 8 * 1", " ")},
// 		{strings.Split("9 - 0 7 * 5", " ")},
// 		{strings.Split("9 + 1 / 0 2", " ")},
// 		{strings.Split("+ 6 / 0 0 2", " ")},
// 	}
// 	for i, tt := range tests {
// 		numbers, operators, _ := InitGameStructure(tt.args)
// 		got, _, _ := CalculatePrecedenceFirst(numbers, operators)
// 		if got != expected[i] {
// 			t.Errorf("test %d: expected %v but got %v", i, expected[i], got)
// 		}
// 	}
// }

func TestCheckNumberOfOperations(t *testing.T) {
	expected := []int{0, 1, 0, 0, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2}
	var tests = []struct {
		args []string
	}{
		{strings.Split("1 1 5 5 2 4", " ")},
		{strings.Split("0 1 + 5 2 4", " ")},
		{strings.Split("+ 0 0 0 0 0", " ")},
		{strings.Split("- 0 0 0 0 1", " ")},
		{strings.Split("0 0 1 0 - 1", " ")},
		{strings.Split("1 + - 1 0 0", " ")},
		{strings.Split("1 - + 1 0 0", " ")},
		{strings.Split("- + 1 0 0 0", " ")},
		{strings.Split("+ - 1 0 0 0", " ")},
		{strings.Split("- 1 0 + 1 0", " ")},
		{strings.Split("- 1 0 + - 5", " ")},
		{strings.Split("- 1 + - 5 5", " ")},
		{strings.Split("- 1 * - 5 5", " ")},
		{strings.Split("- 1 / - 5 5", " ")},
		{strings.Split("1 0 0 * 0 5", " ")},
		{strings.Split("1 0 0 0 / 5", " ")},
		{strings.Split("- 1 0 0 / 5", " ")},
		{strings.Split("- 1 0 / - 5", " ")},
		{strings.Split("- 1 + 1 + 5", " ")},
		{strings.Split("+ 1 + 1 + 5", " ")},
		{strings.Split("1 + 1 + 4 5", " ")},
		{strings.Split("1 + - 1 + 1", " ")},
		{strings.Split("1 * - 1 - 1", " ")},
	}
	for i, tt := range tests {
		_, operators, _ := InitGameStructure(tt.args)
		got := CheckNumberOfOperations(operators)
		if got != expected[i] {
			t.Errorf("test %d: expected %v but got %v", i, expected[i], got)
		}
	}
}
func TestGetValue(t *testing.T) {
	expected := []int{115524, 15524, 0, 1}
	var tests = []struct {
		args []string
	}{
		{strings.Split("1 1 5 5 2 4", " ")},
		{strings.Split("0 1 5 5 2 4", " ")},
		{strings.Split("0 0 0 0 0 0", " ")},
		{strings.Split("0 0 0 0 0 1", " ")},
	}
	for i, tt := range tests {
		numbers, operators, _ := InitGameStructure(tt.args)
		got := GetInitialValue(numbers, operators)
		if got != expected[i] {
			t.Errorf("test %d: expected %v but got %v", i, expected[i], got)
		}
	}
}

func TestCalculateSingle_passing_single_operations_expect_correct_results(t *testing.T) {
	expected := []int{5525, 1555, 1555, 40, 2000, 0, -25, 5, 0, 25, 50, -1, -2, -20, 2, -5, -15, 95, 5, -1, 5, -15, -15, -10, -10}
	var tests = []struct {
		args []string
	}{
		{strings.Split("1 + 5 5 2 4", " ")},
		{strings.Split("1 5 5 5 + 0", " ")},
		{strings.Split("1 5 5 5 - 0", " ")},
		{strings.Split("1 0 * 0 0 4", " ")},
		{strings.Split("1 0 * 2 0 0", " ")},
		{strings.Split("1 0 * 0 0 0", " ")},
		{strings.Split("- 5 * 0 0 5", " ")},
		{strings.Split("2 5 / 0 0 5", " ")},
		{strings.Split("0 1 / 1 0 0", " ")},
		{strings.Split("- 5 * - 0 5", " ")},
		{strings.Split("- 1 0 * - 5", " ")},
		{strings.Split("- 0 0 1 * 1", " ")},
		{strings.Split("- 0 1 0 / 5", " ")},
		{strings.Split("1 0 0 / - 5", " ")},
		{strings.Split("- 1 0 / - 5", " ")},
		{strings.Split("- 1 0 + 0 5", " ")},
		{strings.Split("- 1 0 + - 5", " ")},
		{strings.Split("1 0 0 + - 5", " ")},
		{strings.Split("+ 0 0 + 0 5", " ")},
		{strings.Split("+ - 6 + 0 5", " ")},
		{strings.Split("+ 1 0 + - 5", " ")},
		{strings.Split("+ - 1 0 - 5", " ")},
		{strings.Split("- + 1 0 - 5", " ")},
		{strings.Split("- + 5 + - 5", " ")},
		{strings.Split("+ - 5 + - 5", " ")},
	}
	for i, tt := range tests {
		numbers, operators, _ := InitGameStructure(tt.args)
		got, _ := CalculateSingle(numbers, operators)
		if got != expected[i] {
			t.Errorf("test %d: expected %v but got %v", i, expected[i], got)
		}
	}
}

func TestCalculateSingle_DivisionByZero_expect_error(t *testing.T) {
	var tests = []struct {
		args []string
	}{
		{strings.Split("1 / 0 0 0 0", " ")},
		{strings.Split("1 0 / 0 0 0", " ")},
		{strings.Split("1 0 0 / 0 0", " ")},
		{strings.Split("1 0 0 0 / 0", " ")},
		{strings.Split("0 0 / 0 0 0", " ")},
		{strings.Split("- 1 / 0 0 0", " ")},
		{strings.Split("- 1 0 / 0 0", " ")},
		{strings.Split("- 1 0 0 / 0", " ")},
		{strings.Split("- 1 0 / - 0", " ")},
		{strings.Split("1 0 0 / - 0", " ")},
		{strings.Split("1 / - 0 0 0", " ")},
	}
	for i, tt := range tests {
		numbers, operators, _ := InitGameStructure(tt.args)
		_, err := CalculateSingle(numbers, operators)
		if err == nil {
			t.Errorf("test %d: expected error. didnt get one", i)
		}
	}
}

func TestIsPrecedence(t *testing.T) {
	expected := []bool{true, false, false, true, true, false, false, false, true}
	var tests = []struct {
		args []string
	}{
		{strings.Split("1 1 + 5 * 4", " ")},
		{strings.Split("1 + 2 + - 3", " ")},
		{strings.Split("- 1 + 5 - 2", " ")},
		{strings.Split("- 1 + 5 / 2", " ")},
		{strings.Split("1 + 1 * 1 2", " ")},
		{strings.Split("1 + 1 + 4 0", " ")},
		{strings.Split("2 / 1 + 4 0", " ")},
		{strings.Split("2 / - 1 + 2", " ")},
		{strings.Split("2 + 2 / - 2", " ")},
	}
	for i, tt := range tests {
		_, operators, _ := InitGameStructure(tt.args)
		got := IsPrecedence(operators)
		if got != expected[i] {
			t.Errorf("test %d: expected %v but got %v", i, expected[i], got)
		}
	}
}
func TestGetPrecedenceIndex(t *testing.T) {
	expected := []int{4, 3, 3, 3, 4}
	var tests = []struct {
		args []string
	}{
		{strings.Split("1 1 + 5 * 4", " ")},
		{strings.Split("1 1 2 / 2 4", " ")},
		{strings.Split("+ 1 2 / - 3", " ")},
		{strings.Split("1 + 2 * - 3", " ")},
		{strings.Split("- 2 - 1 / 3", " ")},
	}
	for i, tt := range tests {
		_, operators, _ := InitGameStructure(tt.args)
		got := GetPrecedenceIndex(operators)
		if got != expected[i] {
			t.Errorf("test %d: expected %v but got %v", i, expected[i], got)
		}
	}
}

func TestCalculateTwice(t *testing.T) {
	// expected := []int{30, 0, 7, -2}
	expected := []int{0, -3, 4, 17, 0, 0, -6, -5, 10, -8, 2, 11, 1, 1, -3, 20, 40, -9, -9, 0}
	var tests = []struct {
		args []string
	}{
		{strings.Split("- 1 + 1 + 0", " ")},
		{strings.Split("- 1 * 2 - 1", " ")},
		{strings.Split("- 1 + 2 + 3", " ")},
		{strings.Split("1 0 + 5 + 2", " ")},
		{strings.Split("+ 6 / 2 - 3", " ")},
		{strings.Split("1 + 2 + - 3", " ")},
		{strings.Split("- 2 - 1 - 3", " ")},
		{strings.Split("- 2 * 1 - 3", " ")},
		{strings.Split("6 / 3 + 0 8", " ")},
		{strings.Split("- 2 * 1 * 4", " ")},
		{strings.Split("0 + 1 + 0 1", " ")},
		{strings.Split("1 0 + 0 + 1", " ")},
		{strings.Split("+ 1 + 0 + 0", " ")},
		{strings.Split("+ 0 + 0 + 1", " ")},
		{strings.Split("- 1 - 1 - 1", " ")},
		{strings.Split("1 * 2 0 - 0", " ")},
		{strings.Split("1 * 2 0 * 2", " ")},
		{strings.Split("- 6 / 2 * 3", " ")},
		{strings.Split("- 6 * 3 / 2", " ")},
		{strings.Split("+ 6 / 2 - 3", " ")},
	}
	for i, tt := range tests {
		numbers, operators, _ := InitGameStructure(tt.args)
		got, _ := CalculateTwice(numbers, operators)
		if got != expected[i] {
			t.Errorf("test %d: expected %v but got %v", i, expected[i], got)
		}
	}
}

// func TestCalculateTwice_passing_sum_and_sub_equations_expect_results(t *testing.T) {
// 	// (1 + 1) + 40 = 2 + 40 = 42
// 	expected := []int{-151, -50}
// 	var tests = []struct {
// 		args []string
// 	}{
// 		{strings.Split("- 1 5 5 + 4", " ")},
// 		{strings.Split("1 - 5 5 + 4", " ")},
// 	}
// 	for i, tt := range tests {
// 		numbers, operators, _ := InitGameStructure(tt.args)
// 		got := CalculateTwice(numbers, operators)
// 		if got != expected[i] {
// 			t.Errorf("test %d: expected %v but got %v", i, expected[i], got)
// 		}
// 	}
// }

// func TestCalculateTwice_passing_sum_and_sub_equations_expect_results(t *testing.T) {
// 	// (1 + 1) + 40 = 2 + 40 = 42
// 	expected := []int{42, 62, 162, 0, -50, -151, -50}
// 	var tests = []struct {
// 		args []string
// 	}{
// 		// {strings.Split("1 + 1 + 4 0", " ")},
// 		// {strings.Split("+ 1 0 + 5 2", " ")},
// 		// {strings.Split("1 0 + 1 5 2", " ")},
// 		// {strings.Split("1 0 - 5 - 5", " ")},
// 		{strings.Split("- 1 5 5 + 4", " ")},
// 		{strings.Split("1 - 5 5 + 4", " ")},
// 	}
// 	for i, tt := range tests {
// 		numbers, operators, _ := InitGameStructure(tt.args)
// 		got := CalculateTwice(numbers, operators)
// 		if got != expected[i] {
// 			t.Errorf("test %d: expected %v but got %v", i, expected[i], got)
// 		}
// 	}
// }

// func TestCalculate3(t *testing.T) {
// 	expected := []int{6}
// 	var tests = []struct {
// 		args []string
// 	}{
// 		{strings.Split("2 * 3 + 4 0", " ")},
// 	}
// 	for i, tt := range tests {
// 		numbers, operators, _ := InitGameStructure(tt.args)
// 		got, _ := Calculate3(numbers, operators)
// 		if got != expected[i] {
// 			t.Errorf("test %d: expected %v but got %v", i, expected[i], got)
// 		}
// 	}
// }
