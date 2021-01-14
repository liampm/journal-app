package app

import (
	"bufio"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

type frontMatter struct {
	Title       string   `yaml:"title"`
	DateWritten string   `yaml:"dateWritten"`
	Tags        []string `yaml:"tags"`
}

func FindEntryFilesWithTag(searchTag string) []string {
	var matchedFiles []string

	scanEntries(func(fileName string) {
		matter := frontMatterFromFile(fileName)

		for _, tag := range matter.Tags {
			if tag == searchTag {
				matchedFiles = append(matchedFiles, fileName)
				continue
			}
		}
	})
	return matchedFiles
}

func scanEntries(cb func(fileName string)) {
	entryDir := "entry/"

	files, err := ioutil.ReadDir(entryDir)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".jrnl" {
			continue
		}

		fileName := filepath.Join(entryDir, file.Name())

		cb(fileName)
	}
}

func frontMatterFromFile(fileName string) frontMatter {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	inFrontMatter := false
	var frm []byte
	lineNumber := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if lineNumber == 0 && scanner.Text() == "---" {
			inFrontMatter = true
		} else if inFrontMatter && scanner.Text() == "---" {
			inFrontMatter = false
		} else if inFrontMatter {
			frm = append(frm, scanner.Bytes()...)
			frm = append(frm, '\n') // Readd the newline
		}
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	matter := frontMatter{}

	err = yaml.Unmarshal(frm, &matter)
	if err != nil {
		panic(err)
	}
	return matter
}
