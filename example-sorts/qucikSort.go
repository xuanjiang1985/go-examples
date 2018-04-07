package main

import (
	"fmt"
)

func main(){
	slice := []int{6,3,5,9,0,1,4,8,2}
	quickSort(slice)
	fmt.Println(slice)
}

func quickSort(values []int){
	if (len(values) <= 1){
		return
	}
	mid, i := values[0], 1
	head, tail := 0, len(values) - 1
	for head < tail {
		fmt.Println(values)
		if values[i] > mid {
			values[i], values[tail] = values[tail], values[i]
			tail--
		} else {
			values[i], values[head] = values[head], values[i]
            head++
            i++
		}
	}
	values[head] = mid
	quickSort(values[:head])
	quickSort(values[head+1:])
}