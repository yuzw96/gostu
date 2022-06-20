package main

import "fmt"

func main() {
	var i [3]string
	fmt.Println("字符串类型的零值为:" + i[0] + ".")

	//简化的初始化写法
	j := [...]int{1, 2, 3, 4}
	fmt.Println(j)

	//我们也可以在初始化的时候只指定特定元素的初始值
	//如下，我们将该数组的最后一个元素指定为-1，而其它元素都为默认的零值
	r := [...]int{99: -1}
	fmt.Println(r[1])
	fmt.Println(r[len(r)-1])
}
