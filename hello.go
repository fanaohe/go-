// package main

// import "fmt"

// func main() {
// 	fmt.Printf("Hello word")
// }
package main

import "fmt"

func swap(x string, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}



