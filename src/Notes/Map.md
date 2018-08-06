# Map

## grammar

``` go
var m=make(map[string] int)
m:=make(map[string] int)
m:=map[string] int{  }
var m map[string] int=make(map[string] int)

for key,value range m{
    fmt.Println(value)
}

func main() {
	m1 := map[int]string{0: "a", 1: "b", 2: "c", 3: "d", 4: "e"}
	m2 := map[string]int{}
	for key, value := range m1 {
		m2[value] = key
		value = "asdasdasdada" // so you can see that the origin map "m1" has not been changed ,so it is a copy ,not a refrence type.
	}
	fmt.Println(m1)
	fmt.Println(m2)

}
```

## speciality

iterator will make a copy of origin entity instead of use the key-value entity directly.

value 's type can be changed differently from original declaration.

meke() will just make the outermost entity ,like other method.
