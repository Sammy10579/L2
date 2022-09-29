package main

import (
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

func handleFile(nameFile string) (answer []string, err error) {
	file, err := os.Open(nameFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteFile, _ := io.ReadAll(file)
	stringFile := string(byteFile)
	return strings.Split(stringFile, "\n"), nil
}

//Обработка строки согласно ключам
func handlerStr(str string, temp string, ignoreCase bool, invert bool, fixed bool) bool {

	// если активен ключ игнорированя регистра - приводим обе строки к одному
	if ignoreCase {
		str = strings.ToLower(str)
		temp = strings.ToLower(temp)
	}

	//если активен ключ фиксированной строки, то возвращаем значения учитывая клч отрицания
	if fixed {
		if str == temp {
			return true && !invert
		}
	}

	//производим поиск подстроки и возвращаем значения учитывая ключ инвертора
	if strings.Contains(str, temp) {
		return true && !invert
	}
	return false || invert
}

func main() {
	// ключи отвечающие за количество обработанных строк
	after := flag.Int("A", 0, "after key")
	before := flag.Int("B", 0, "before key")
	context := flag.Int("C", 0, "context key")

	//Общее количество скрок, которые выводят
	count := flag.Int("c", 0, "count key")

	//ключи отвечающие за обработку строк
	ignoreCase := flag.Bool("i", false, "ignore-case key")
	inverter := flag.Bool("v", false, "iverter key")
	fixed := flag.Bool("fixed", false, "fixed key")
	lineNum := flag.Bool("n", false, "live num key")
	flag.Parse()

	lines := struct {
		left  int
		right int
		count int
	}{left: 0, right: 0, count: 0}

	temp := "man"

	sliceFile, err := handleFile("in.txt")
	if err != nil {
		fmt.Println(err)
	}

	//если общее количество не указанно, то возвращаем строки
	if *count != 0 {
		lines.count = *count
	} else {
		lines.count = len(sliceFile)
	}
	//используем если указан хотя бы один параметр
	switch {
	case *context != 0:
		lines.left, lines.right = *context, *context
	case *after != 0:
		lines.right = *after
	case *before != 0:
		lines.left = *before

	//если не указан, то возвращаться будут все строки
	default:
		lines.right = len(sliceFile)
	}

	//проверяем все строки
	for i, str := range sliceFile {
		//если строка подходит - указываются ключи обработки строки
		if handlerStr(str, temp, *ignoreCase, *inverter, *fixed) {
			//отсекаем строки, не входящие в рез
			maxOfMin := math.Max(0, float64(i-lines.left))
			minOfMax := math.Min(float64(len(sliceFile)-1), float64(i+lines.right))

			//захват дополнительных строк, находящихся около нужной строки
			for j := maxOfMin; j <= minOfMax; j++ {
				if lines.count >= 1 {
					if *lineNum {
						fmt.Println("num", j, "str", sliceFile[int(j)])
						lines.count--
					} else {
						fmt.Println("str", sliceFile[int(j)])
						lines.count--
					}
				} else {
					// lines.count >=1 количество строк для вывода
					os.Exit(1)
				}
			}
		}
	}
}
