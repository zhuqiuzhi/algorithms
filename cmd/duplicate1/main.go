package main

import (
	"bufio"
	"fmt"
	"os"
)

// 面试同花顺的后端开发工程师的时候，有道笔试题是要求打印出文件中重复的行
func main() {
	var inFile *os.File
	if len(os.Args) > 1 {
		var err error
		inFile, err = os.Open(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "open %v: %v \n", os.Args[1], err)
			os.Exit(1)
		}
	} else {
		inFile = os.Stdin
	}

	var lines = make(map[string]struct{})
	var line string
	input := bufio.NewScanner(inFile)
	for input.Scan() {
		line = input.Text()
		if _, ok := lines[line]; ok {
			fmt.Println(line)
		} else {
			lines[line] = struct{}{}
		}
	}
}
