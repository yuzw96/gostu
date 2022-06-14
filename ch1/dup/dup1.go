package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//go中的map类似于java中的hashmap，通常使用hash实现
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		//正如for循环一样，if两边也不许要加括号
		if n > 1 {
			//printf可以进行格式转换
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
