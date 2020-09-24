package csv

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io"
	"os"
	"sync"
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

func GetShape(filePath string) {
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		os.Stderr.WriteString("Could not load CSV at " + filePath + "\n")
	}

	// Create context.
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	reader := csv.NewReader(f)
	cols := 0
	rows := 0

	src := make(chan []string)
	out := make(chan int)

	// use a waitgroup to manage synchronization
	var wg sync.WaitGroup

	// declare the workers
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			shapeWorker(ctx, out, src)
		}()
	}

	headers, err := reader.Read()
	if err != nil {
		fmt.Println(err)
	}
	cols = len(headers)
	// read the csv and write it to src
	go func() {
		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Println(err)
			}
			src <- record // you might select on ctx.Done().
		}
		close(src) // close src to signal workers that no more job are incoming.
	}()

	// wait for worker group to finish and close out
	go func() {
		wg.Wait() // wait for writers to quit.
		close(out) // when you close(out) it breaks the below loop.
	}()

	// drain the output
	for res := range out {
		rows += res
	}

	fmt.Println(rows, ",", cols)
}

func shapeWorker(ctx context.Context, dst chan int, src chan []string) {

	for {
		select {
		case _, ok := <-src: // you must check for readable state of the channel.
			if !ok {
				return
			}
			dst <- 1
		case <-ctx.Done(): // if the context is cancelled, quit.
			return
		}
	}

}