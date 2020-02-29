package main

import "fmt"

// defer 多用于函数结束之前释放资源，（文件语柄， 数据库连接，socket连接）
func deferDemo()  {
	fmt.Println("start")
	defer fmt.Println("hhh")   // 多个defer语句会按照后进先出的顺序进行执行
	defer fmt.Println("nnnnn")
	fmt.Println("0000")

}

func main()  {
	deferDemo()
}