package main

import (
	"fmt"
	"os"
)

const (
	recordSize = 64
	nameSize   = recordSize - 4 // 4 bytes for the ID, the rest for the name
	dbFile     = "warmup.db"
)

func main() {
	// open the file - if there are any errors we print them - then end the program
	f, err := os.OpenFile(dbFile, os.O_RDWR|os.O_CREATE, 0o644)
	if err != nil {
		fmt.Println("open failed:", err)
		os.Exit(1)
	}
	defer f.Close() // defer means - run f.close() when this surrounding function finishes

	record := make([]byte, recordSize)
	copy(record, "Hello Database!")
	_, err = f.WriteAt(record, 0)
	if err != nil {
		fmt.Println("write failed:", err)
		os.Exit(1)
	}

	buffer := make([]byte, recordSize)
	_, err = f.ReadAt(buffer, 0)
	if err != nil {
		fmt.Println("read failed:", err)
		os.Exit(55)
	}
	fmt.Print(string(buffer))
	fmt.Print(buffer)

	fmt.Println("done")
}
