package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	var data [][]string
	var examples [][]float64
	var labels []float64

	csvFile, _ := os.Open("data_banknote_authentication.txt")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		data = append(data, line)
	}

	for _, row := range data {
		temp := []float64{}
		for _, feature := range row {
			floatFeature, _ := strconv.ParseFloat(feature, 64)
			temp = append(temp, floatFeature)
		}
		examples = append(examples, temp[:len(temp)-1])
		labels = append(labels, temp[len(temp)-1])
	}

	// fmt.Println(examples[0])
	// fmt.Println(labels[0])

}
