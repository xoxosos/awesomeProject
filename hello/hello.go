package hello

import (
	"fmt"
)

func GetOne(a string) string {
	fmt.Println("返回输入的值：", a)
	return a
}

func GetMore(age int, name string) (a int, b string) {
	fmt.Println("a =", a, "b =", b)
	return age, name
}
