package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input")
	defer file.Close()
	reader := bufio.NewReaderSize(file, 1024*1024)
	bytes := []byte{}
	var (
		err error
		m   int
		s   string
	)
	arr := []int{}
	for {
		bytes, _, err = reader.ReadLine()
		if err == io.EOF {
			break
		}
		for _, s = range strings.Fields(string(bytes)) {
			m, _ = strconv.Atoi(s)
			arr = append(arr, m)
		}
		arr = mergeSortIt(arr)
		for _, m = range arr {
			fmt.Printf("%d ", m)
		}
		fmt.Println()
		arr = nil
	}
}

func mergeSortIt(arr []int) []int {
	secarr := make([]int, len(arr))
	mergeSortInternal(arr, secarr, 0, len(arr)-1)
	return arr
}

func mergeSortInternal(arr, secarr []int, left, right int) {
	if left >= right {
		return
	}
	mid := left + ((right - left) / 2)

	mergeSortInternal(arr, secarr, left, mid)
	mergeSortInternal(arr, secarr, mid+1, right)
	mergeTogether(arr, secarr, left, right)
}

func mergeTogether(arr, secarr []int, left, right int) {

	//let's assume the indices are as l1, ... , l2, r1, ... , r2 where l1 and r2 are left and right respectively
	l1 := left
	l2 := left + ((right - left) / 2)
	r1 := l2 + 1
	r2 := right

	//for traversal, assume first half, second half and full array has indices as i, j, k respectively
	i, k := l1, l1
	j := r1

	for i <= l2 && j <= r2 {
		if arr[i] <= arr[j] {
			secarr[k] = arr[i]
			i++
		} else {
			secarr[k] = arr[j]
			j++
		}
		k++
	}
	//finish copying pending from left, if any
	for i <= l2 {
		secarr[k] = arr[i]
		i++
		k++
	}

	//finish copying pending from right, if any
	for j <= r2 {
		secarr[k] = arr[j]
		j++
		k++
	}

	//copy the only operated ones from secondary to primary
	for i := l1; i <= r2; i++ {
		arr[i] = secarr[i]
	}

}
