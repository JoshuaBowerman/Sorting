package main

import (
	"fmt"
//	"../Sorting"
	"../Utils"
	"math"
)

var ArraySize = 20
var Array = Utils.GenArray(ArraySize)

func main() {
	fmt.Println("### Array Sorting ###")
	fmt.Println("Array Length: ", ArraySize)
	if ArraySize <= 32 {
		fmt.Println("Array: ",Array)
	}

}

func PrintArray(a []int){
	var width = math.Floor(math.Log10(float64(ArraySize))) + 2 // this is the number of digits plus one, used for spacing

}