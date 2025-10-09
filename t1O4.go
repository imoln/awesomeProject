package main

import "fmt"

func main() {

	var r float32
	const PI = 3.14

	fmt.Println("Radius of the circle: ")
	fmt.Scanln(&r)

	// C = 2 * PI * r
	circumference := 2 * PI * r
	// A = PI * r * r
	area := PI * r * r

	fmt.Println("Circumference of a circle: ", circumference)
	fmt.Println("Area of the circle: ", area)

}
