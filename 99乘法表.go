//实现9 9 乘法表
package main

import "fmt"

func Feibonaqi(n int) int  {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n > 1 {
		return Feibonaqi(n-1) + Feibonaqi(n-2)
	} else {
		return -1
	}
}


func main() {
	//遍历, 决定处理第几行
	for y := 1; y <= 9; y++ {
		// 遍历, 决定这一行有多少列
		for x := 1; x <= y; x++ {
			fmt.Printf("%d*%d=%d ", x, y, x*y)
		}
		// 手动生成回车
		fmt.Println()
	}
	//n := 7
	//i := Feibonaqi(n)
	//println(i)

}

