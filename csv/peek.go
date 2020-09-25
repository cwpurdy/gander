package csv

import (
	"encoding/csv"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io"
	"os"
)

func PeekNItems(filePath string, n int, headless bool,) {

	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		os.Stderr.WriteString("Could not load CSV" + filePath + "\n")
		return
	}

	reader := csv.NewReader(f)
	reader.ReuseRecord = true

	head, err := reader.Read()

	if err != nil {
		fmt.Println("CSV is empty")
		return
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(head)
	if headless {
		n -= 1
	}

	for i := 0; i < n; i++ {
		record, err := reader.Read()
		if err == io.EOF {
			table.Render()
			return
		} else if err != nil {
			fmt.Println(err)
			return
		}
		table.Append(record)
	}
	table.Render()

}
