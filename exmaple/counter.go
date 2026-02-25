package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./words.txt")
	log.SetFlags(0) // disables time in log lines
	if err != nil {
		log.Fatalln("failed to open file: ", err)
	}
	wordCount := CountWords(file)
	fmt.Println(wordCount)
}

func CountWords(file *os.File) int {
	const bufferSize = 4096
	buffer := make([]byte, bufferSize)
	totalCount := 0

	for {
		size, err := file.Read(buffer)
		if err != nil {
			break
		}
		totalCount += size
	}
	return totalCount
}
