package main

import (
	"fmt"
	"math"
)

func main() {
	var f float32 = 16777216
	var f64 float64 = 16777216
	//因为float32类型的累计计算误差很容易扩散，并且float32能精确表示的正整数并不是很大
	//因为float32的有效bit位只有23个，其它的bit位用于指数和符号；当整数大于23bit能表达的范围时，float32的表示将出现误差
	fmt.Println(f)
	fmt.Println(f + 1)
	fmt.Println(f == f+1)

	fmt.Println(f64)
	fmt.Println(f64 + 1)
	fmt.Println(f64 == f64+1)

	//小数点前面或后面的数字都可能被省略（例如.707或1.）。很小或很大的数最好用科学计数法书写，通过e或E来指定指数部分
	const e = 2.71828
	const Avogadro = 6.02214129e23 // 阿伏伽德罗常数
	const Planck = 6.62606957e-34  // 普朗克常数

	//用Printf函数的%g参数打印浮点数，将采用更紧凑的表示形式打印，并提供足够的精度，
	//但是对应表格的数据，使用%e（带指数）或%f的形式打印可能更合适。
	//所有的这三个打印形式都可以指定打印的宽度和控制打印精度
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) //0 -0 +Inf -Inf NaN

	man := math.NaN()
	fmt.Println(man == man, man < man, man > man) //false false false
}
