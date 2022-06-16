package main

import "fmt"

func main() {
	//var u uint8 = 255
	////如下，当计算结果超出边界时，会丢调超出边界的高位
	//fmt.Println(u,u+1,u*u)
	//
	//var i int8 = 127
	//fmt.Println(i,i+1,i*i)

	//位运算
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	//以二进制形式打印该数的八位
	fmt.Printf("%08b\n", x) //00100010
	fmt.Printf("%08b\n", y) //00000110

	fmt.Printf("%08b\n", x&y)
	fmt.Printf("%08b\n", x|y)
	fmt.Printf("%08b\n", x^y)
	fmt.Printf("%08b\n", x&^y)

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i)
		}
	}

	medals := []string{"gold", "silver", "bronze"}

	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i])
	}

}
