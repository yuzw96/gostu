package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(1, 3, 5, 7, 9))
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func max(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, errors.New("请保证至少有一个参数")
	}
	max := vals[0]
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max, nil
}

func min(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, errors.New("请保证至少有一个参数")
	}
	min := vals[0]
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min, nil
}

func mustMax(first int, vals ...int) (int, error) {
	max := first
	if len(vals) > 0 {
		for _, val := range vals {
			if val > max {
				max = val
			}
		}
	}
	return max, nil
}

func mustMin(first int, vals ...int) (int, error) {
	min := first
	if len(vals) > 0 {
		for _, val := range vals {
			if val < min {
				min = val
			}
		}
	}
	return min, nil
}
