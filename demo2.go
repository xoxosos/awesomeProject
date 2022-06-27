package main

import (
	"fmt"
)

func main() {
	//声明变量 a,b
	a, b := 2, "b"
	// 用&符号 获得变量a,b的内存地址 ，得到对应类型变量的指针
	ptr1, ptr2 := &a, &b
	// 打印对应指针类型与内存地址
	fmt.Printf("ptr1 type: %T\n ptr2 type: %T\n", ptr1, ptr2)
	fmt.Printf("ptr1  %p\n ptr2  %p\n", ptr1, ptr2)
	// *是解指针符号 得到对应变量的值
	value1, value2 := *ptr1, *ptr2
	fmt.Printf("ptr1 value:%v\n ptr2 value:%v\n", value1, value2)
	//也可以操作指针来改变值
	*ptr1 += 2
	// 2+2输出4
	fmt.Println(a)
	//new() 函数可以创建一个对应类型的指针，创建过程会分配内存，被创建的指针指向默认值。
	str := new(string)
	*str = "hello"
	fmt.Println(*str)
	// 多重指针 指针可以指向任何类型的变量。所以也可以指向另一个指针
	p := 3.14
	//得到变量p指针
	pp := &p
	//得到变量pp指针
	ppp := &pp
	fmt.Println("p =", p)
	fmt.Println("pp =", pp)
	fmt.Println("ppp =", ppp)
	fmt.Println("*pp =", *pp)
	fmt.Println("*ppp =", *ppp)
	fmt.Println("**ppp =", **ppp)
	//==运算符来比较两个相同类型的指针是否相等
	if !(pp != *ppp) {
		fmt.Println("两个相同类型的指针不相等")
	}

}
