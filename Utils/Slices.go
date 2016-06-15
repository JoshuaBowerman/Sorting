package Utils

import (
	"math/rand"
)

func Delete(array []int, index int) []int{
	a := array
	a = append(a[:index], a[index+1:]...)
	return a
}

func GenArray(len int) []int{
	var Sorted_Array = make([]int, len)

	for i:= 0; i < len; i++{
		Sorted_Array[i] = i+1
	}

	var Unsorted_Array = make([]int, 0)

	for i:= 0; i < len - 1; i++{
		a := rand.Intn(len - i - 1)
		Unsorted_Array = append(Unsorted_Array,Sorted_Array[a])
		Sorted_Array = Delete(Sorted_Array,a)
	}
	Unsorted_Array = append(Unsorted_Array,Sorted_Array[0])

	return Unsorted_Array
}