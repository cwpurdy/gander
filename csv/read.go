package csv

import (
	"encoding/csv"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io"
	"os"
)

func PeekRecords(filePath string) {

	table, err := tablewriter.NewCSV(os.Stdout, filePath, true)
	if err != nil {
		os.Stderr.WriteString("Could not load CSV at " + filePath + "\n")
	}
	table.SetAlignment(tablewriter.ALIGN_LEFT)   // Set Alignment
	table.Render()
}

func GetHeaders(filePath string) {
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		os.Stderr.WriteString("Could not load CSV at " + filePath + "\n")
	}

	parser := csv.NewReader(f)

	headers, err := parser.Read()
	if err == io.EOF {
		return
	}
	if err != nil {
		fmt.Println(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)
	table.Render()
}