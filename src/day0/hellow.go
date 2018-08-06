package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// for value := range feib() {
	// 	fmt.Println(value)
	// }
	// fmt.Println(feib())

	// a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	// sa := a[2:7]
	// sa = append(sa, 100)
	// sb := sa[3:8]
	// sb[0] = 99
	// fmt.Println(a)  //输出：[1 2 3 4 5 99 7 100 9 0]
	// fmt.Println(sa) //输出：[3 4 5 99 7 100]
	// fmt.Println(sb)
	// fmt.Println(feib())
	// fmt.Println(fb2(15))

	outDate()
}

func feib() []int {
	result := []int{}
	var pre, value = 1, 1
	for a := 0; a < 15; a++ {
		if a < 1 {
		} else {
			pre += value
			value += pre
		}
		result = append(result, pre, value)
	}
	return result
}

func fb2(count int) []int {
	result := []int{}
	for i := 0; i < count; i++ {
		if i > 1 {
			result = append(result, result[i-2]+result[i-1])
		} else {
			result = append(result, 1)
		}
	}
	return result
}

func outDate() {
	today := time.Now()
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("I am running at a linux PC.")
	case "darwin":
		fmt.Println("OS X.")
	case "windows":
		fmt.Println("windows ,the best os.")
		fallthrough
	default:
		fmt.Println(today.String())
	}
}
