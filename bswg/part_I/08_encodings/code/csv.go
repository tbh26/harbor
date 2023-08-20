package code

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func CsvDemo() {
	fmt.Println("Hello code.CsvDemo() world! (ch8)")

	readCsvDemo()
	writeCsvDemo()

	fmt.Println()
}

func readCsvDemo() {
	fmt.Println("\n=-= read csv demo =-=")

	csvData := `user_id,score,password
"Gopher",1000,"admin"
"BigJ",10,"1234"
"GGBoom",,"1111"
`
	r := csv.NewReader(strings.NewReader(csvData))

	isHeader := true
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(record)
		for _, e := range record {
			if isHeader {
				fmt.Printf("  %-11q ", e) // with quotes
			} else {
				fmt.Printf("   %-9s  ", e)
			}
		}
		fmt.Println()
		isHeader = false
	}

	fmt.Println()
}

func writeCsvDemo() {
	fmt.Println("\n=-= write csv demo =-=")

	data := [][]string{
		{"user_id", "score", "password"},
		{"Gopher", "1000", "admin"},
		{"BigJ", "10", "1234"},
		{"GGBoom", "", "1111"},
	}
	writer := csv.NewWriter(os.Stdout)
	for _, rec := range data {
		err := writer.Write(rec)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()

	fmt.Println()
}
