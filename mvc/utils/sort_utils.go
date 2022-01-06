package utils

import (
	"sort"

)

// BubbleSort
// []int {9,8,7,6,5}
// []int {5,6,7,8,9}
func BubbleSort(elements []int){
	keepRunning := true
	//fmt.Printf("%v\n", elements)
	for keepRunning {
		keepRunning = false
		for i:=0; i < len(elements)-1; i++{
			if elements[i] > elements[i+1] {

			elements[i],elements[i+1] = elements[i+1],elements[i]
			//fmt.Printf("%v\n", elements)
			keepRunning = true
			}
		}
	}
}

func Sort(els []int) {
	if len(els) < 1000 {
		BubbleSort(els)
		return
	}
	sort.Ints(els)
}