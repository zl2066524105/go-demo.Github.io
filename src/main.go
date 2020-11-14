package main

import "fmt"

func main() {
	var grade string
	score := 78.5
	switch {
	case score >= 90:
		grade = "A"
	case score >= 80:
		grade = "B"
	case score >= 70:
		grade = "C"
	case score >= 60:
		grade = "D"
	default:
		grade = "E"
	}
	fmt.Printf("%s\n", grade)
	switch grade {
	case "A":
		fmt.Printf("优秀\n")
	case "B":
		fmt.Printf("良好\n")
	case "C":
		fmt.Printf("中等\n")
	case "D":
		fmt.Printf("及格\n")
	default:
		fmt.Printf("差\n")
	}
}
