package main

import (
	"fmt"
)

func main() {
	var scores = []int {90 , 70 , 50 , 80 , 60 , 85 }
	fmt.Printf("请输入要删除的索引:\n")
	var index int
	fmt.Scan(&index)

	var length=len(scores)
	var tempArray=make([]int,length-1)//创建临时数组
	for i := 0; i < length; i++ {
		if i<index {//将索引前面的数据复制到tempArray的前面
			tempArray[i]=scores[i]
		}else {//将索引后的数组复制到tempArray的后面
			tempArray[i-1]=scores[i]
		}
	}

	scores=tempArray
	for i := 0; i < length-1; i++ {
		fmt.Printf("%d,",scores[i])
	}
}
