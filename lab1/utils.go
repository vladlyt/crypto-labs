package main

import (
	"bufio"
	"os"
	"strings"
	"unicode"
)

func ClearText(path string, outPath string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	out, err := os.Create(outPath)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		res := ""
		for _, c := range scanner.Text() {
			if unicode.IsLetter(c) {
				res += strings.ToUpper(string(c))
			}
		}
		_, err := out.WriteString(res)
		if err != nil {
			panic(err)
		}
	}
}

func GetTextFromFile(path string) string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	res := ""
	for {
		line, isPrefix, _ := reader.ReadLine()
		res += string(line)
		if !isPrefix {
			break
		}
	}
	return res
}
