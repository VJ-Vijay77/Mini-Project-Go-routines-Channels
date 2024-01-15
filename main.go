package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

type fileInfo struct {
	fileName  string
	wordCount int
	err       error
}

func main() {
	filepaths := []string{"file1.txt", "file2.txt", "file3.txt"}
	resultChannel := make(chan fileInfo, len(filepaths)) //buffered channel

	var wg sync.WaitGroup

	for _, filepath := range filepaths {
		wg.Add(1)
		go processFile(filepath, &wg, resultChannel)
	}

	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	totalWordCount := 0

	for result := range resultChannel {
		if result.err != nil {
			fmt.Printf("\nError processing file %s\nerr:%v\n", result.fileName, result.err)
			continue
		}
		fmt.Printf("\nFileName : %s\nFile Word Count : %d\n", result.fileName, result.wordCount)
		totalWordCount += result.wordCount
	}
	fmt.Printf("\nTotal File Word Count = %d\n", totalWordCount)
}

func processFile(filename string, wg *sync.WaitGroup, result chan<- fileInfo) {
	defer wg.Done()
	var lines []string
	file, err := os.Open(filename)
	if err != nil {
		result <- fileInfo{fileName: filename, err: err}
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// scan each line
		word := scanner.Text()
		// avoid if a line is empty
		if strings.TrimSpace(word) != "" {
			lines = append(lines, word)
		}
	}
	if err := scanner.Err(); err != nil {
		result <- fileInfo{fileName: filename, err: err}
		return
	}
	// joining empty lines
	new := strings.Join(lines, "\n")
	// gettting the whole words
	s := strings.Fields(new)
	result <- fileInfo{fileName: filename, wordCount: len(s)}
	return
}
