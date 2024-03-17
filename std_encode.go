package main

import (
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func testBase64() {
	s := "aditya pratama"
	encode := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println(encode)
	decode, err := base64.StdEncoding.DecodeString(encode)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(decode))
	}
}

func testReadCSV() {
	csvString := "name, gender, age\n" +
		"aditya, Men, 29\n" +
		"budi, Men, 30\n" +
		"Celine, Women, 28\n"

	reader := csv.NewReader(strings.NewReader(csvString))
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record)
	}
}

func testWriteCSV() {
	rec := [][]string{
		{"name", "gender", "age"},
		{"aditya", "Men", "29"},
		{"budi", "Men", "30"},
		{"Celine", "Women", "28"},
	}

	write := csv.NewWriter(os.Stdout)
	write.WriteAll(rec) // calls Flush internally

	if err := write.Error(); err != nil {
		fmt.Println("error when write csv")
	}

}

func main() {
	// testBase64()
	// testReadCSV()
	testWriteCSV()
}
