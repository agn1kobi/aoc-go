package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkErr(err error) bool {
	if err != nil {
		fmt.Println("Error :", err)
		return true
	}
	return false
}

func main() {
	data, err := os.ReadFile("input_2015_2")
	checkErr(err)
	input := string(data)
	boxes := strings.Split(strings.TrimSpace(input), "\n")
	fmt.Printf("%d\n", len(boxes))
	feetPaper := 0
	feetRibbon := 0
	for _, box := range boxes {
		measurements := strings.Split(box, "x")
		l, err := strconv.Atoi(measurements[0])
		if checkErr(err) {
			return
		}
		w, err := strconv.Atoi(measurements[1])
		if checkErr(err) {
			return
		}
		h, err := strconv.Atoi(measurements[2])
		if checkErr(err) {
			return
		}
		fmt.Printf("length: %d, width: %d, height: %d\n", l, w, h)
		feetPaper += (2*l*w + 2*w*h + 2*h*l)
		feetRibbon += (l * w * h)
		smallest, secondSmallest := l, w
		if w < l {
			smallest, secondSmallest = w, l
		}
		if h < smallest {
			smallest, secondSmallest = h, smallest
		} else if h < secondSmallest {
			secondSmallest = h
		}
		feetPaper += (smallest * secondSmallest)
		feetRibbon += (2*smallest + 2*secondSmallest)

	}
	fmt.Println(feetPaper, feetRibbon)
}
