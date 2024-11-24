package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("input_2015_1")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	var x int = 0
	var y int = -1
	for _, value := range data {
		if value == ')' {
			x = x - 1
		} else if value == '(' {
			x = x + 1
		}
	}
	fmt.Println(x)
	x = 0
	for index, value := range data {
		if value == ')' {
			x = x - 1
		} else if value == '(' {
			x = x + 1
		}
		if x < 0 {
			y = index + 1
			break
		}
	}
	fmt.Println(y)
}
