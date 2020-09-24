package csv

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sync"
)

func GetShape(filePath string) {

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
	cols := 0
	rows := 0

	src := make(chan int)
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
			_, err := reader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Println(err)
			}
			// you might select on ctx.Done().
			src <- 1
		}
		// close src to signal workers that no more job are incoming.
		close(src)
	}()

	// wait for worker group to finish and close out
	go func() {
		// wait for writers to quit.
		wg.Wait()
		close(out) // when you close(out) it breaks the below loop.
	}()

	// drain the output
	for res := range out {
		rows += res
	}

	fmt.Println(rows, ",", cols) // Done, return
}

func shapeWorker(ctx context.Context, dst chan int, src chan int) {

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