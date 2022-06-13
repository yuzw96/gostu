package main

import (
	"fmt"
	"os"
	"strings"
)

//这里go语言使用range来获取参数，我们不能通过把索引赋值给一个零时变量然后再忽略它的值的方式来进行遍历，因为go是不允许使用无用的局部变量的
//相应的go语言使用空标识符（_）来表示
//另外我们也可以采用实例中给变量赋值的方式来进行初始化
func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		//当遍历的参数很多时，+=的方式代价将会变高，因为在这个过程中会产生很多新的字符串需要等待gc
		//因此，对于字符串的拼接我们可以使用strings包中的join函数
		s += sep + arg
		sep = ""
	}
	fmt.Println(s)
	//join用法如下
	fmt.Println(strings.Join(os.Args[1:], " "))

	//循环打印每一对键值
	for k, v := range os.Args[1:] {
		fmt.Println(k, v)
	}
}
