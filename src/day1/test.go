package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{0, 1, 3, 2, 4, 5, 7, 6}
	var b = make([]int, 5, 10)
	var c = []int{1, 2, 3, 5, 6}
	fmt.Println(a, cap(a), "\r\n", b, cap(b), "\r\n", c, cap(c))
	// sort.Ints(a) //rese order
	fmt.Println(a)
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println(a)
}

// func main() {
// 	m1 := map[int]string{0: "a", 1: "b", 2: "c", 3: "d", 4: "e"}
// 	m2 := map[string]int{}
// 	for key, value := range m1 {
// 		m2[value] = key
// 		value = "asdasdasdada" // so you can see that the origin map "m1" has not been changed ,so it is a copy ,not a refrence type.
// 	}
// 	fmt.Println(m1)
// 	fmt.Println(m2)

// }
