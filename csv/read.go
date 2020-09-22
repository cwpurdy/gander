package csv

import (
	"os"
	"fmt"
)

func GetFile(path string) (*os.File, error) {
	f, err := os.Open("../nasdaq-listed.csv")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("NO ERROR")
	return f, nil
}