package main

import "fmt"

func main() {
	var C, c int
	C = 1
loop:
	for C < 50 {
		C++
		for c = 2; c < C; c++ {
			if C%c == 0 {
				goto loop
			}
		}
		fmt.Printf("%3d\n", C)
	}
}
