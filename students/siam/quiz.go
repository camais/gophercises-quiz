package main

import (
	"flag"
	"fmt"
)

func main() {
	filePathPtr := flag.String("csv", "./problems.csv", "Path to the quiz problem set.")

	flag.Parse()

	fmt.Println("File Path: ", *filePathPtr)

}
