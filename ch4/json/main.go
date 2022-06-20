package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title    string
	Year     int  `json:"released"`
	Color    bool `json:"color,omitempty"`
	Actors   []string
	isImport bool
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1900, Color: false, Actors: []string{"test1", "test1"}, isImport: false},
	{Title: "Cool Hand Luke", Year: 1900, Color: false, Actors: []string{"test2", "test2"}, isImport: false},
	{Title: "Bullitt", Year: 1900, Color: false, Actors: []string{"test3", "test3"}, isImport: false},
}

//这样的数据结构非常适合Json格式，并且两者很容易相互转换
//在go语言中，将类似movies的结构体slice转为JSON的过程叫做编组(marshaling)
func main() {
	//data,err := json.Marshal(movies)
	//在编码过程中默认使用Go语言结构体的成员名字作为JSON对象，需要注意的是只有导出的结构体成员从才会被编码
	//除了Marshal方法，我们为了便于阅读，也可以使用一下方法来生成缩进格式的数据
	data, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	//我们可以从打印结果看到未导出的结构体成员isImport并没有被编码
	//此外，由于我们在申明成员Year、Color时指定了编码后的对象名称，因此打印出的json成员Year的名称变为了released
	fmt.Printf("%s\n", data)

	//与编码先对应的解码如下
	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	//通过打印可以看到在解码的时候我们只会解码结构中已有的成员
	fmt.Println(titles)
}
