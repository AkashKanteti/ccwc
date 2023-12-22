package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	wc := flag.NewFlagSet("wc", flag.ExitOnError)
	byteCounter := wc.Bool("c", false, "to set byte counter")
	lineCounter := wc.Bool("l", false, "to set line countet")
	switch os.Args[1] {
	case "wc":
		wc.Parse(os.Args[2:])
		arguments := wc.Args()
		// fmt.Printf("%v", arguments)
		if len(arguments) < 1 {
			fmt.Printf("%v", errors.New("very few arguments"))
			os.Exit(1)
		}
		fileName := arguments[0]
		cnt, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Printf("%v", err)
		}
		if *byteCounter {
			fmt.Printf("%v %v", len(cnt), fileName)
		} else if *lineCounter {
			lines := splitter(cnt)
			fmt.Printf("%v %v", len(lines), fileName)
		} else {
			fmt.Printf("No option choosen")
		}
	}
}

func splitter(text []byte) []string {
	str := string(text)
	str = strings.Trim(str, " \n")
	arr := strings.Split(str, "\n")
	return arr
}
