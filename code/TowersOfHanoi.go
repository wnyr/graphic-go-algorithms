package main

import "fmt"

func hanoi(n int , A string , B string , C string ) {
	if n == 1 {
		fmt.Printf("移动 %d 号 %s 到 %s \n" , n, A, C)
	} else {
		hanoi(n-1 , A, C, B) //移动A上的第n-1个圆盘    通过C到B
		fmt.Printf("移动 %d 号从 %s to %s \n" , n, A, C)
		hanoi(n-1 , B, A, C) //移动B上的第n-1个圆盘    通过A到C
	}
}

func main() {
	fmt.Printf("请输入盘的数量 : \n" )
	var n int
	fmt.Scanf("%d" , &n)
	hanoi(n, "A" , "B" , "C" )
}