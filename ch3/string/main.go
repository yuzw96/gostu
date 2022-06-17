package main

import "fmt"

func main() {
	s := "hello world"
	fmt.Println(s[0:5])
	//忽略部分将被默认为0或者len
	fmt.Println(s[:5])
	fmt.Println(s[7:])
	fmt.Println(s[:])

	// +操作符将两个字符串链接构造一个新字符串
	fmt.Println("goodbye" + s[5:])

	//字符串的值是不可变的：
	//一个字符串包含的字节序列永远不会被改变，
	//当然我们也可以给一个字符串变量分配一个新字符串值。
	//可以像下面这样将一个字符串追加到另一个字符串
	ss := "new string"
	tt := ss
	ss += ",change string"
	fmt.Println(ss)
	fmt.Println(tt)

	//因为字符串是不可修改的，因此尝试修改字符串内部数据的操作也是被禁止的
	//ss[0] = 'L'
}
