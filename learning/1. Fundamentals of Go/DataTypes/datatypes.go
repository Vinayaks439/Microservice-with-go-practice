package main

import "fmt"

func Pointer() {
	var i int = 8
	var pi *int
	pi = &i
	*pi = 5
	fmt.Println(i)
}

type Circle struct {
	X, Y, R int
}

func Structs() {
	c1 := Circle{2, 3, 5}
	c2 := Circle{Y: 1, R: 5}
	c3 := Circle{}

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
}

func Slices() {
	s1 := make([]int, 5)
	s2 := make([]int, 0, 5)

	s1 = append(s1, 1)
	s2 = append(s2, 2)

	fmt.Println(len(s1), cap(s1))
	fmt.Println(len(s2), cap(s2))
}

func Maps() {
	m := make(map[int]string)

	m[1] = "Monday"
	m[3] = "Wednesday"

	value, ok := m[1]

	fmt.Println(value, ok)

	value1, ok1 := m[2]

	fmt.Println(value1, ok1)
}

func main() {
	fmt.Println("Executing Pointer code...")
	Pointer()
	fmt.Println("Executing Structs code...")
	Structs()
	fmt.Println("Executing Slices code...")
	Slices()
	fmt.Println("Executing Maps code...")
	Maps()
}
