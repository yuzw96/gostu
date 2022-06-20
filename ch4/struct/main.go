package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID int
	//在创建结构体的时候，我们也可以把数据类型相同的成员合并放到一行
	//例：Name,Position string
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var dilbert Employee

func main() {
	fmt.Println("新建实例:", dilbert)

	//通过点操作符来访问结构体的具体成员
	dilbert.Salary -= 5000
	fmt.Println("访问成员并更新:", dilbert)

	//也可以对成员取地址，然后通过指针访问
	position := &dilbert.Position
	*position = "Senior " + *position
	fmt.Println("指针访问:" + *position)
	fmt.Println(dilbert)

	//点操作符也可以和指向结构体的指针一起工作
	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += "(proactive team player)"
	fmt.Println("点操作符与指针一起工作:", employeeOfTheMonth)
}
