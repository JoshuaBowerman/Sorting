package Sorting

import (
	"../Utils"
)
func T_Merge_Sort (a []int) []int{
	var ch = make(chan []int)
	go Merge_Sort_Threaded(a,ch)
	return <-ch
}
func Merge_Sort_Threaded(a []int, ch chan []int){

	//If we have an array of size one, return that array
	if len(a) == 1{
		ch <- a
	}
	if(len(a) < 10000){
		var First_Half, Second_Half = TSplit(a)
		var cha = make(chan []int)
		var chb = make(chan []int)
		Merge_Sort_Threaded(First_Half,  cha)
		Merge_Sort_Threaded(Second_Half, chb)
		var First_Sorted = <- cha
		var Second_Sorted = <- chb
		ch <- TMerge(First_Sorted,Second_Sorted)
	} else {
		var First_Half, Second_Half = TSplit(a)
		var cha = make(chan []int)
		var chb = make(chan []int)
		go Merge_Sort_Threaded(First_Half,  cha)
		go Merge_Sort_Threaded(Second_Half, chb)
		var First_Sorted = <-cha
		var Second_Sorted = <-chb
		ch <- TMerge(First_Sorted,Second_Sorted)
	}

}

func TSplit(a []int)(first_half,second_half []int){
	var half   = len(a) / 2
	var first  = make([]int, half) 		// This makes a new array that will contain half the array
	var second = make([]int, len(a) - half) // This one takes the other half and may have an extra item depending on rounding

	for i := 0; i < half; i++{
		first[i] = a[i]
	}
	for i := 0; i < len(a) - half; i++{
		second[i] = a[half + i]
	}

	return first, second
}

/*
	Keep in mind that this function requires 2 sorted lists
 */
func TMerge (a []int, b []int) []int{
	var Result = make([]int,0)
	var t = len(a) + len(b)
	for i := 0; i < t; i++{
		if len(a) == 0{
			Result = append(Result,b[0])
			b = Utils.Delete(b,0)
		} else if len(b) == 0{
			Result = append(Result,a[0])
			a = Utils.Delete(a,0)
		} else if a[0] < b[0]{
			Result = append(Result,a[0])
			a = Utils.Delete(a,0)
		} else {
			Result = append(Result,b[0])
			b = Utils.Delete(b,0)
		}
	}
	return Result
}