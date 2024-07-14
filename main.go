package main

import (
	"fmt"
	"log"
	"os"
)

func getFlagsAndParams(args []string) ([]string, []string) {
	flags := []string{}
	params := []string{}
	for _, arg := range args {
		if len(arg) > 0 {
			if arg[0] == '-' {
				flags = append(flags, arg)
			} else {
				params = append(params, arg)
			}
		}
	}
	return flags, params
}

func countBytes(content []byte) int {
	return len(content)
}
func countWords(content []byte) int {
	return len(content)
}
func countLines(content []byte) int {
	return len(content)
}
func countChars(content []byte) int {
	runes := []rune(string(content))
	return len(runes)
}

func main() {
	flags, files := getFlagsAndParams(os.Args[1:])

	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			log.Fatal("error reading file: ", err)
		}
		if len(flags) == 0 {
			fmt.Printf("%v\t%v\t%v\t%v\n", countBytes(content), countWords(content), countLines(content), file)
			continue
		}
		output := ""
		for _, flag := range flags {
			if flag == "-c" {
				output = fmt.Sprint(output, countBytes(content), "\t")
				continue
			}
			if flag == "-w" {
				output = fmt.Sprint(output, countWords(content), "\t")
				continue
			}
			if flag == "-l" {
				output = fmt.Sprint(output, countLines(content), "\t")
				continue
			}
			if flag == "-m" {
				output = fmt.Sprint(output, countChars(content), "\t")
				continue
			}
			// if flag doesn't match any known flag
			log.Fatal("unknown flag: ", flag)
		}
		output = output + file + "\n"
		fmt.Print(output)
	}

}
