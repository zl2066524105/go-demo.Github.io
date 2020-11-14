package main

import "fmt"

func main() {
	username := ""
	age := 0
	_, _ = fmt.Scanln(&username, &age)
	fmt.Println("帐号信息为：", username, age)
}
