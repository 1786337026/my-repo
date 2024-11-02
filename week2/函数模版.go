package main

import "fmt"

func function_swag(a int, b int) (int, int) { // 交换函数
	return b, a
}

func main() {
	var x, y int
	fmt.Scanf("%d%d", &x, &y)
	fmt.Printf("%d %d\n", x, y)
	x, y = function_swag(x, y)
	fmt.Printf("%d %d", x, y)
}
