package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func fileGrepWg(filePath string, search string, wg *sync.WaitGroup) {
	defer wg.Done()
	fileGrep(filePath, search)
}

func fileGrep(filePath string, search string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 1
	matchCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, search) {
			matchCount += 1
			fmt.Printf("%q line %4d; %q \n", filePath, lineNum, line)
		}
		lineNum += 1
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	if matchCount == 0 {
		fmt.Printf("No match found in %q \n", filePath)
	}
}

func IsRegularFile(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.Mode().IsRegular(), nil
}

func dirGrep(dirPath string, searchStr string, fileMatch string) {
	//fmt.Printf("dir_path: %q, search: %q, fileMatch: %q \n", dirPath, searchStr, fileMatch)
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	var wg sync.WaitGroup

	for _, entry := range entries {
		if entry.IsDir() {
			newDirPath := filepath.Join(dirPath, entry.Name())
			dirGrep(newDirPath, searchStr, fileMatch)
		} else {
			path := filepath.Join(dirPath, entry.Name())
			isRegular, err := IsRegularFile(path)
			if err != nil {
				fmt.Printf("Skipping %q, error; %q\n", path, err)
			} else {
				if isRegular {
					if fileMatch == "" {
						wg.Add(1)
						fileGrepWg(path, searchStr, &wg)
					} else {
						if strings.Contains(entry.Name(), fileMatch) {
							wg.Add(1)
							fileGrepWg(path, searchStr, &wg)
						} else {
							fmt.Printf("Exclude %q \n", path)
						}
					}
				} else {
					fmt.Printf("Skipping %q, not a regular file\n", path)
				}
			}
		}
	}
	wg.Wait()
}

func usage() {
	name := filepath.Base(os.Args[0])
	fmt.Printf("Usage: %s <dir_path> <search> [match] \n", name)
}

func dirGrepWrapper() {
	var dirPath string
	var searchStr string
	var fileMatch string
	fileMatch = ""
	//fmt.Println("args len: ", len(os.Args))
	switch len(os.Args) {
	case 3:
		dirPath = os.Args[1]
		searchStr = os.Args[2]
		dirGrep(dirPath, searchStr, fileMatch)
	case 4:
		dirPath = os.Args[1]
		searchStr = os.Args[2]
		fileMatch = os.Args[3]
		dirGrep(dirPath, searchStr, fileMatch)
	default:
		usage()
		os.Exit(22)
	}
}

func main() {
	dirGrepWrapper()
}
