package main

import (
	"fmt"
	"math/rand"
)

func compare(coins []int , i int , j int , k int ) { //coin[k] true,coin[i]>coin[j]
	if coins[i] > coins[k] { //coin[i]>coin[j]&&coin[i]>coin[k] ----->coin[i] 是一枚重的假币
		fmt.Printf("\n假币 %d 较重\n" , (i + 1 ))
	} else { //coin[j] is a light counterfeit coin
		fmt.Printf("\n假币 %d 较轻\n " , (j + 1 ))
	}
}

func eightcoins(coins []int ) {
	if coins[0 ]+coins[1 ]+coins[2 ] == coins[3 ]+coins[4 ]+coins[5 ] { //(a+b+c)==(d+e+f)
		if coins[6 ] > coins[7 ] { //g>h?(g>a?g:a):(h>a?h:a)
			compare(coins, 6 , 7 , 0 )
		} else { //h>g?(h>a?h:a):(g>a?g:a)
			compare(coins, 7 , 6 , 0 )
		}
	} else if coins[0 ]+coins[1 ]+coins[2 ] > coins[3 ]+coins[4 ]+coins[5 ] { //(a+b+c)>(d+e+f)
		if coins[0 ]+coins[3 ] == coins[1 ]+coins[4 ] { //(a+e)==(d+b)
			compare(coins, 2 , 5 , 0 )
		} else if coins[0 ]+coins[3 ] > coins[1 ]+coins[4 ] { //(a+e)>(d+b)
			compare(coins, 0 , 4 , 1 )
		}
		if coins[0 ]+coins[3 ] < coins[1 ]+coins[4 ] { //(a+e)<(d+b)
			compare(coins, 1 , 3 , 0 )
		}
	} else if coins[0 ]+coins[1 ]+coins[2 ] < coins[3 ]+coins[4 ]+coins[5 ] { //(a+b+c)<(d+e+f)
		if coins[0 ]+coins[3 ] == coins[1 ]+coins[4 ] { //(a+e)>(d+b)
			compare(coins, 5 , 2 , 0 )
		} else if coins[0 ]+coins[3 ] > coins[1 ]+coins[4 ] { //(a+e)>(d+b)
			compare(coins, 3 , 1 , 0 )
		}
		if coins[0 ]+coins[3 ] < coins[1 ]+coins[4 ] { //(a+e)<(d+b)
			compare(coins, 4 , 0 , 1 )
		}
	}
}

func main() {
	var coins = make ([]int , 8 )
	// Initial coin weight is 10
	for i := 0 ; i < 8 ; i++ {
		coins[i] = 10
	}
	fmt.Printf("输入假币权重(大于或小于10):" )
	var coin int
	fmt.Scanf("%d" , &coin)
	var index = rand.Intn(8 )
	coins[index] = coin
	eightcoins(coins)
	for i := 0 ; i < 8 ; i++ {
		fmt.Printf("%d , " , coins[i])
	}
}
