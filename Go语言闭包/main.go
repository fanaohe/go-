package main

import (
	"fmt"
)

func f1(x, y int) {
	fmt.Println("this is f1")
	fmt.Println(x + y)
}

func lixiang(x func(int, int), m, n int) int {
	tmp := func() {
		x(m, n)
	}
	tmp()

	return 1

}

func main()  {
	a := lixiang(f1, 100, 200)
	fmt.Println(a)
	fmt.Printf("%T", f1)
}


