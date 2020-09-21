package main

import "fmt"

func swap(arrays []int,a int,b int)  {
	arrays[a]=arrays[a]+arrays[b]
	arrays[b]=arrays[a]-arrays[b]
	arrays[a]=arrays[a]-arrays[b]
}

func shellSort(array []int,length int)  {
	for gap := length/2; gap >0 ; gap=gap/2 {
		for i := gap; i < length; i++ {
			var j=i
			for{
				if j-gap<0 || array[j]>=array[j-gap]{
					break
				}
				swap(array,j,j-gap)
				j=j-gap
			}
		}
	}
}

func main() {
	var scores = []int {9 , 6 , 5 , 8 , 0 , 7 , 4 , 3 , 1 , 2 }
	var length = len (scores)
	shellSort(scores,length)
	for i := 0; i < length; i++ {
		fmt.Printf("%d,",scores[i])
	}
}