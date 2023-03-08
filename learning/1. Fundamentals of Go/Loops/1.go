package main

import "fmt"

1.
// func main() {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(i)
// 	}
// }

2. 
// func main() {
// 	i := 0
// 	for ; i < 5; i++ {
// 		fmt.Println(i)
// 	}
// }

3.
// func main() {
// 	i := 0
// 	for ; i < 5; i++ {
// 		fmt.Println(i)
// 	}
// 	fmt.Println(i)
// }

4.
// func main() {
// 	i := 0
// 	for i < 5 {
// 		fmt.Println(i)
// 		i++
// 	}
// }

5. 
func main() {
	i := 0
	for {
		fmt.Println(i)
		i++
		if i == 5 {
			break
		}
	}
}
