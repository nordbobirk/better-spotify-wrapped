package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World!")
	for i := range os.Args {
		fmt.Println(os.Args[i])
	}
}
