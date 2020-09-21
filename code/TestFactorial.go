package main

import "fmt"

func factorial(n int) int {
	if n==1{
		return 1
	}else {
		return factorial(n-1)*n//递归地调用自己，直到返回结束
	}
}

func main() {
	var n=5
	var fac =factorial(n)
	fmt.Printf("5的阶乘是：%d",fac)
}