### go基本命令
go run xxx.go 运行

go build xxx.go 构建，并生成相应的exe文件

go get xxx/xxx 获取指定地址的代码

###go代码管理
go语言的代码是由包(package)来确定的，类似与其他语言中的库或者模块。每个源代码都由一条package的语句开始。

main包是一个特殊的存在，它定义了一个独立可执行的程序，而不是一个库。main包中的main函数也比较特殊，它是整个程序执行时的入口。

每个源代码中通过import的形式来申明用到的包，且必须是恰到好处的引用，缺少必要的包或者映入非必要的包都会导致程序无法编译。

###命令行参数
在go的标准库中os库以跨平台的方式，提供了一些与操作系统交互的函数和变量

os.Args变量是一个字符串（string）的切面（slice-动态数组），其中切片是go的一个基础概念

###go中的循环
go中只有for循环一种循环，具体有一下几种形式

``` go
//最常见的一种形式
for initialization;condition;post {
    //dosomething
} 
```

``` go
//相当于while循环
for condition {
    //dosomething
        }
```

``` go
//也就是一个无限循环，我们可以使用break或者return来使其停止执行
for {
    //dosomething
}

```

除此之外，for循环还有一种形式，也就是的数据的区间（range）上进行遍历，[详见代码echo2]