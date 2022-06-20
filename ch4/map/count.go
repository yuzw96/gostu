package main

import (
	"bufio"
	"fmt"
	"os"
)

//统计计数
func main() {
	wordcount := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		wordcount[input.Text()]++
	}

	for word, count := range wordcount {
		fmt.Printf("word:%s,count=%d\n", word, count)
	}
}
