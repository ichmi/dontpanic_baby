package main

import (
	"app/src"
	"fmt"
	"os"
)

func main() {
	av := os.Args[1:]
	if err := src.ValidateArguments(av); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	numbers, operators, err := src.InitGameStructure(av)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("num", numbers)
	fmt.Println("ops", operators)
}
