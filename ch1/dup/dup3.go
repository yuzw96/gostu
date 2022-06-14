package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		//ReadFile函数返回一个字节切片,必须把他转换成string才能进行strings.Split
		data, error := ioutil.ReadFile(filename)
		if error != nil {
			fmt.Fprintf(os.Stderr, "dups: %v\n", error)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
