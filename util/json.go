package util

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func GetDataToArraysString(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var data [][]string
	if err := json.Unmarshal(bytes, &data); err != nil {
		log.Fatal(err)
	}

	return data
}
