package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fileName := flag.String("c", "test file", "for taking file name")
	flag.Parse()
	cnt, err := os.ReadFile(*fileName)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("%v %v", len(cnt), *fileName)
}
