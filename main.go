package main

import (
	"fmt"
	"io/fs"
	"log"

	"naupaGenerator/filehandler"
	"naupaGenerator/naupaProcessor"
	"path/filepath"
	"strings"
)

func main() {
	dir := "." // Directory to scan for .xlsx files
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Error with finding .xlsx file: %s\n", path)
			fmt.Println("Error: ", err)
			return err
		}

		// If the current entry is a directory and it's not the root directory, skip it
		if d.IsDir() && path != dir {
			return filepath.SkipDir
		}

		// Ignore temporary Excel files that start with "~$"
		if strings.HasPrefix(filepath.Base(path), "~$") {
			fmt.Printf("Ignoring temporary Excel file: %s\n", path)
			return nil // Skip processing for temporary files
		}

		if filepath.Ext(path) == ".xlsx" {
			fmt.Printf("Found .xlsx file: %s\n", path)
			holderRecord, propertyRecords, summaryRecord, err := naupaProcessor.WriteRecords(path)
			if err != nil {
				log.Printf("Error importing .xlsx file: %v\n", err)
				return err
			}
			err = filehandler.CreateNAUPATxtFile("output.txt", holderRecord, propertyRecords, summaryRecord)
			if err != nil {
				log.Printf("Error creating NAUPA .txt file: %v\n", err)
				return err
			}
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Error scanning directory: %v", err)
	}
}
