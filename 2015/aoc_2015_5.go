package main

import (
	"fmt"
	"os"
	"strings"
)

func checkErr(err error) bool {
	if err != nil {
		fmt.Println("Error :", err)
		return true
	}
	return false
}

func checkVowels(str string) bool {

	vowels := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	vowelCount := 0

	for i := 0; i < len(str); i++ {
		if vowels[rune(str[i])] {

			vowelCount++
		}
	}
	if vowelCount >= 3 {
		return true
	}
	return false
}

func hasDouble(str string) bool {
	for i := 0; i < len(str)-1; i++ {
		if str[i] == str[i+1] {
			return true
		}
	}
	return false
}

func isNotDirty(str string) bool {
	badSeq := []string{"ab", "cd", "pq", "xy"}
	for _, sub := range badSeq {
		if strings.Contains(str, sub) {
			return false
		}
	}
	return true
}

//func pairMagick(str string) bool {
//
//	pairs := map[str]int{}
//
//	return false
//}
//
//func shortPalindrome(str string) bool {
//	return false
//}

func main() {
	data, err := os.ReadFile("input_2015_5")
	if checkErr(err) {
		return
	}

	strs := strings.Split(strings.TrimSpace(string(data)), "\n")

	niceCounter := 0
	for _, str := range strs {
		partialNiceness := 0

		if checkVowels(str) {
			partialNiceness++
		}

		if hasDouble(str) {
			partialNiceness++
		}

		if isNotDirty(str) {
			partialNiceness++
		}

		if partialNiceness == 3 {
			niceCounter++
		}
	}

	niceCounter = 0
	//	for _, str := range strs {
	//		partialNiceness := 0
	//
	//		if pairMagick(str) {
	//			partialNiceness++
	//		}
	//
	//		if shortPalindrome(str) {
	//			partialNiceness++
	//		}
	//
	//		if partialNiceness == 2 {
	//			niceCounter++
	//		}
	//	}

	fmt.Println(niceCounter)
}
