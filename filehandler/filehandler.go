package filehandler

import (
	"fmt"
	"naupaGenerator/models"
	"naupaGenerator/naupaProcessor"
	"os"
)

func CreateNAUPATxtFile(fileName string, holderRecord models.HolderRecord, propertyRecords []models.PropertyRecord, summaryRecord models.SummaryRecord) error {
	// Create a new file or overwrite if exists
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Note: some virus scanners will prevent this program from running")
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	// Write the holder record
	formattedHolderRecord := naupaProcessor.FormatHolderRecord(holderRecord)

	_, err = file.Write(formattedHolderRecord)
	if err != nil {
		return fmt.Errorf("failed to write holder record: %v", err)
	}

	// Write each property record
	for _, propertyRecord := range propertyRecords {
		// fmt.Println("Writing property record to file...", propertyRecord)
		// fmt.Println("propertyRecord: ", propertyRecord)

		formattedPropertyRecord := naupaProcessor.FormatPropertyRecord(propertyRecord)
		if len(formattedPropertyRecord) != 627 {
			return fmt.Errorf("failed to write property record with incorrect length: %d, error: %v", len(formattedPropertyRecord), err)
		}
		// fmt.Println("write formattedPropertyRecord: ", formattedPropertyRecord)
		_, err = file.Write(formattedPropertyRecord)
		if err != nil {
			return fmt.Errorf("failed to write property record: %v", err)
		}
	}

	formattedSummaryRecord := naupaProcessor.FormatSummaryRecord(summaryRecord)

	_, err = file.Write(formattedSummaryRecord)
	if err != nil {
		return fmt.Errorf("failed to write summary record: %v", err)
	}

	return nil
}
