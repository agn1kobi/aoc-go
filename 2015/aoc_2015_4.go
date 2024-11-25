package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func hashString(input string) string {
	hash := md5.Sum([]byte(input))
	return fmt.Sprintf("%x", hash)
}

func main() {
	secretKey := "yzbqklnj"
	i := 1
	for {
		input := secretKey + strconv.Itoa(i)
		hash := hashString(input)
		if strings.HasPrefix(hash, "000000") {
			break
		}
		i++
	}
	fmt.Println(i)

}
