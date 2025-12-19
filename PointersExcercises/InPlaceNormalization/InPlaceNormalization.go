package main

import (
	"fmt"
	"math"
)

func f(arr *[]int){
	minOfArr := math.MaxInt

	for i:=0;i<len(*arr);i++ {
		if (*arr)[i] <minOfArr {
			minOfArr = (*arr)[i]
		}
	}
	
	for i:=0;i<len(*arr); i++ {
		(*arr)[i] -= minOfArr
	}
}
func main() {
	arr := []int{11,2,2,4,5,6}
	f(&arr)
	fmt.Println(arr)
}