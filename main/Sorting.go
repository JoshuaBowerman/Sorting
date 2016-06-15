package main

import (
	"fmt"
	"../Sorting"
	"../Utils"
	"math"
	"strconv"
	"strings"
)

var ArraySize = 30000

func main() {
	fmt.Println("How Big Should The Array Be?")
	fmt.Scanf("%d \n",&ArraySize)

	fmt.Println("### Array Sorting ###")
	fmt.Println("Sorting Method: ", Sorting.MethodList[Sorting.Method])
	fmt.Println("Array Length:   ", ArraySize)
	fmt.Scan()

	var Array = Utils.GenArray(ArraySize)
	var Sorted_Array = Sorting.Sort(Array)

	if ArraySize <= 32 {
		PrintArray("Scrambled Array",Array)
		PrintArray("Sorted Array",Sorted_Array)
	}
	var isSorted = true
	for i:= 0; i < ArraySize; i++{
		if Sorted_Array[i] != i + 1 {
			isSorted = false
			break
		}
	}
	if isSorted{
		fmt.Println("The List Was Sucessfully Sorted")
	} else {
		fmt.Println("The List Was Not Sorted Sucessfully")
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

