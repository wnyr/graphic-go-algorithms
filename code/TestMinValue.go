package main

import "fmt"

func min(arrays []int, length int) int {
	var minIndex=0 //最小索引
	for j := 0; j < length; j++ {
		if arrays[minIndex]>arrays[j]{
			minIndex=j
		}
	}
	return arrays[minIndex]
}
func main() {
	var scores=[]int{60,80,95,50,70}
	var length=len(scores)
	var minValue=min(scores,length)
	fmt.Printf("Min Value= %d\n",minValue)
}