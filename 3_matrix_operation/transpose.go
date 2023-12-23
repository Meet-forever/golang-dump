package main

import (
	"fmt"
	"sync"
	"time"
)

func transposeRow(src [][]int, dest [][]int, idx int, wg *sync.WaitGroup) {
	defer wg.Done()
	N := len(dest)
	for i := 0; i < N; i++ {
		dest[i][idx] = src[idx][i]
	}
}

func transpose(arr [][]int) [][]int {
	M, N := len(arr), len(arr[0])
	var mat [][]int
	for i := 0; i < N; i++ {
		var temp []int
		for j := 0; j < M; j++ {
			temp = append(temp, 0)
		}
		mat = append(mat, temp)
	}
	var wg sync.WaitGroup
	wg.Add(M)
	for i := 0; i < M; i++ {
		go transposeRow(arr, mat, i, &wg)
	}
	wg.Wait()
	return mat
}

func transpose2(arr [][]int) [][]int {
	M, N := len(arr), len(arr[0])
	var mat [][]int
	for i := 0; i < N; i++ {
		var temp []int
		for j := 0; j < M; j++ {
			temp = append(temp, arr[j][i])
		}
		mat = append(mat, temp)
	}
	return mat
}

func main() {
	var arr [][]int
	M, N := 10000, 10000
	for i := 0; i < M; i++ {
		var temp []int
		for j := 0; j < N; j++ {
			temp = append(temp, i*10+(j+1))
		}
		arr = append(arr, temp)
	}
	start := time.Now()
	arr = transpose(arr)
	elapsed := time.Since(start)
	start2 := time.Now()
	arr = transpose2(arr)
	elapsed2 := time.Since(start2)
	fmt.Printf("Total time taken:\nGo func: %d nanoseconds\nNormal func: %d nanoseconds\n", elapsed.Nanoseconds(), elapsed2.Nanoseconds())
	fmt.Printf("Total time taken:\nGo func: %.6f seconds\nNormal func: %.6f seconds\n", elapsed.Seconds(), elapsed2.Seconds())
}
