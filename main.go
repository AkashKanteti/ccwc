package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var err error
	var cnt []byte
	fileName := ""
	wc := flag.NewFlagSet("wc", flag.ExitOnError)
	byteCounter := wc.Bool("c", false, "to set byte counter")
	lineCounter := wc.Bool("l", false, "to set line counter")
	wordCounter := wc.Bool("w", false, "to set word counter")
	characterCounter := wc.Bool("m", false, "to set character counter")

	switch os.Args[1] {
	case "wc":
		wc.Parse(os.Args[2:])
		arguments := wc.Args()
		if len(arguments) < 1 {
			in := bufio.NewReader(os.Stdin)
			lines := ""

			for {
				line, err := in.ReadString('\n')
				if err == io.EOF {
					break
				}
				if err != nil {
					fmt.Printf("%v", err)
					os.Exit(1)
				}
				lines = lines + line
			}

			cnt = []byte(lines)
		} else {
			fileName := arguments[0]
			cnt, err = os.ReadFile(fileName)
			if err != nil {
				fmt.Printf("%v", err)
				os.Exit(1)
			}
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
		} else if *characterCounter {
			str := string(cnt)
			arr := strings.Split(str, "")

			fmt.Printf("%v %v", len(arr), fileName)
		} else {
			lines := splitter(cnt, "\n")
			words := findCountOfWords(lines)

			fmt.Printf("%v %v %v %v", len(lines), words, len(cnt), fileName)
		}
	}
}

func splitter(text []byte, sep string) []string {
	str := string(text)
	arr := strings.FieldsFunc(str, fun)
	return arr
}

func findCountOfWords(lines []string) int {
	var count int

	for _, line := range lines {
		line = strings.TrimSpace(line)
		arr := strings.Fields(line)
		count = count + len(arr)
	}

	return count
}

func fun(r rune) bool {
	return r == '\n'
}
