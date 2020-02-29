package main

import "fmt"

// 标识符: 变量名，函数名，类型名，方法名
// Go语言中如果标识符首字母是大写的，就表示对外部包可见(暴露的，公有的)，另一个文件进行导入时可以使用


type Dog struct {
	name string
}

// 构造函数
func newDog(name string) Dog  {
	return Dog{
		name:name,
	}
}

// 方法是作用于特定类型的
// 接收者表示的是调用该方法的具体类型变量，多用类型名首字母小写表示
func (d Dog)wang()  {     // 在函数名称前面加上括号，表示可以约束当前函数那些属性可以调用，当前是dog类型可以调用
	fmt.Printf("%s:汪汪汪\n", d.name)
}

func main()  {
	d1 := newDog("张凉")
	d1.wang()
}