package main

import (
	"fmt"
)

func main(){
	var s  string = "hello"
	fmt.Println(s)
	c := []byte(s)
	c[0] = 'c'
	s2 := string(c)
	fmt.Print(s2)
}

