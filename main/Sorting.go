package main

import (
	"fmt"
	"../Sorting"
	"../Utils"
	"math"
	"strconv"
	"strings"
)

var ArraySize = 4
var Array = Utils.GenArray(ArraySize)

func main() {
	fmt.Println("### Array Sorting ###")
	fmt.Println("Sorting Method: ", Sorting.MethodList[Sorting.Method])
	fmt.Println("Array Length:   ", ArraySize)
	if ArraySize <= 32 {
		PrintArray("Scrambled Array",Array)
		PrintArray("Sorted Array",Sorting.Sort(Array))
	}

}

func PrintArray(label string,a []int){
	var width = int(math.Floor(math.Log10(float64(ArraySize))) + 2) // this is the number of digits plus one, used for spacing
	var FirstItem = "| %20s |"
	var Item = " %" + strconv.Itoa(width) + "d |"
	var FString = strings.Repeat(Item,ArraySize) + "\n"

	var IArray = make([]interface{}, len(a))
	for i, v := range a {
		IArray[i] = v
	}

	fmt.Printf(FirstItem,label)
	fmt.Printf(FString,IArray...)

}