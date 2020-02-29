package main

import "fmt"

// map 操作

func main()  {
	var m1 map[string]int   //使用map时必须进行初始化
	fmt.Println(m1)
	m1 = make(map[string]int, 10)
	m1["理想"] = 19
	m1["年龄"] = 200
	// 在go语言中map类似与python中的字典，在取值时map类型返回两个值   value 和 bool
	value, ok :=m1["理想"]
	fmt.Println(value)
	fmt.Println(ok)

	// map遍历可以使用for range
	for k, v := range m1{
		fmt.Println(k, v)
	}
	// 在go语言中只遍历k
	for k :=range m1{
		fmt.Println(k)
	}
	// 在go语言中只遍历v
	for _, bools :=range m1{
		fmt.Println(bools)
	}

	// go语言中字典中有列表
	var s1 = make([]map[int]string, 10, 10)    //第一个10指几个，第二个10指容量
	fmt.Println(s1[0])
	s1[0] = make(map[int]string, 1)
	s1[0][10] = "沙河"     // go语言键值对添加操作，10为键，沙河为值
	fmt.Println(s1)
	// 值为切片类型的map
	var s2 = make(map[string][]int, 10)
	s2["北京"] = []int{10, 20, 30}
	fmt.Println(s2)


}

