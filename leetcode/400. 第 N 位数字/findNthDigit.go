package main

import (
	"fmt"
	"math"
)

func findNthDigit(n int) int {
	// 一位的数字有9个，两位的有90个，三位的有900个，以此类推
	// 先求出n所在的区间，如n=13在10-99区间内
	// 然后再锁定是哪个数字的第几位
	// num 所选数字     pos 所选数字的第pos位   i 该区间的数字有几位
	i := 1
	for ; n > i*9*int(math.Pow(10, float64(i-1))); i++ {
		n -= i * 9 * int(math.Pow(10, float64(i-1)))
	}
	pos := n % i
	num := 0
	pow := int(math.Pow(10, float64(i-1)))
	if pos == 0 {
		num = pow + n/i - 1
		pos = i
	} else {
		num = pow + n/i
	}

	// 将pos反转
	pos = i - pos + 1
	ret := 0
	for pos > 0 {
		ret = num % 10
		num /= 10
		pos--
	}
	return ret
}

func main() {
	ret := findNthDigit(13)
	fmt.Println(ret)
}
