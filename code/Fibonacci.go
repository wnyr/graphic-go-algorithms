package main

import "fmt"

func fibonacci(n int ) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fibonacci(n-1 ) + fibonacci(n-2 )
	}
}
func main() {
	fmt.Printf("请输入月数 : \n" )
	var number int
	fmt.Scanf("%d" , &number)
	for i := 1 ; i <= number; i++ {
		fmt.Printf("%d 个月: %d \n" , i, fibonacci(i))
	}
}
