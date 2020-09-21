package main
import "fmt"

const MAXSIZE = 8
const MINSIZE = 1

type Fruit struct {
	name   string
	size   int
	price int
}

func main() {
	var item = [MAXSIZE + 1 ]int {0 }
	var value = [MAXSIZE + 1 ]int {0 }
	var fruits = [5 ]Fruit{
		{"Plum" , 4 , 4500 },
		{"Apple" , 5 , 5700 },
		{"Orange" , 2 , 2250 },
		{"Strawberry" , 1 , 1100 },
		{"Melon" , 6 , 6700 },
	}
	var length = len (fruits)
	for i := 0 ; i < length; i++ {
		for j := fruits[i].size; j <= MAXSIZE; j++ {
			var p = j - fruits[i].size
			var newValue = value[p] + fruits[i].price
			if newValue > value[j] { // 找到最佳解决方案
				value[j] = newValue
				item[j] = i
			}
		}
	}
	fmt.Printf("Item \t Price \n" )
	for i := MAXSIZE; i >= MINSIZE; i = i - fruits[item[i]].size {
		fmt.Printf("%s\t %d \n" , fruits[item[i]].name, fruits[item[i]].price)
	}
	fmt.Printf("Total \t %d" , value[MAXSIZE])
}
