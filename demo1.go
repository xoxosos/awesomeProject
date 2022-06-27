package main

import (
	"awesomeProject/hello"
	"fmt"
)

/*
 变量声明多种方法
 var 支持全局访问,也可以声明在方法体类当局部变量
 如果想要在外部包中使用全局变量的首字母必须大写
*/
//var a int =1
var a = 1
var b = "b"
var d, f = "d", 3
var (
	g = 1
	h = "h"
)

//const 用来定义常量
const name = "lrj"

const ( //const与iota搭配可用来定义枚举 ， 第一行iota默认为0 会累加
	beijing = iota + 1
	shanghai
	shenzhen
	guangzhou
)

func main() {
	//	name="a" idea会报错
	c := 3 // ：=声明变量 ，只能用在方法体中外部调用会报错
	fmt.Println("a = ", a)
	fmt.Println("C = ", c)
	fmt.Println("d =", d, "f = ", f)
	fmt.Println("g = ", g, "h = ", h)
	fmt.Println(beijing, shanghai, shenzhen, guangzhou)
	//注意大写开头的方法才能被其他模块调用 返回string
	result := hello.GetOne("我是Hello方法")
	fmt.Println(result)
	//返回多个值
	result1, result2 := hello.GetMore(20, "lrj")
	fmt.Println("result1 = ", result1, "result2 = ", result2)
	//匿名变量 "_" 当只要第一个返回值的时候可以这样声明
	result3, _ := hello.GetMore(30, "xxx")
	//匿名变量不占用内存空间，不会分配内存。匿名变量与匿名变量之间也不会因为多次声明而无法使用。
	fmt.Println("result3=", result3)
	// scan 获取值
	var email string
	fmt.Print("输入邮箱：")
	fmt.Scan(&email)
	fmt.Printf("你的邮箱为：%v\n", email)

}
