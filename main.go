package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func filterHiddenFiles(fd []os.FileInfo) []string {
	files := make([]string, 0)
	for _, file := range fd {
		if !strings.HasPrefix(file.Name(), ".") {
			files = append(files, file.Name())
		}
	}
	return files
}

func main() {
	args := os.Args

	if len(args) < 3 {
		fmt.Println("Error: some arguments are missing")
	}

	subDir := args[2]
	vidDir := args[1]

	vidFiles, _ := ioutil.ReadDir(vidDir)
	subFiles, _ := ioutil.ReadDir(subDir)

	filteredVidFiles := filterHiddenFiles(vidFiles)
	filteredSubFiles := filterHiddenFiles(subFiles)

	for i, subFile := range filteredSubFiles {
		sub := subFile
		vid := filteredVidFiles[i]

		newSubName := strings.TrimSuffix(vid, filepath.Ext(vid)) + filepath.Ext(sub)

		subPath := filepath.Join(subDir, sub)
		renameSubPath := filepath.Join(vidDir, newSubName)

		if err := os.Rename(subPath, renameSubPath); err != nil {
			fmt.Println("Error: couldn't rename file ", sub)
		}
	}

	os.Remove(subDir)
}
