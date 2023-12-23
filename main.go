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
	wordCounter := wc.Bool("w", false, "to set word countet")

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
			os.Exit(1)
		}
		if *byteCounter {
			fmt.Printf("%v %v", len(cnt), fileName)
		} else if *lineCounter {
			lines := splitter(cnt, "\n")
			fmt.Printf("%v %v", len(lines), fileName)
		} else if *wordCounter {
			var count int

			lines := splitter(cnt, "\n")

			for _, line := range lines {
				line = strings.TrimSpace(line)
				arr := strings.Fields(line)
				count = count + len(arr)
			}

			fmt.Printf("%v %v", count, fileName)
		} else {
			fmt.Printf("No option choosen")
		}
	}
}

func splitter(text []byte, sep string) []string {
	str := string(text)
	arr := strings.FieldsFunc(str, fun)
	return arr
}

func fun(r rune) bool {
	return r == '\n'
}
