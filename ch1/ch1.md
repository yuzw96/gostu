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

###Printf动词
>  %d                十进制整数 
>         
>  %x, %o, %b        十六进制，八进制，二进制整数。
>         
>  %f, %g, %e        浮点数： 3.141593 3.141592653589793 3.141593e+00
>         
>  %t                布尔：true或false
>         
>  %c                字符（rune） (Unicode码点)
>         
>  %s                字符串
>         
>  %q                带双引号的字符串"abc"或带单引号的字符'c'
>         
>  %v                变量的自然形式（natural format）
>         
>  %T                变量的类型
>         
>  %%                字面上的百分号标志（无操作数）


###switch
应用demo如下

``` go
//go语言不会显式的在每个指向语句后面加break，
//语言默认在执行case后面的逻辑语句就会退出，
//如果你想要相邻的几个case执行同一逻辑就要显式的写上fallthrough来覆盖这种默认行为
switch coinflip() {
case: "heads"
    heads++
case: "tails"
    tails++
default:
    fmt.Println("mismatch")
}
```

switch除此之外还有一种形式
``` go
//这种叫无tag switch和switch true相同
func Signum(x int) int{
    switch{
    case x>0:
        return +1
    default:    
        return 0
    case x<0:
        return -1
    }
}
```

###命名类型
类型声明使得我们很方便地给一个特殊类型一个名字
``` go
type Point struct{
    X,Y int
}

var p Point
```

###指针
go语言提供了指针，是一种存储了变量内存地址的数据类型，且指针操作完全不受约束

###方法和接口
方法和命名类型关联的一类函数，go语言中比较特殊的是方法可以被给关联到任意一种命名类型

###包
go目前已经有了很多的标准库，方面我们去引用到自己的项目[go标准库查询](https://pkg.go.dev/)