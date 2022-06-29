package main

import "fmt"

func main() {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

//通过这种方式定义的函数可以访问完整的词法环境，这就意味着在函数中定义的内部函数可以引用该函数的变量
func squares() func() int {
	var x int
	return func() int {
		x++
		return x + x
	}
}
