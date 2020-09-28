package csv

import (
	"context"
	"encoding/csv"
	"os"
	"sync"
)

type Shape struct {
	Rows int
	Cols int
}

func GetShape(filePath string, headless bool) (*Shape, error) {

	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return nil, err
	}

	// Create context.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	reader := csv.NewReader(f)
	reader.ReuseRecord = true

	record, err := reader.Read()

	if err != nil {
		return nil, err
	}
	cols := len(record)
	rows := 0
	if headless {
		rows += 1
	}

	// use a wait group to manage synchronization
	var wg sync.WaitGroup
	out := make(chan int, 100)
	m := &sync.Mutex{}

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
	// when you close(out) it breaks the goroutine loop above ^^
	close(out)

	shape := &Shape{rows, cols}
	return shape, nil
}

func shapeWorker(ctx context.Context, dst chan int, m *sync.Mutex, reader *csv.Reader) {

	for {
		m.Lock()
		_, err := reader.Read()
		m.Unlock()
		if err != nil {
			return
		}
		// you might select on ctx.Done().
		dst <- 1
	}
}