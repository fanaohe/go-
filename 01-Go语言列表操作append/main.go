package main

import "fmt"

func main()  {
	s1 := []string{"北京","上海","广州"}
	s1[2] = "guang"
	s1 = append(s1, "999") // go 语言中使用append必须使用同变量
	fmt.Println("1111111")
	fmt.Println("1111111")

}
