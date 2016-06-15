package Sorting

import (
	"../Utils"
)

func Merge_Sort(a []int) []int{

	//If we have an array of size one, return that array
	if len(a) == 1{
		return a
	}

	var First_Half, Second_Half = Split(a)
	return Merge(Merge_Sort(First_Half),Merge_Sort(Second_Half))

}

func Split(a []int)(first_half,second_half []int){
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
func Merge (a []int, b []int) []int{
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