package main

import (
	"bufio"
	"flag"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readingFile(file string) []string {
	var result []string

	inFile, _ := os.Open(file)
	defer inFile.Close()

	fileScanner := bufio.NewScanner(inFile)

	for fileScanner.Scan() {
		str := fileScanner.Text()
		result = append(result, str)
	}
	return result
}

func recordFile(file string, array []string) {
	outFile, _ := os.Create(file)
	defer outFile.Close()

	for i := 0; i < len(array)-1; i++ {
		outFile.WriteString(array[i] + "\n")
	}
	outFile.WriteString(array[len(array)-1])
}

func UniqueString(noneUniqueString []string) []string {
	for i, str := range noneUniqueString {
		for j := i + 1; j < len(noneUniqueString); j++ {
			if str == noneUniqueString[j] {
				noneUniqueString = append(noneUniqueString[:i], noneUniqueString[j:]...)
			}
		}
	}
	return noneUniqueString
}

func SortNum(unsorted []string) []string {

	for i := range unsorted {
		var count1 int
		l := strings.Split(unsorted[i], " ")

		for j := range l {
			_, err := strconv.Atoi(l[j])
			if err == nil {
				count1++
			}
		}

		if count1 == len(l) {
			var result []int

			for j := range l {
				k, _ := strconv.Atoi(l[j])
				result = append(result, k)
			}
			sort.Ints(result)

			for m := range result {
				l = append(l[:m], strconv.Itoa(result[m]))
			}
		}
		unsorted[i] = strings.Join(l, " ")
	}
	return unsorted
}

func ReversSort(unsorted []string) []string {
	sort.Sort(sort.Reverse(sort.StringSlice(unsorted)))
	return unsorted
}

func SortByColumn(data []string, colNum int) []string {
	sort.Slice(data, func(i, j int) bool {
		return data[i][colNum] < data[j][colNum]
	})
	return data
}

func main() {

	k := flag.Int("k", 0, "enter column")
	n := flag.Bool("n", false, "sort by num")
	r := flag.Bool("r", false, "revers sort")
	u := flag.Bool("d", false, "do not dublicate lines")

	dataFile := readingFile("test.txt")

	switch {
	case *u:
		dataFile = UniqueString(dataFile)
	case *k != 0:
		dataFile = SortByColumn(dataFile, *k)
	case *r == true:
		dataFile = ReversSort(dataFile)
	case *n == true:
		dataFile = SortNum(dataFile)
	}
	recordFile("outText.txt", dataFile)
}
