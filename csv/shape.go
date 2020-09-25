package csv

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

func GetShape(filePath string, headless bool) {

	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		os.Stderr.WriteString("Could not load CSV" + filePath + "\n")
		return
	}

	// Create context.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	reader := csv.NewReader(f)
	reader.ReuseRecord = true

	record, err := reader.Read()

	if err != nil {
		fmt.Println("CSV is empty")
		return
	}
	cols := len(record)
	rows := 0
	if headless {
		rows += 1
	}

	// use a waitgroup to manage synchronization
	var wg sync.WaitGroup
	out := make(chan int)
	m := &sync.Mutex{}

	start := time.Now()
	// declare the workers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			shapeWorker(ctx, out, m,  reader)
		}()
	}

	// drain the output and add to rows count
	go func() {
		for res := range out {
			rows += res
		}
	}()

	// wait for worker group to finish and close out
	wg.Wait()
	// when you close(out) it breaks the below loop.
	close(out)


	fmt.Println(rows, ",", cols) // Done, return
	fmt.Println(time.Since(start).Seconds())
}

func shapeWorker(ctx context.Context, dst chan int, m *sync.Mutex, reader *csv.Reader) {

	for {
		m.Lock()
		_, err := reader.Read()
		m.Unlock()
		if err == io.EOF {
			return
		} else if err != nil {
			fmt.Println(err)
			return
		}
		// you might select on ctx.Done().
		dst <- 1
	}
}

func PracticeShape(filePath string) {

	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		os.Stderr.WriteString("Could not load CSV at " + filePath + "\n")
	}

	parser := csv.NewReader(f)

	rows := 0
	cols := 0

	start := time.Now()
	for {
		record, err := parser.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}

		if cols == 0 {
			cols = len(record)
		}
		rows += 1

	}

	fmt.Println(rows, cols)
	fmt.Println(time.Since(start).Seconds())
}