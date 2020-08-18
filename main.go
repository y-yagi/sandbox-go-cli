package main

import (
	"fmt"
	"os"

	"github.com/y-yagi/goext/arr"
)

func main() {
	fmt.Printf("Args: %v\n", os.Args)
	fmt.Printf("Unique Args: %v\n", arr.UniqueStrings(os.Args))
}
