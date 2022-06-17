package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	var x complex128 = complex(1, 2)
	var y complex128 = complex(3, 4)

	fmt.Println(x * y)
	fmt.Println(real(x * y))
	fmt.Println(imag(x * y))

	//如果一个浮点数面值或一个十进制整数面值后面跟着一个i，例如3.141592i或2i，它将构成一个复数的虚部，复数的实部是0
	fmt.Println(1i * 1i)

	//复数的比较也可以用== !=，但是要考虑浮点数精度问题
	//此外cmplx包提供了复数处理的很多函数
	fmt.Println(cmplx.Sqrt(-1))

}
