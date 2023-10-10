package main

import (
	"fmt"
)

func parent(i int) int { return (i - 1) / 2 }
func left(i int) int   { return i*2 + 1 }
func right(i int) int  { return i*2 + 2 }

func pushUp(arr []int, idx int) {
	for arr[parent(idx)] > arr[idx] {
		pid := parent(idx)
		arr[idx], arr[pid] = arr[pid], arr[idx]
		idx = pid
	}
}

func pushDown(arr []int, idx int) {
	N := len(arr)
	l, r := left(idx), right(idx)
	smallest := idx
	if l < N && arr[l] < arr[idx] {
		smallest = l
	}
	if r < N && arr[r] < arr[smallest] {
		smallest = r
	}
	if smallest != idx {
		arr[idx], arr[smallest] = arr[smallest], arr[idx]
		pushDown(arr, smallest)
	}
}

func heapify(arr []int) {
	N := len(arr)
	for i := N/2 - 1; i >= 0; i-- {
		pushDown(arr, i)
	}
}

func push(arr *[]int, val int) {
	*arr = append(*arr, val)
	pushUp(*arr, len(*arr)-1)
}

func pop(arr *[]int) int {
	N := len(*arr)
	if N == 0 {
		panic("Empty Heap")
	}
	val := (*arr)[N-1]
	*arr = (*arr)[:N-1]
	val, (*arr)[0] = (*arr)[0], val
	pushDown(*arr, 0)
	return val
}

func main() {
	arr := []int{5, 2, 1, 4, 3}
	fmt.Println(arr)
	heapify(arr)
	fmt.Println(arr)
	pop(&arr)
	fmt.Println(arr)
	push(&arr, 1)
	fmt.Println(arr)
}
