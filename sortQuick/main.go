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
		arr = quickSortIt(arr)
		for _, m = range arr {
			fmt.Printf("%d ", m)
		}
		fmt.Println()
		arr = nil
	}
}

func quickSortIt(arr []int) []int {
	quickSortInternal(arr, 0, len(arr)-1)
	return arr
}

func quickSortInternal(arr []int, left, right int) {
	if left >= right {
		return
	}
	pivot := arr[(left+right)/2]
	i := partitionIt(arr, left, right, pivot)
	quickSortInternal(arr, left, i-1)
	quickSortInternal(arr, i, right)
}

func partitionIt(arr []int, start, end, pivot int) int {
	tmp := 0
	for start <= end {
		for arr[start] < pivot {
			start++
		}
		for arr[end] > pivot {
			end--
		}
		if start <= end {
			tmp = arr[start]
			arr[start] = arr[end]
			arr[end] = tmp
			start++
			end--
		}
	}
	return start
}
