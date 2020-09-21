package main

import "fmt"

//调整堆
func adjustHeap(array []int , currentIndex int , maxLength int ) {
	var noLeafValue = array[currentIndex] // 目前非叶节点
	//2 * currentIndex + 1   当前左子树下标
	for j := 2 *currentIndex + 1 ; j <= maxLength; j = currentIndex*2 + 1 {
		if j < maxLength && array[j] < array[j+1 ] {
			j++ // j 大的下标
		}
		if noLeafValue >= array[j] {
			break
		}
		array[currentIndex] = array[j] // 向上移动到父节点
		currentIndex = j
	}
	array[currentIndex] = noLeafValue // 放到位置上
}

//初始化堆
func createHeap(array []int , length int ) {
	// 建立一个堆, (length - 1) / 2 扫描一半的子节点
	for i := (length - 1 ) / 2 ; i >= 0 ; i-- {
		adjustHeap(array, i, length-1 )
	}
}

func heapSort(array []int , length int ) {
	for i := length - 1 ; i > 0 ; i-- {
		var temp = array[0 ]
		array[0 ] = array[i]
		array[i] = temp
		adjustHeap(array, 0 , i-1 )
	}
}

func main() {
	var scores = []int {10 , 90 , 20 , 80 , 30 , 70 , 40 , 60 , 50 }
	var length = len (scores)
	fmt.Printf("在建立堆之前 : \n" )
	for i := 0 ; i < length; i++ {
		fmt.Printf("%d, " , scores[i])
	}
	fmt.Printf("\n\n" )
	fmt.Printf("在建立堆之后 : \n" )
	createHeap(scores, length)
	for i := 0 ; i < length; i++ {
		fmt.Printf("%d, " , scores[i])
	}
	fmt.Printf("\n\n" )
	fmt.Printf("堆排序 : \n" )
	heapSort(scores, length)
	for i := 0 ; i < length; i++ {
		fmt.Printf("%d, " , scores[i])
	}
}
