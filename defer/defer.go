package main

import "fmt"

/*
  defer 类似finally 程序结束执行
*/
func fun1() {
	fmt.Println("main::fun1")
}
func fun2() {
	fmt.Println("main::fun2")
}

func fun3() {
	fmt.Println("main::fun3")
}
func deferfunc() int {
	fmt.Println("defferfunc...")
	return 0
}
func returnfunc() int {
	fmt.Println("returnfunc...")
	return 0
}
func deferandreturn() int {
	defer deferfunc()
	return returnfunc()
}

// defer执行顺序 先进后出
func main() {
	fmt.Println("main1")
	defer fun1()
	defer fun2()
	defer fun3()
	fmt.Println("main2")
	// return先执行 , defer后执行
	deferandreturn()
}
