package main

import (
	"bufio"
	"encoding/csv"
	"os"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"
)

type Data struct {
	sum float64
	desc string
	date string
}

var bankMap map[int]Data
var drebedengiMap map[int]Data

func SpaceMap(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

func main() {
	iteration := 0

	bankMap = make(map[int]Data)
	drebedengiMap = make(map[int]Data)

	file, _ := os.Open("sample.csv")

	r := csv.NewReader(bufio.NewReader(file))

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if len(record) == 4 {
			sum, _ := strconv.ParseFloat(record[2], 64)
			bankMap[iteration] = Data{
				sum, record[1],record[0],
			}
			iteration++
		}
	}

	file2, _ := os.Open("drebedengi.csv")

	r2 := csv.NewReader(bufio.NewReader(file2))

	iteration = 0

	for {
		record, err := r2.Read()

		if err == io.EOF {
			break
		}

		if len(record) >= 3 {
			if record[1] == "Citi Зарплатная" {
				preparedSum := strings.Replace(record[3], ",", ".", -1)
				preparedSum = SpaceMap(preparedSum)
				sum, _ := strconv.ParseFloat(preparedSum, 64)
				drebedengiMap[iteration] = Data{
					sum, record[2],record[5],
				}
				iteration++
			}
		}
	}


	for dbI, dbV := range drebedengiMap {
		for bI, bV := range bankMap {
			if dbV.sum == bV.sum {
				delete(bankMap, bI);
				delete(drebedengiMap, dbI);
				break
			}
		}
	}
	//fmt.Println(bankMap)
	//fmt.Println(drebedengiMap)

	fo, _ := os.Create("output.txt")
	writer := bufio.NewWriter(fo)
	defer fo.Close()
	defer writer.Flush()

	writer.WriteString("************BANK***************\n")

	//fmt.Println(bankMap)
	for _, dbV := range bankMap {
		writer.WriteString(fmt.Sprintf("%v %v %v \n", dbV.sum, dbV.date, dbV.desc))
	}

	writer.WriteString("*********DREBEDENGI************\n")
	for _, dbV := range drebedengiMap {
		writer.WriteString(fmt.Sprintf("%v %v %v \n", dbV.sum, dbV.date, dbV.desc))
	}


}