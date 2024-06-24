package utils

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
)

func ReadFileToStrings(path string) [][]string {
	originFile, _ := os.Open(path)
	defer originFile.Close()

	byteFile, _ := io.ReadAll(originFile)
	readerCSV := csv.NewReader(strings.NewReader(string(byteFile)))

	csvFile, err := readerCSV.ReadAll()
	if err != nil {
		log.Fatalf("Can't read CSV file: %s", err)
		return csvFile
	}
	return csvFile
}

func ReadFileToByte(path string) []byte {
	file, _ := os.Open(path)
	defer file.Close()

	byteFile, err := io.ReadAll(file)
	if err != nil {
		return nil
	}
	return byteFile
}
