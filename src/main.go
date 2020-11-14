package main

import "fmt"

func main() {
	i := 1
	sum := 0
	for i <= 40 {
		if i%3 == 0 {
			sum += i
			fmt.Printf("%d", i)
			if i < 39 {
				fmt.Printf("+")
			} else {
				fmt.Printf("=%d", sum)
			}
		}
		i++
	}
}
