package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func grepPath(path string, fileInfo os.FileInfo, searchStr string) {
	fullPath := filepath.Join(path, fileInfo.Name())
	if !fileInfo.IsDir() {
		content, err := os.ReadFile(fullPath)
		if err != nil {
			log.Fatal(err)
		}
		if strings.Contains(string(content), searchStr) {
			fmt.Println(fullPath, "contains a match with", searchStr)
		} else {
			fmt.Println(fullPath, "does NOT contain a match with", searchStr)
		}
	}

}

func main() {
	searchStr := os.Args[1]
	path := os.Args[2]
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, fileInfo := range files {
		go grepPath(path, fileInfo, searchStr)
	}

	time.Sleep(2 * time.Second)
}
