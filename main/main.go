// https://www.hackerrank.com/challenges/find-maximum-index-product/problem

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func left(idx int, arr []int) int {
	for i := idx - 1; i >= 0; i-- {
		if arr[i] > arr[idx] {
			return i + 1
		}
	}

	return 0
}

func right(idx int, arr []int) int {
	for i := idx + 1; i < len(arr); i++ {
		if arr[i] > arr[idx] {
			return i + 1
		}
	}

	return 0
}

func solve(arr []int) int {
	var max int
	var prevValue int

	for i := 0; i < len(arr); i++ {
		if prevValue != 0 && prevValue == arr[i] {
			// The number is the same as previous number, don't bother to check
			continue
		}

		prevValue = arr[i]

		var leftIndex = left(i, arr)
		if leftIndex == 0 {
			continue
		}

		var rightIndex = right(i, arr)
		var indexProduct = leftIndex * rightIndex

		if indexProduct > max {
			max = indexProduct
		}
	}

	return max
}

func main() {
	fmt.Println(solve([]int{5, 4, 3, 4, 5}))
	fmt.Println(solve([]int{5}))

	var inputFile, err = os.Open("input00.txt")

	if err == nil {
		defer func() {
			_ = inputFile.Close()
		}()
	}

	var reader = bufio.NewReader(inputFile)
	var numOfInputStr, _ = reader.ReadString('\n')
	var numOfInput, _ = strconv.Atoi(strings.TrimSpace(numOfInputStr))
	var data = make([]int, numOfInput)

	for i := 0; i < numOfInput; i++ {
		var numberStr, _ = reader.ReadString(' ')
		var number, _ = strconv.Atoi(strings.TrimSpace(numberStr))
		data[i] = number
	}

	fmt.Println(solve(data))
}
