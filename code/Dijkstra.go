package main

import "fmt"

func main() {
	var flags = []string {"R" , "B" , "W" , "B" , "R" , "W" }
	var length = len (flags)
	var b = 0
	var w = 0
	var r = length - 1
	var count = 0
	for {
		if w > r {
			break
		}
		if flags[w] == "W" {
			w++
		} else if flags[w] == "B" {
			var temp = flags[w]
			flags[w] = flags[b]
			flags[b] = temp
			w++
			b++
			count++
		} else if flags[w] == "R" {
			var m = flags[w]
			flags[w] = flags[r]
			flags[r] = m
			r--
			count++
		}
	}
	for i := 0 ; i < length; i++ {
		fmt.Printf("%s" , flags[i])
	}
	fmt.Printf("\n总交换计数  : %d" , count)
}
