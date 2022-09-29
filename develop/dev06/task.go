package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func delimiterBytes(lines [][]byte, delimiterByte []byte, fields *int) {
	res := make([][][]byte, 0)
	for i := range lines {
		temp := bytes.Split(lines[i], delimiterByte)
		if field := *fields; field > -1 {
			if len(temp) > field {
				temp = [][]byte{temp[field]}
			} else {
				temp = [][]byte{}
			}
		}
		res = append(res, temp)
	}
	fmt.Printf("%q", res)
}

func separatedBytes(lines [][]byte, delimiterByte []byte) [][]byte {
	temp := make([][]byte, 0)
	for i := range lines {
		if bytes.Contains(lines[i], delimiterByte) {
			temp = append(temp, lines[i])
		}
	}
	lines = temp
	return lines
}

func main() {
	fields := flag.Int("f", 1, "выбрать поля(колонки)")
	delimit := flag.String("d", "\t", "использовать другой разделитель")
	separated := flag.Bool("s", false, "только строки с разделителем")

	log.SetFlags(0)
	flag.Parse()

	content, err := ioutil.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}

	delimiterByte := []byte(*delimit)
	lines := bytes.Split(content, []byte("\n"))

	if *separated {
		separatedBytes(lines, delimiterByte)
	}

	delimiterBytes(lines, delimiterByte, fields)
}
