package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	var student1 string = "John Doe"
	var student2 = "Jane Doe"
	x := 2

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	var v, l, n, m int = 1, 2, 3, 4
	fmt.Println(v)
	fmt.Println(l)
	fmt.Println(n)
	fmt.Println(m)

	var (
		aa int
		bb int    = 2
		cc string = "hi"
	)

	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)

	const PI = 3.14

	fmt.Println(PI)

	var arr1 = [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr1, arr2)

	var arr3 = [3]string{"a", "b", "c"}

	fmt.Println(arr3)

}
