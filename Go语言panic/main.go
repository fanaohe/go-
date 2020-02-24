package main

import "fmt"

func funcA()  {
	// 在打开数据库资源时，出现意外，防止数据库数据泄露
	defer func() {
		fmt.Println("释放数据库资源")
	}()
	panic("出现了意外")
}
func main()  {
	funcA()
}



