package Sorting

var MethodList = []string{"Merge Sort","Bubble Sort"}
var Method = 0

func SelectMethod(m string) bool{
	for i:= 0; i < len(MethodList); i++{
		if m == MethodList[i] {
			Method = i
			return true
		}
	}
	return false // the method chosen does not exist
}

func Sort(a []int) []int{
	switch Method {
	case 0:
		return Merge_Sort(a)
	case 1:
		return Bubble_Sort(a)
	}
	return make([]int,0)
}