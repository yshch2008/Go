# Slice

```golang

func main() {
	a := []int{0, 1, 3, 2, 4, 5, 7, 6}
	var b = make([]int, 5, 10)
	var c = []int{1, 2, 3, 5, 6}
	fmt.Println(a, cap(a), "\r\n", b, cap(b), "\r\n", c, cap(c))
	// sort.Ints(a) //rese order
	fmt.Println(a)
	sort.Sort(sort.Reverse(sort.IntSlice(a)))//desending,looks so stupid
	fmt.Println(a)
}
```