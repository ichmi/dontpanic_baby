package src

import (
	"errors"
	"fmt"
	"unicode"
)

func ValidateNumArgs(av []string) error {
	if len(av) != 6 {
		return errors.New("invalid number of arguments")
	}
	return nil
}

func CheckArgLen(s string) error {
	if len(s) > 1 {
		return errors.New("argument have more than two digits")
	}
	return nil
}

func IsDigit(ch int) bool {
	if unicode.IsDigit(rune(ch)) {
		return true
	}
	return false
}
func IsValidArg(ch int) bool {
	if (ch == '+' || ch == '-' || ch == '*' || ch == '/') || IsDigit(ch) {
		return true
	}
	return false
}

func ValidateEntry(av []string) error {
	if err := ValidateNumArgs(av); err != nil {
		return err
	}
	if err := CheckFirstSpot(av); err != nil {
		return err
	}
	if err := CheckLastSpot(av); err != nil {
		return err
	}
	return nil
}

func ValidateArguments(av []string) error {
	if err := ValidateEntry(av); err != nil {
		return err
	}
	for _, s := range av {
		if err := CheckArgLen(s); err != nil {
			return err
		}
		for _, ch := range s {
			if !IsValidArg(int(ch)) {
				return fmt.Errorf("argument `%c` is invalid", ch)
			}
		}
	}
	return nil
}

func CheckFirstSpot(ss []string) error {
	first := ss[0][0]
	switch first {
	case '*', '/':
		return errors.New("operator (* and /) shouldt be in the first spot")
	}
	return nil
}

func CheckLastSpot(ss []string) error {
	last := ss[len(ss)-1][0]
	switch last {
	case '+', '-', '*', '/':
		return errors.New("should not use any operator in the last spot")
	}
	return nil
}
