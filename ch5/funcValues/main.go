package main

import (
	"fmt"
	"strings"
)

func main() {
	//将函数赋值给变量
	f := square
	fmt.Println(f(3))

	f = negative
	fmt.Println(f(3))
	fmt.Printf("%T\n", f)

	//Cannot use 'product' (type func(n int, m int) int) as the type func(n int) int
	//f = product

	//函数类型的零值为nil，调用值为nil的函数值会引起panic错误
	//var f1 func(int) int
	//f1(3)

	//函数值使得我们不仅可以通过数据来参数化函数，亦可通过行为，如下
	fmt.Println(strings.Map(add1, "HAL-9000"))
	fmt.Println(strings.Map(add1, "VMS"))
	fmt.Println(strings.Map(add1, "Admix"))

}

//另外函数值之间是不能比较的，也不能用函数值作为map的key
func add1(r rune) rune {
	return r + 1
}

func square(n int) int {
	return n * n
}

func negative(n int) int {
	return -n
}

func product(n, m int) int {
	return m * n
}
