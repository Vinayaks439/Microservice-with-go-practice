// Range Keyword

package main

import "fmt"

func main() {
	s := [3]int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}
}
