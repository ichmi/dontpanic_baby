package main

import (
	"app/src"
	"testing"
)

func TestValidateArgs_pass_invalid_number_of_args_expect_error(t *testing.T) {
	var tests = []struct {
		isErr bool
		args  []string
	}{
		{true, []string{}},
		{true, []string{""}},
		{true, []string{"", ""}},
		{true, []string{"", "", ""}},
		{true, []string{"", "", "", ""}},
		{true, []string{"", "", "", "", ""}},
		{true, []string{"", "", "", "", "", "", ""}},
	}
	for _, tt := range tests {
		err := src.ValidateNumArgs(tt.args)
		if tt.isErr {
			if err == nil {
				t.Error("expected a error but we didnt get one")
			}
		} else {
			if err != nil {
				t.Error("received an error where we didnt expected")
			}
		}
	}
}

func TestValidateArgs_passing_valid_number_of_args_expect_nil(t *testing.T) {
	args := []string{"", "", "", "", "", ""}
	err := src.ValidateNumArgs(args)
	if err != nil {
		t.Errorf("received an error where we didnt expected %s", err)
	}
}

func TestCheckArgLen_passing_two_digits_expect_error(t *testing.T) {
	test := "10"
	err := src.CheckArgLen(test)
	if err == nil {
		t.Error(err)
	}
}
func TestCheckArgLen_passing_multiple_digits_expect_error(t *testing.T) {
	test := "4096"
	err := src.CheckArgLen(test)
	if err == nil {
		t.Error(err)
	}
}

func TestIsValidArg_passing_invalid_arg_expect_false(t *testing.T) {
	expected := false
	symbols := []int{'_', '"', '?', '>', '.', '!', '<', '=', '~', '`', ' ', '#'}
	for _, symbol := range symbols {
		got := src.IsValidArg(symbol)
		if got != expected {
			t.Errorf("expected %t but got %t", expected, got)
		}
	}
}
func TestIsValidArg_passing_valid_args_expect_true(t *testing.T) {
	expected := true
	symbols := []int{'+', '-', '*', '/'}
	for _, symbol := range symbols {
		got := src.IsValidArg(symbol)
		if got != expected {
			t.Errorf("expected %t but got %t", expected, got)
		}
	}
}

