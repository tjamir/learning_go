package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	args := os.Args[1:]

	for _, filename := range args {
		fmt.Printf("\n===Printing file %v===\n\n", filename)
		file, error := os.Open(filename)
		if (error) != nil {
			fmt.Printf("Error attempting to open file %v:  %v", filename, error)
		}
		io.Copy(os.Stdout, file)
	}
}
