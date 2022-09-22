package main

import (
	"fmt"
	"os"
)

func main() {
	av := os.Args[1:]
	if err := ValidateArguments(av); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	numbers, operators, err := InitGameStructure(av)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("num", numbers)
	fmt.Println("ops", operators)
}
