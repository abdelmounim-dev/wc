package main

import (
	"bufio"
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
	chars := []rune(string(content))
	count := 0
	isBlankChar := func(c rune) bool {
		return c == '\n' || c == ' ' || c == '\t' || c == '\r' || c == '\f' || c == '\v'
	}
	inWord := false
	for i := 0; i < len(chars); i++ {
		if isBlankChar(chars[i]) {
			inWord = false
		} else if !inWord {
			inWord = true
			count++
		}
	}
	return count
}

func countLines(content []byte) int {
	str := string(content)
	count := 0
	for _, c := range str {
		if c == '\n' {
			count++
		}
	}
	return count
}

func countChars(content []byte) int {
	runes := []rune(string(content))
	return len(runes)
}

func count(content []byte, flags []string) (string, error) {

	var output string
	if len(flags) == 0 {
		output = fmt.Sprintf("  %v\t%v\t%v\t", countLines(content), countWords(content), countBytes(content))
		return output, nil
	}
	for _, flag := range flags {
		if flag == "-l" {
			output = fmt.Sprint(output, countLines(content), "\t")
			continue
		}
		if flag == "-w" {
			output = fmt.Sprint(output, countWords(content), "\t")
			continue
		}
		if flag == "-c" {
			output = fmt.Sprint(output, countBytes(content), "\t")
			continue
		}
		if flag == "-m" {
			output = fmt.Sprint(output, countChars(content), "\t")
			continue
		}
		return "", fmt.Errorf("unknown flag: %s", flag)
	}
	return output, nil
}

func main() {
	flags, files := getFlagsAndParams(os.Args[1:])
	if len(files) == 0 {
		content := []byte{}
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			content = append(content, scanner.Bytes()...)
			content = append(content, byte('\n'))
		}
		if err := scanner.Err(); err != nil {
			log.Fatal("error reading from stdin: ", err)
		}
		output, err := count(content, flags)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", output)
	}

	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			log.Fatal("error reading file: ", err)
		}
		output, err := count(content, flags)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\t%s\n", output, file)
	}

}
