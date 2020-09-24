package csv

import (
	"io"
	"os"
	"fmt"
	"encoding/csv"
)

func PeekRecords() {
	f, err := os.Open("../nasdaq-listed.csv")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}

	parser := csv.NewReader(f)

	for {
		record, err := parser.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(record)
	}
}

func GetHeaders(filePath string) {
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}

	parser := csv.NewReader(f)

	record, err := parser.Read()
	if err == io.EOF {
		return
	}
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("HEADERS", record)
}