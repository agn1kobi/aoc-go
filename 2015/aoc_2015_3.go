package main

import (
	"fmt"
	"os"
)

func checkErr(err error) bool {
	if err != nil {
		fmt.Println("Error :", err)
		return true
	}
	return false
}

type Point struct {
	x int
	y int
}

func main() {
	data, err := os.ReadFile("input_2015_3")
	if checkErr(err) {
		return
	}
	x, y := 0, 0
	houseMap := make(map[Point]int)
	curHouse := Point{x: x, y: y}
	houseMap[curHouse] = 1

	for i := 0; i < len(data); i++ {
		switch data[i] {
		case '^':
			y++
		case 'v':
			y--
		case '>':
			x++
		case '<':
			x--
		}
		curHouse = Point{x: x, y: y}
		if _, exists := houseMap[curHouse]; exists {
			houseMap[curHouse] += 1
		} else {
			houseMap[curHouse] = 1
		}
	}
	fmt.Println(len(houseMap))

	x, y = 0, 0
	xa, ya := 0, 0

	secMap := make(map[Point]int)
	curHouse = Point{x: x, y: y}
	secMap[curHouse] = 2

	for i := 0; i < len(data); i++ {
		if i%2 == 0 {
			switch data[i] {
			case '^':
				y++
			case 'v':
				y--
			case '>':
				x++
			case '<':
				x--
			}
			curHouse = Point{x: x, y: y}
			if _, exists := secMap[curHouse]; exists {
				secMap[curHouse] += 1
			} else {
				secMap[curHouse] = 1
			}
		} else {
			switch data[i] {
			case '^':
				ya++
			case 'v':
				ya--
			case '>':
				xa++
			case '<':
				xa--
			}
			curHouse = Point{x: xa, y: ya}
			if _, exists := secMap[curHouse]; exists {
				secMap[curHouse] += 1
			} else {
				secMap[curHouse] = 1
			}
		}
	}
	fmt.Println(len(secMap))
}
