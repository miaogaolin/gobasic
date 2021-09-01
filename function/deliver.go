package main

// callback 是一个函数类型
func Calc(callback func(n1 int, n2 int) int) int {
	x, y := 3, 4
	return callback(x, y)
}

// 计算两数之积
func Mul(n1 int, n2 int) int {
	return n1 * n2
}

func main() {
	// 第一个：传递一个匿名函数
	Calc(func(n1 int, n2 int) int {
		return n1 + n2
	})

	// 第二个：传递一个定义好的函数
	Calc(Mul)
}
