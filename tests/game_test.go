package main

import (
	"app/src"
	"strings"
	"testing"
)

func TestGetIntVal_passing_operators_expect_minus_one_as_value(t *testing.T) {
	test := []string{"+", "-", "*", "/"}
	for _, s := range test {
		v, _ := src.GetIntVal(s)
		if v != -1 {
			t.Errorf("expected -1 but got %d", v)
		}
	}
}

func TestGetIntVal_passing_valid_numbers_str_expect_int_values(t *testing.T) {
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	test := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for i, s := range test {
		v, _ := src.GetIntVal(s)
		if v != expected[i] {
			t.Errorf("expected %d and got %d", expected[i], v)
		}
	}
}

func TestGetIntVal_passing_invalid_numbers_str_expect_error(t *testing.T) {
	test := []string{".", "a", "#", "^", ",", ")", ";", ":", "=", "x"}
	for i, s := range test {
		_, err := src.GetIntVal(s)
		if err == nil {
			t.Errorf("test %d: expect error. we didnt get one", i)
		}
	}
}

func TestIsOperator_passing_valid_operator_expect_true(t *testing.T) {
	operators := []string{"-", "+", "/", "*"}
	for i, s := range operators {
		op := src.IsOperator(s)
		if op != true {
			t.Errorf("test %d: expect true but got false", i)
		}
	}
}

func TestIsOperator_passing_invalid_operator_expect_false(t *testing.T) {
	operators := []string{">", "<", "=", ".", "x", "(", ")", "{", "}", "1", "2", "9", "0"}
	for i, s := range operators {
		op := src.IsOperator(s)
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
		operators := src.GetSliceOfOperators(tt.args)
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
		operators := src.GetSliceOfOperators(tt.args)
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
		operators := src.GetSliceOfOperators(tt.args)
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
		numbers := src.GetSliceOfNumbers(tt.args)
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
		numbers := src.GetSliceOfNumbers(tt.args)
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
		numbers := src.GetSliceOfNumbers(tt.args)
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
		_, _, err := src.InitGameStructure(tt.args)
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
		_, _, err := src.InitGameStructure(tt.args)
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
		_, operators, _ := src.InitGameStructure(tt.args)
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
		numbers, _, _ := src.InitGameStructure(tt.args)
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
		numbers, _, _ := src.InitGameStructure(tt.args)
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
		_, operators, _ := src.InitGameStructure(tt.args)
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
	numbers, _, _ := src.InitGameStructure(test)
	got := src.IsSliceNumberedOnly(numbers)
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
		numbers, _, _ := src.InitGameStructure(tt.args)
		got := src.IsSliceNumberedOnly(numbers)
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
		_, operators, _ := src.InitGameStructure(tt.args)
		got := src.CheckForOperatorPrecedence(operators)
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
		_, operators, _ := src.InitGameStructure(tt.args)
		got := src.CheckForOperatorPrecedence(operators)
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
		_, operators, _ := src.InitGameStructure(tt.args)
		got := src.GetOperatorPrecedenceIndex(operators)
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
		_, operators, _ := src.InitGameStructure(tt.args)
		got := src.GetOperatorPrecedenceIndex(operators)
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
		_, operators, _ := src.InitGameStructure(tt.args)
		got := src.GetOperatorIndex(operators)
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
		_, operators, _ := src.InitGameStructure(tt.args)
		got := src.GetOperatorIndex(operators)
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
		_, operators, _ := src.InitGameStructure(tt.args)
		got := src.GetIndex(operators, true)
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
		_, operators, _ := src.InitGameStructure(tt.args)
		got := src.GetIndex(operators, false)
		if got != expected[i] {
			t.Errorf("test %d: expected %v but got %v", i, expected[i], got)
		}
	}
}

func TestGetMultiPlaceValue_passing_valid_args_expecting_its_numerical_value(t *testing.T) {
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
		numbers, operators, _ := src.InitGameStructure(tt.args)
		got := src.GetMultiPlaceValue(numbers, operators)
		if got != expected[i] {
			t.Errorf("test %d: expected %v but got %v", i, expected[i], got)
		}
	}
}