func TestCheckWhetherArgIsDigit_passing_nondigit_arg_expect_false(t *testing.T) {
	expected := false
	symbols := []int{'_', '\'', '#', ',', ' ', '=', '<', '?', '>', '!', '"', 'x', '-', '+', '*', '/', -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, symbol := range symbols {
		got := src.IsDigit(symbol)
		if got != expected {
			t.Errorf("expected %t but got %t (%c is valid)", expected, got, symbol)
		}
	}
}

func TestCheckWhetherArgIsDigit_passing_digit_arg_expect_true(t *testing.T) {
	expected := true
	symbols := []int{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0', 48, 49, 50, 51, 52, 53, 54, 55, 56, 57}
	for _, symbol := range symbols {
		got := src.IsDigit(symbol)
		if got != expected {
			t.Errorf("expected %t but got %t (%c is invalid)", expected, got, symbol)
		}
	}
}

func TestValidateArguments_passing_valid_args_expect_nil(t *testing.T) {
	var tests = []struct {
		args []string
	}{
		{[]string{"0", "0", "0", "0", "0", "0"}},
		{[]string{"1", "2", "4", "8", "9", "0"}},
		{[]string{"4", "2", "+", "0", "0", "0"}},
		{[]string{"2", "0", "*", "2", "+", "2"}}, // 42
		{[]string{"1", "6", "8", "/", "0", "4"}},
		{[]string{"8", "*", "8", "-", "2", "2"}},
		{[]string{"8", "*", "9", "-", "3", "0"}}, // 42
	}
	for i, tt := range tests {
		err := src.ValidateArguments(tt.args)
		if err != nil {
			t.Errorf("test %d: got error `%s` where we didnt expected", i, err)
		}
	}
}

func TestValidateArguments_passing_invalid_args_expect_error(t *testing.T) {
	var tests = []struct {
		args []string
	}{
		{[]string{}},
		{[]string{"0", "0", "0", "0", "0"}},
		{[]string{"0", "0", "0", "0", "0", "0", "0"}},
		{[]string{"10", "0", "0", "0", "0", "0"}},
		{[]string{"10", "0", "0", "0", "0", "10"}},
		{[]string{"-1", "0", "0", "0", "0", "0"}},
		{[]string{"0", "0", "0", "0", "0", "-1"}},
		{[]string{"1", "1025", "0", "0", "2", "2"}},
		{[]string{"1", "1", "1", "0", "2", "4096321233"}},
		{[]string{"1", "1", "1", "0324234", "2", "40"}},
		{[]string{"42", "1", "1", "1", "2", "4"}},
		{[]string{"42", "3443", "12349034", "324", "342142", "222234344"}},
		{[]string{"9223372036854775807", "0", "0", "0", "0", "0"}},
		{[]string{"0", "0", "0", "0", "0", "9223372036854775807"}},
		{[]string{"_", "0", "0", "0", "0", "0"}},
		{[]string{"0", "_", "0", "0", "0", "0"}},
		{[]string{"0", "0", "_", "0", "0", "0"}},
		{[]string{"0", "0", "0", "_", "0", "0"}},
		{[]string{"0", "0", "0", "0", "_", "0"}},
		{[]string{"0", "0", "0", "0", "0", "_"}},
		{[]string{"A", "0", "0", "0", "0", "0"}},
		{[]string{"0", "0", "0", "0", "0", "a"}},
		{[]string{"1", "0", "x", "2", "2", "0"}},
		{[]string{"*", "0", "+", "2", "2", "0"}},
		{[]string{"/", "0", "+", "2", "2", "2"}},
		{[]string{"2", "0", "+", "2", "2", "+"}},
		{[]string{"2", "0", "+", "2", "2", "-"}},
		{[]string{"2", "0", "+", "2", "2", "/"}},
		{[]string{"2", "0", "+", "2", "2", "*"}},
	}
	for i, tt := range tests {
		err := src.ValidateArguments(tt.args)
		if err == nil {
			t.Errorf("test %d: expected error. didnt get one", i)
		}
	}
}

func TestCheckFirstSpot_passing_invalid_first_args_expect_error(t *testing.T) {
	var tests = []struct {
		args []string
	}{
		{[]string{"*", "0", "0", "0", "0"}},
		{[]string{"/", "0", "0", "0", "0"}},
	}
	for i, tt := range tests {
		if err := src.CheckFirstSpot(tt.args); err == nil {
			t.Errorf("test %d: expected error. didnt get one", i)
		}
	}
}

func TestCheckFirstSpot_passing_valid_first_args(t *testing.T) {
	var tests = []struct {
		args []string
	}{
		{[]string{"+", "-", "1", "+", "0"}},
		{[]string{"-", "*", "3", "+", "5"}},
		{[]string{"2", "/", "1", "0", "9"}},
		{[]string{"5", "*", "1", "0", "1"}},
		{[]string{"0", "*", "1", "0", "1"}},
		{[]string{"1", "*", "1", "0", "1"}},
	}
	for i, tt := range tests {
		if err := src.CheckFirstSpot(tt.args); err != nil {
			t.Errorf("test %d: receive error `%s` where we didnt expected", i, err)
		}
	}
}

func TestCheckLastSpot_passing_invalid_last_args_expect_error(t *testing.T) {
	var tests = []struct {
		args []string
	}{
		{[]string{"0", "0", "0", "0", "+"}},
		{[]string{"0", "0", "0", "0", "-"}},
		{[]string{"0", "0", "0", "0", "*"}},
		{[]string{"0", "0", "0", "0", "/"}},
	}
	for i, tt := range tests {
		if err := src.CheckLastSpot(tt.args); err == nil {
			t.Errorf("test %d: expected error. didnt get one", i)
		}
	}
}

func TestCheckLastSpot_passing_valid_last_args(t *testing.T) {
	var tests = []struct {
		args []string
	}{
		{[]string{"1", "-", "1", "+", "0"}},
		{[]string{"5", "*", "3", "+", "5"}},
		{[]string{"9", "/", "1", "0", "9"}},
		{[]string{"5", "*", "1", "0", "1"}},
	}
	for i, tt := range tests {
		if err := src.CheckLastSpot(tt.args); err != nil {
			t.Errorf("test %d: receive error `%s` where we didnt expected", i, err)
		}
	}
}
