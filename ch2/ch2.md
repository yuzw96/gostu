##程序结构
###命名
go语言中的函数名、变量名、常量名、类型名、语句标号和包名等所有的命名都遵循一个规则：
一个名字必须以一个字符(unicode字母)或者下划线开头，后面可以跟任意数量的字母、数字和下划线。大写字母和小写字母是不同的。

go语言中有25个关键字，关键字是不能用于自定义名字：
>break      default       func     interface   select
> 
>case       defer         go       map         struct
> 
>chan       else          goto     package     switch
> 
>const      fallthrough   if       range       type
> 
>continue   for           import   return      var

此外还有三十多个预定义的名字
>内建常量: true false iota nil
> 
>
>内建类型: int int8 int16 int32 int64
> 
>uint uint8 uint16 uint32 uint64 uintptr
> 
>float32 float64 complex128 complex64
> 
>bool byte rune string error
>
> 
>内建函数: make len cap new append copy close delete
> 
>complex real imag
> 
>panic recover
> 

####名字的定义
如果一个名字在函数内部定义，那它只在函数内部有效。如果在函数外部定义名字，那么将在当前包的所有文件都可以访问。名字开头字母的大小写决定了名字在包外的可见性。如果一个名字是大写字母开头的，那么它将是导出的，也就是说可以被外部的包调用。

###声明
声明语句定义了程序中的各种实体以及部分或者全部的属性。go语言主要有四种类型的声明语句：var、const、type、func对应这变量、常量、类型和函数实体对象的声明。

###变量
声明形式如下：
>var 变量名字 类型 = 表达式
