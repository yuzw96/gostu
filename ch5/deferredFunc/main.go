package main

import (
	"log"
	"sync"
	"time"
)

//随着处理逻辑的复杂，我们很难去保证清晰的业务逻辑,这时就可以用deferred机制来简化处理
//仅需在调用函数或方法前加上关键字，就完成了defer所需的语法
//当defer语句被执行的时候，跟在defer后面的函数会被延时执行。直到包含该defer语句的函数执行完毕时，defer后的函数才会执行
//我们也可以在一个函数中执行多条defer语句，它们的执行顺序和声明顺序相反
func main() {
	bigSlowOperation()
}

//如下用在释放锁时
var mu sync.Mutex
var m = make(map[string]int)

func lookup(key string) int {
	mu.Lock()
	defer mu.Unlock()
	return m[key]
}

//调试复杂程序时，defer也被用来记录何时进出函数
//我们来观察下面的例子

func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	//extra parentheses
	time.Sleep(10 * time.Second)
	//operation by sleeping
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

//我们再来具体分析一下几个例子

func f(res int) int {

	defer func() {
		res++
	}()
	return 0

	//我们可以将上述代码改写成如下
	//res = 0   return语句不是一条原子调用，return xxx其实是赋值+ret指令
	//func(){	defer被插入到return之前执行
	//	res++
	//}()
	//
	//return 因此最终的返回值为1
}

func f2(res int) int {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t

	//为了理解将其该写如下
	//t:=5
	//r = t  赋值命令
	//func() {  defer被插入到赋值与返回之间执行，这个例子中返回r没被修改过
	//	t = t+5
	//}
	//return  空的return指令 返回结果还是5
}

func f3(res int) int {
	defer func(res int) {
		res = res + 5
	}(res)

	return 1

	//res = 1 给返回值赋值
	//func(res int) { 这里修改的r是传值传进去的r，不会改变要返回的那个r值
	//	r = r+5
	//}(r)
	//
	//return 空的return
}
