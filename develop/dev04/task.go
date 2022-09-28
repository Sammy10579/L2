package main

import (
	"fmt"
	"sort"
	"strings"
)

func UniqLower(in []string) []string {
	res := make([]string, 0, len(in))
	u := make(map[string]bool)

	for _, i := range in {
		if !u[i] {
			res = append(res, strings.ToLower(i))
			u[i] = true
		}
	}
	return res
}

func AnagramDict(in []string) map[string][]string {
	if len(in) < 2 {
		return nil
	}

	buffer := make(map[string][]string)

	uniqIn := UniqLower(in)
	for _, i := range uniqIn {
		sorted := []rune(i)
		sort.Slice(sorted, func(i, j int) bool {
			return sorted[i] < sorted[j]
		})

		word := string(sorted)
		buffer[word] = append(buffer[word], i)
	}

	res := make(map[string][]string)
	for _, words := range buffer {
		if len(words) > 1 {
			sort.Strings(words)
			res[words[0]] = words
		}
	}
	return res
}

func main() {
	input := []string{"тест", "листок", "пятка", "пятак", "тяпка", "листок", "пятка", "слиток", "столик"}

	fmt.Println(input)
	fmt.Println(AnagramDict(input))
}
