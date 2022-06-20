package main

import "fmt"

func main() {
	//创建
	ages := map[string]int{
		"tony": 11,
		"jhon": 22,
	}

	fmt.Println("初始化结束：", ages)

	//添加
	ages["alice"] = 33

	fmt.Println("添加元素：", ages)

	//删除
	delete(ages, "alice")

	fmt.Println("删除元素：", ages)

	//以下代码是可以正常运行的，当我们查询一个不存在的值时，会默认返回该类型的零值
	ages["bob"] = ages["bob"] + 1

	fmt.Println("添加未知元素：", ages)

}
