package main

import (
	"fmt"
	. "github.com/mphilpot/gocyberq/cyberq"
	"os"
)

func main() {
	// `os.Args` provides access to raw command-line
	// arguments. Note that the first value in this slice
	// is the path to the program, and `os.Args[1:]`
	// holds the arguments to the program.
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]
	arg1 := os.Args[1]

	// You can get individual args with normal indexing.
	// arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg1)

	cyberq := &CyberQ{URL: arg1}

	fmt.Println(cyberq.All())
}
