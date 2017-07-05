package main

import (
	"fmt"
	"strconv"
	"strings"
)

func explodeString(str string) []string {
	var arr []string

	for i := 0; i < len(str) && str[i] != 0x03; {
		i2 := strings.Index(str[i:], ":")
		if i2 == -1 {
			fmt.Println("error", i, str[i:], len(str))
			break
		}
		i2 += i
		sl := str[i:i2]
		l, _ := strconv.Atoi(sl)
		i2++
		if (i2+l) < len(str) {
			arr = append(arr, str[i2:i2+l])
			i = i2 + l
		} else {
			arr = append(arr, str[i2:])
			i = len(str)
		}
	}
	return arr
}
