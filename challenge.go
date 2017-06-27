package main

import (
	"fmt"
	"os"

	"github.com/crerwin/stringchallenge/item"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Fprint(os.Stderr, "Use: stringchallenge (input(text,of),some,sort)\n")
		os.Exit(1)
	}
	item, err := item.CreateItem(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
	fmt.Println("Non-alphabetized:")
	fmt.Println(item.GetOutput(false))
	fmt.Println("\nAlphabetized:")
	fmt.Println(item.GetOutput(true))
}
