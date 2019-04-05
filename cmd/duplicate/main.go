package main

import (
	"fmt"
	"os"
)

// 面试同花顺的后端开发工程师的时候，有道笔试题是要求打印出文件中重复的行
func main() {
	inFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "open %v: %v \n", os.Args[1], err)
		os.Exit(1)
	}

	var lines = make(map[string]struct{})
	var line string
	for {
		// 注意 Fscanf 的第三个参数得是
		if _, err := fmt.Fscanf(inFile, "%s\n", &line); err != nil {
			return
		} else {
			if _, ok := lines[line]; ok {
				fmt.Println(line)
			} else {
				lines[line] = struct{}{}
			}
		}
	}
}
