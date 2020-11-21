package main

import "fmt"

type processFunc func(int) bool

func main() {
	slice := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd)
	fmt.Println("奇数元素：", odd)
	even := filter(slice, isEven)
	fmt.Println("偶数元素:", even)
}
func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}
func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}
func filter(slice []int, f processFunc) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}
