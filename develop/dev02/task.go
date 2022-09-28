package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Unpack(s string) (str string, err error) {
	if s == "" {
		return "", nil
	}
	if s[0] < 97 || s[0] > 122 {
		return str, fmt.Errorf("incorrent string")
	}

	for i := 0; i <= len(s)-1; i++ {
		if s[i] >= 97 && s[i] <= 122 {
			str += string(s[i])
		} else if s[i] == '\\' {
			str += string(s[i+1])
			i++
		} else {
			bytes := make([]byte, 0)
			for n := i; n <= len(s)-1; n++ {
				if s[n] >= 48 && s[n] <= 57 {
					bytes = append(bytes, s[n])
				} else {
					break
				}
			}
			atoi, err := strconv.Atoi(string(bytes))
			if err != nil {
				panic(err)
			}
			y := strings.Repeat(string(s[i-1]), atoi-1)
			str += y
		}
	}

	return
}

func main() {
	fmt.Println(Unpack("a4bc2d5e"))
}
