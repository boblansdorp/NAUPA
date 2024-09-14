package naupaProcessor

import (
	"bytes"
	"fmt"
	"naupaGenerator/country"
	"naupaGenerator/models"
	"naupaGenerator/utils"

	"github.com/tealeg/xlsx"
)

// Create a HolderRecord initialized with spaces
func NewHolderRecord() models.HolderRecord {
	// Create a new HolderRecord
	record := models.HolderRecord{
		TRCode:         '1',          // Always set TRCode to "1" according to the PDF spec
		HolderContact1: NewContact(), // Initialize Contact1 with spaces
		HolderContact2: NewContact(), // Initialize Contact2 with spaces
	}

	// Fill each byte array field with spaces
	utils.FillByteArrayWithSpaces(record.HolderTaxID[:], len(record.HolderTaxID))
	utils.FillByteArrayWithSpaces(record.HolderTaxIDExt[:], len(record.HolderTaxIDExt))
	utils.FillByteArrayWithSpaces(record.HolderRptYear[:], len(record.HolderRptYear))
	utils.FillByteArrayWithSpaces(record.HolderRptNumber[:], len(record.HolderRptNumber))
	utils.FillByteArrayWithSpaces(record.HolderSICCode[:], len(record.HolderSICCode))
	utils.FillByteArrayWithSpaces(record.HolderIncState[:], len(record.HolderIncState))
	utils.FillByteArrayWithSpaces(record.HolderIncDateCCYY[:], len(record.HolderIncDateCCYY))
	utils.FillByteArrayWithSpaces(record.HolderIncDateMM[:], len(record.HolderIncDateMM))
	utils.FillByteArrayWithSpaces(record.HolderIncDateDD[:], len(record.HolderIncDateDD))
	utils.FillByteArrayWithSpaces(record.HolderName[:], len(record.HolderName))
	utils.FillByteArrayWithSpaces(record.HolderCity[:], len(record.HolderCity))
	utils.FillByteArrayWithSpaces(record.HolderCounty[:], len(record.HolderCounty))
	utils.FillByteArrayWithSpaces(record.HolderState[:], len(record.HolderState))
	utils.FillByteArrayWithSpaces(record.HolderFaxAC[:], len(record.HolderFaxAC))
	utils.FillByteArrayWithSpaces(record.HolderFaxNbr[:], len(record.HolderFaxNbr))
	utils.FillByteArrayWithSpaces(record.HolderNAICSCode[:], len(record.HolderNAICSCode))
	utils.FillByteArrayWithSpaces(record.Filler[:], len(record.Filler))

	return record
}

// Create a Contact initialized with spaces
func NewContact() models.Contact {
	contact := models.Contact{}

	// Fill each byte array field with spaces
	utils.FillByteArrayWithSpaces(contact.Name[:], len(contact.Name))
	utils.FillByteArrayWithSpaces(contact.Addr1[:], len(contact.Addr1))
	utils.FillByteArrayWithSpaces(contact.Addr2[:], len(contact.Addr2))
	utils.FillByteArrayWithSpaces(contact.Addr3[:], len(contact.Addr3))
	utils.FillByteArrayWithSpaces(contact.City[:], len(contact.City))
	utils.FillByteArrayWithSpaces(contact.State[:], len(contact.State))
	utils.FillByteArrayWithSpaces(contact.ZIP[:], len(contact.ZIP))
	utils.FillByteArrayWithSpaces(contact.Country[:], len(contact.Country))
	utils.FillByteArrayWithSpaces(contact.TelAC[:], len(contact.TelAC))
	utils.FillByteArrayWithSpaces(contact.TelNbr[:], len(contact.TelNbr))
	utils.FillByteArrayWithSpaces(contact.Ext[:], len(contact.Ext))
	utils.FillByteArrayWithSpaces(contact.Email[:], len(contact.Email))

	return contact
}

// Create a PropertyRecord initialized with spaces
func NewPropertyRecord() models.PropertyRecord {
	// Create a new PropertyRecord
	record := models.PropertyRecord{
		TRCode:        '2', // Set TRCode to "2" as per specification
		PropOwnerType: 'P', // Set PropOwnerType to "P" as per specification
		PropNameID:    'C', // Set PropNameID to "C" as per specification for a business
	}

	// Fill each byte array field with spaces
	utils.FillByteArrayWithZeroes(record.PropSequenceNumber[:], len(record.PropSequenceNumber))
	utils.FillByteArrayWithSpaces(record.PropOwnerNameLast[:], len(record.PropOwnerNameLast))
	utils.FillByteArrayWithSpaces(record.PropOwnerNameFirst[:], len(record.PropOwnerNameFirst))
	utils.FillByteArrayWithSpaces(record.PropOwnerNameMiddle[:], len(record.PropOwnerNameMiddle))
	utils.FillByteArrayWithSpaces(record.PropOwnerNamePrefix[:], len(record.PropOwnerNamePrefix))
	utils.FillByteArrayWithSpaces(record.PropOwnerNameSuffix[:], len(record.PropOwnerNameSuffix))
	utils.FillByteArrayWithSpaces(record.PropOwnerNameTitle[:], len(record.PropOwnerNameTitle))
	utils.FillByteArrayWithSpaces(record.PropOwnerAddress1[:], len(record.PropOwnerAddress1))
	utils.FillByteArrayWithSpaces(record.PropOwnerAddress2[:], len(record.PropOwnerAddress2))
	utils.FillByteArrayWithSpaces(record.PropOwnerAddress3[:], len(record.PropOwnerAddress3))
	utils.FillByteArrayWithSpaces(record.PropOwnerCity[:], len(record.PropOwnerCity))
	utils.FillByteArrayWithSpaces(record.PropOwnerCounty[:], len(record.PropOwnerCounty))
	utils.FillByteArrayWithSpaces(record.PropOwnerState[:], len(record.PropOwnerState))
	utils.FillByteArrayWithSpaces(record.PropOwnerZIP[:], len(record.PropOwnerZIP))
	utils.FillByteArrayWithSpaces(record.PropOwnerCountry[:], len(record.PropOwnerCountry))
	utils.FillByteArrayWithSpaces(record.PropOwnerTaxID[:], len(record.PropOwnerTaxID))
	utils.FillByteArrayWithSpaces(record.PropOwnerTaxIDExt[:], len(record.PropOwnerTaxIDExt))
	utils.FillByteArrayWithSpaces(record.PropOwnerDOBCCYY[:], len(record.PropOwnerDOBCCYY))
	utils.FillByteArrayWithZeroes(record.PropOwnerDOBMM[:], len(record.PropOwnerDOBMM))
	utils.FillByteArrayWithZeroes(record.PropOwnerDOBDD[:], len(record.PropOwnerDOBDD))
	utils.FillByteArrayWithZeroes(record.PropStartTransCCYY[:], len(record.PropStartTransCCYY))
	utils.FillByteArrayWithSpaces(record.PropStartTransMM[:], len(record.PropStartTransMM))
	utils.FillByteArrayWithSpaces(record.PropStartTransDD[:], len(record.PropStartTransDD))
	utils.FillByteArrayWithSpaces(record.PropEndTransCCYY[:], len(record.PropEndTransCCYY))
	utils.FillByteArrayWithSpaces(record.PropEndTransMM[:], len(record.PropEndTransMM))
	utils.FillByteArrayWithSpaces(record.PropEndTransDD[:], len(record.PropEndTransDD))
	utils.FillByteArrayWithSpaces(record.PropType[:], len(record.PropType))
	utils.FillByteArrayWithSpaces(record.PropAmountReported[:], len(record.PropAmountReported))
	utils.FillByteArrayWithSpaces(record.PropDeductionType[:], len(record.PropDeductionType))
	utils.FillByteArrayWithSpaces(record.PropDeductionAmount[:], len(record.PropDeductionAmount))
	utils.FillByteArrayWithSpaces(record.PropAmountAdvertised[:], len(record.PropAmountAdvertised))
	utils.FillByteArrayWithSpaces(record.PropAdditionType[:], len(record.PropAdditionType))
	utils.FillByteArrayWithSpaces(record.PropAdditionAmount[:], len(record.PropAdditionAmount))
	utils.FillByteArrayWithSpaces(record.PropDeletionType[:], len(record.PropDeletionType))
	utils.FillByteArrayWithSpaces(record.PropDeletionAmount[:], len(record.PropDeletionAmount))
	utils.FillByteArrayWithSpaces(record.PropAmountRemitted[:], len(record.PropAmountRemitted))
	utils.FillByteArrayWithSpaces(record.PropInterestRate[:], len(record.PropInterestRate))
	utils.FillByteArrayWithSpaces(record.PropStockIssueName[:], len(record.PropStockIssueName))
	utils.FillByteArrayWithSpaces(record.PropStockCUSIP[:], len(record.PropStockCUSIP))
	utils.FillByteArrayWithSpaces(record.PropNumberOfShares[:], len(record.PropNumberOfShares))
	utils.FillByteArrayWithSpaces(record.PropAddShares[:], len(record.PropAddShares))
	utils.FillByteArrayWithSpaces(record.PropDelShares[:], len(record.PropDelShares))
	utils.FillByteArrayWithSpaces(record.PropRemShares[:], len(record.PropRemShares))
	utils.FillByteArrayWithSpaces(record.PropUnexchangedIssueName[:], len(record.PropUnexchangedIssueName))
	utils.FillByteArrayWithSpaces(record.PropUnexchangedCUSIP[:], len(record.PropUnexchangedCUSIP))
	utils.FillByteArrayWithSpaces(record.PropUnexchangedShares[:], len(record.PropUnexchangedShares))
	utils.FillByteArrayWithSpaces(record.PropAccountNumber[:], len(record.PropAccountNumber))
	utils.FillByteArrayWithSpaces(record.PropCheckNumber[:], len(record.PropCheckNumber))
	utils.FillByteArrayWithSpaces(record.PropDescription[:], len(record.PropDescription))
	utils.FillByteArrayWithSpaces(record.PropRelationshipCode[:], len(record.PropRelationshipCode))
	utils.FillByteArrayWithSpaces(record.PropOwnerTypeCode[:], len(record.PropOwnerTypeCode))
	utils.FillByteArrayWithSpaces(record.Filler[:], len(record.Filler))

	return record
}

// Create a SummaryRecord initialized with spaces
func NewSummaryRecord() models.SummaryRecord {
	// Create a new SummaryRecord
	record := models.SummaryRecord{
		TRCode:             '9', // As per the spec
		SummNegativeReport: 'Y', // Default to 'Y', write as space as soon as a propery record is detected
	}

	// Fill each byte array field with spaces
	utils.FillByteArrayWithZeroes(record.SummNbrOfRecords[:], len(record.SummNbrOfRecords))
	utils.FillByteArrayWithZeroes(record.SummNbrOfProperties[:], len(record.SummNbrOfProperties))
	utils.FillByteArrayWithZeroes(record.SummAmountReported[:], len(record.SummAmountReported))
	utils.FillByteArrayWithZeroes(record.SummDeductionAmount[:], len(record.SummDeductionAmount))
	utils.FillByteArrayWithZeroes(record.SummAmountAdvertised[:], len(record.SummAmountAdvertised))
	utils.FillByteArrayWithZeroes(record.SummAdditionAmount[:], len(record.SummAdditionAmount))
	utils.FillByteArrayWithZeroes(record.SummDeletionAmount[:], len(record.SummDeletionAmount))
	utils.FillByteArrayWithZeroes(record.SummAmountRemitted[:], len(record.SummAmountRemitted))
	utils.FillByteArrayWithZeroes(record.SummNbrOfShares[:], len(record.SummNbrOfShares))
	utils.FillByteArrayWithZeroes(record.SummSharesAdd[:], len(record.SummSharesAdd))
	utils.FillByteArrayWithZeroes(record.SummSharesDel[:], len(record.SummSharesDel))
	utils.FillByteArrayWithZeroes(record.SummSharesRemitted[:], len(record.SummSharesRemitted))

	// Pad or truncate the software version, creator, and contact to fit the respective fields
	softwareVersionPadded := utils.PadOrTruncate(models.SoftwareVersion, len(record.SummSoftwareVersion))
	softwareCreatorPadded := utils.PadOrTruncate(models.SoftwareCreator, len(record.SummCreator))
	softwareCreatorContactPadded := utils.PadOrTruncate(models.SoftwareCreatorContact, len(record.SummCreatorContact))

	// Copy the padded or truncated values into the respective fields
	copy(record.SummSoftwareVersion[:], softwareVersionPadded)
	copy(record.SummCreator[:], softwareCreatorPadded)
	copy(record.SummCreatorContact[:], softwareCreatorContactPadded)

	utils.FillByteArrayWithZeroes(record.Filler[:], len(record.Filler))
	return record
}

func WriteRecords(filePath string) (models.HolderRecord, []models.PropertyRecord, models.SummaryRecord, error) {
	holderRecord := NewHolderRecord()
	propertyRecords := []models.PropertyRecord{}
	summaryRecord := NewSummaryRecord()
	xlFile, err := xlsx.OpenFile(filePath)
	if err != nil {
		return holderRecord, propertyRecords, summaryRecord, fmt.Errorf("error opening .xlsx file: %v", err)
	}

	var numRecords = 0
	var numPropertyRecords = 0
	var summPropertyAmount = 0.00

	for _, sheet := range xlFile.Sheets {

		if sheet.Name == "Company" {
			fmt.Printf("Processing sheet: %s\n", sheet.Name)
			//var holderRecord HolderRecord
			numRecords++ // one record for holder
			for rowIndex, row := range sheet.Rows {
				// Skip the first row (header)
				if rowIndex < 1 {
					continue
				}

				// Check if there's data in column 1 (Field Name) and column 2 (Value)
				if len(row.Cells) < 2 {
					continue // Skip if there isn't enough data
				}

				fieldName := row.Cells[0].String() // Get the field name from column 1
				cellValue := row.Cells[1].String() // Get the value from column 2

				// Find the corresponding function in the map and populate the HolderRecord
				if populateFunc, found := fieldMap[fieldName]; found {
					populateFunc(&holderRecord, cellValue)
				} else {
					fmt.Printf("Unknown field name: %s\n", fieldName)
				}
			}
			// After processing, the holderRecord should be populated
			fmt.Printf("Populated Holder Record \n")
			//fmt.Printf("Populated Holder Record: %+v\n", holderRecord)
		}

		// Process the "Data" sheet for PropertyRecords
		if sheet.Name == "Data" {
			// continue //  for testing purposes we can skip the property records
			fmt.Printf("Processing sheet: %s\n", sheet.Name)

			// Define the column mappings for each PropertyRecord field
			columnMap := map[string]int{
				"Last Transaction Dt":         0,   // Column A
				"Property Type":               1,   // Column B
				"Reporting to State":          2,   // Column C
				"Cash Reported":               3,   // Column D
				"Deduction Code":              4,   // Column E
				"Cash Deduction":              5,   // Column F
				"Addition Code":               6,   // Column G
				"Cash Addition":               7,   // Column H
				"Interest Rate":               8,   // Column I
				"Account Number":              9,   // Column J
				"Check Number":                10,  // Column K
				"Prop. Comments":              11,  // Column L
				"Stock Issue Name":            12,  // Column M
				"Stock CUSIP":                 13,  // Column N
				"Stock Ticker Symbol":         14,  // Column O
				"Subissue Name":               15,  // Column P
				"Stk. Delivery Method":        16,  // Column Q
				"Stk. Delivery Acct. #":       17,  // Column R
				"Reported Shares":             18,  // Column S
				"Deleted Shares":              19,  // Column T
				"Original Shares":             20,  // Column U
				"Original Certificate #":      21,  // Column V
				"Original Registration Name":  22,  // Column W
				"Owner1 Corporate Status":     23,  // Column X
				"Owner1 Relationship":         24,  // Column Y
				"Owner1 Type":                 25,  // Column Z
				"Owner1 Tax ID":               26,  // Column AA
				"Owner1 Tax ID Extension":     27,  // Column AB
				"Owner1 Date of Birth":        28,  // Column AC
				"Owner1 Lastname":             29,  // Column AD
				"Owner1 Firstname":            30,  // Column AE
				"Owner1 Middlename":           31,  // Column AF
				"Owner1 Prefix":               32,  // Column AG
				"Owner1 Suffix":               33,  // Column AH
				"Owner1 Title":                34,  // Column AI
				"Owner1 Address line 1":       35,  // Column AJ
				"Owner1 Address line 2":       36,  // Column AK
				"Owner1 Address line 3":       37,  // Column AL
				"Owner1 City":                 38,  // Column AM
				"Owner1 County":               39,  // Column AN
				"Owner1 State":                40,  // Column AO
				"Owner1 Zipcode":              41,  // Column AP
				"Owner1 Country":              42,  // Column AQ
				"Owner1 Email":                43,  // Column AR
				"Owner1 Driver License":       44,  // Column AS
				"Owner1 Driver License State": 45,  // Column AT
				"Description":                 102, // Column AT

				// Add more mappings for Owner2 and Owner3 if necessary
			}

			for rowIndex, row := range sheet.Rows {
				if rowIndex < 4 {
					// Skip the first 4 rows
					continue
				}
				// fmt.Println("Processing row:", row)
				// Ensure the row has enough columns (in this case, 10 as an example)
				if len(row.Cells) < 10 {
					fmt.Printf("Skipping row %d due to insufficient data\n", rowIndex)
					continue
				}
				// Increment the number of records for each property
				numRecords++
				numPropertyRecords++

				summaryRecord.SummNegativeReport = ' ' // Set to space if there are property records
				// Create a new PropertyRecord
				var propertyRecord = NewPropertyRecord()

				// Convert numPropertyRecords to a zero-padded string
				sequenceNumberStr := fmt.Sprintf("%06d", numPropertyRecords) // Pads to 6 digits

				// Copy the padded string into PropSequenceNumber
				copy(propertyRecord.PropSequenceNumber[:], sequenceNumberStr)

				// Populate the PropertyRecord fields using the column mappings

				// Safely check for the single-character fields
				if cellValue := row.Cells[columnMap["Owner1 Type"]].String(); len(cellValue) > 0 {
					propertyRecord.PropOwnerType = cellValue[0] // Use the first character
				}

				// Populate the string fields with the appropriate column values
				copy(propertyRecord.PropType[:], utils.PadOrTruncate(row.Cells[columnMap["Property Type"]].String(), len(propertyRecord.PropType)))
				copy(propertyRecord.PropOwnerTypeCode[:], utils.PadOrTruncate(row.Cells[columnMap["Owner1 Type"]].String(), len(propertyRecord.PropOwnerTypeCode)))
				copy(propertyRecord.PropRelationshipCode[:], utils.PadOrTruncate(row.Cells[columnMap["Owner1 Relationship"]].String(), len(propertyRecord.PropRelationshipCode)))
				// Use this function before copying names
				lastname := utils.PadOrTruncate(row.Cells[columnMap["Owner1 Lastname"]].String(), len(propertyRecord.PropOwnerNameLast))

				firstname := row.Cells[columnMap["Owner1 Firstname"]].String()

				copy(propertyRecord.PropOwnerNameLast[:], utils.PadOrTruncate(lastname, len(propertyRecord.PropOwnerNameLast)))

				copy(propertyRecord.PropOwnerNameFirst[:], utils.PadOrTruncate(firstname, len(propertyRecord.PropOwnerNameFirst)))

				copy(propertyRecord.PropOwnerNameMiddle[:], utils.PadOrTruncate(row.Cells[columnMap["Owner1 Middlename"]].String(), len(propertyRecord.PropOwnerNameMiddle)))
				copy(propertyRecord.PropOwnerAddress1[:], utils.PadOrTruncate(row.Cells[columnMap["Owner1 Address line 1"]].String(), len(propertyRecord.PropOwnerAddress1)))
				copy(propertyRecord.PropOwnerAddress2[:], utils.PadOrTruncate(row.Cells[columnMap["Owner1 Address line 2"]].String(), len(propertyRecord.PropOwnerAddress2)))
				copy(propertyRecord.PropOwnerCity[:], utils.PadOrTruncate(row.Cells[columnMap["Owner1 City"]].String(), len(propertyRecord.PropOwnerCity)))
				copy(propertyRecord.PropOwnerState[:], utils.PadOrTruncate(row.Cells[columnMap["Owner1 State"]].String(), len(propertyRecord.PropOwnerState)))
				copy(propertyRecord.PropOwnerZIP[:], utils.PadOrTruncate(row.Cells[columnMap["Owner1 Zipcode"]].String(), len(propertyRecord.PropOwnerZIP)))

				// Populate other string fields
				copy(propertyRecord.PropOwnerTaxID[:], utils.PadOrTruncate(row.Cells[columnMap["Owner1 Tax ID"]].String(), len(propertyRecord.PropOwnerTaxID)))
				copy(propertyRecord.PropOwnerTaxIDExt[:], utils.PadOrTruncate(row.Cells[columnMap["Owner1 Tax ID Extension"]].String(), len(propertyRecord.PropOwnerTaxIDExt)))

				countryString := row.Cells[columnMap["Owner1 Country"]].String()
				countryFieldLength := 3 // Length of the field
				// Look up the 3-digit country code and pad/truncate it as needed
				countryCode := utils.PadOrTruncate(country.LookupCountryCode(countryString), countryFieldLength)

				copy(propertyRecord.PropOwnerCountry[:], countryCode)

				// Dates and numeric fields
				copy(propertyRecord.PropStartTransCCYY[:], utils.PadOrTruncate(row.Cells[columnMap["Last Transaction Dt"]].String(), len(propertyRecord.PropStartTransCCYY)))

				cashReported, err := row.Cells[columnMap["Cash Reported"]].Float()
				if err != nil {
					// Handle the error here, for example, log it or return from the function
					fmt.Println("row.Cells[columnMap[Cash Reported]]:", row.Cells[columnMap["Cash Reported"]])

					fmt.Println("Error converting cell to float:", err)
					panic(err)
				}
				summPropertyAmount += cashReported

				copy(propertyRecord.PropAmountReported[:], utils.PadOrTruncate(fmt.Sprintf("%010d", int64(cashReported*100)), len(propertyRecord.PropAmountReported))) // use 10 digit padded string in cents not dollars

				copy(propertyRecord.PropAmountAdvertised[:], utils.PadOrTruncate(row.Cells[columnMap["Cash Deduction"]].String(), len(propertyRecord.PropAmountAdvertised)))

				// Stock and CUSIP fields
				copy(propertyRecord.PropStockIssueName[:], utils.PadOrTruncate(row.Cells[columnMap["Stock Issue Name"]].String(), len(propertyRecord.PropStockIssueName)))
				copy(propertyRecord.PropStockCUSIP[:], utils.PadOrTruncate(row.Cells[columnMap["Stock CUSIP"]].String(), len(propertyRecord.PropStockCUSIP)))
				copy(propertyRecord.PropStockIssueName[:], utils.PadOrTruncate(row.Cells[columnMap["Stock Ticker Symbol"]].String(), len(propertyRecord.PropStockIssueName)))

				// Check Number and Account Number fields
				copy(propertyRecord.PropAccountNumber[:], utils.PadOrTruncate(row.Cells[columnMap["Account Number"]].String(), len(propertyRecord.PropAccountNumber)))
				copy(propertyRecord.PropCheckNumber[:], utils.PadOrTruncate(row.Cells[columnMap["Check Number"]].String(), len(propertyRecord.PropCheckNumber)))

				// Description
				desc := row.Cells[columnMap["Description"]].String()
				fmt.Println("desc: ", desc)
				copy(propertyRecord.PropDescription[:], utils.PadOrTruncate(row.Cells[columnMap["Description"]].String(), len(propertyRecord.PropDescription)))

				// Append the populated PropertyRecord to the list of propertyRecords
				propertyRecords = append(propertyRecords, propertyRecord)

			}
			fmt.Printf("Processed %d property records\n", len(propertyRecords))
		}
	}

	// add a summary record

	// Convert the summPropertyAmount to a 12-digit zero-padded string
	// Multiply summPropertyAmount by 100 to account for the two decimal places
	summPropertyAmountCents := int64(summPropertyAmount * 100)

	// Convert the summPropertyAmountCents to a 12-digit zero-padded string
	summPropertyAmountStr := fmt.Sprintf("%012d", summPropertyAmountCents)

	// Copy the padded string into the SummAmountReported field
	copy(summaryRecord.SummAmountReported[:], summPropertyAmountStr)

	numRecords++ // one record for summary record

	// Convert the numRecords to a 6-digit zero-padded string
	numRecordsStr := fmt.Sprintf("%06d", numRecords)

	// Copy the string into the SummNbrOfRecords field
	copy(summaryRecord.SummNbrOfRecords[:], numRecordsStr)

	// Convert the numPropertyRecords to a 6-digit zero-padded string
	numPropertyRecordsStr := fmt.Sprintf("%06d", numPropertyRecords)

	// Copy the string into the SummNbrOfRecords field
	copy(summaryRecord.SummNbrOfProperties[:], numPropertyRecordsStr)

	return holderRecord, propertyRecords, summaryRecord, nil
}

// Format the HolderRecord by looping through its bytes
func FormatHolderRecord(holderRecord models.HolderRecord) []byte {
	// Use a bytes buffer to collect all the bytes
	var buffer bytes.Buffer

	// Append all the fields to the buffer one by one
	buffer.WriteByte(holderRecord.TRCode)

	buffer.Write(holderRecord.HolderTaxID[:])
	buffer.Write(holderRecord.HolderTaxIDExt[:])
	buffer.Write(holderRecord.HolderRptYear[:])
	// Safely write HolderRptType
	if holderRecord.HolderRptType != 0 {
		buffer.WriteByte(holderRecord.HolderRptType)
	} else {
		buffer.WriteByte(' ') // Write space if not set
	}

	// Write HolderRptNumber (no need for extra checks as it's a byte array)
	buffer.Write(holderRecord.HolderRptNumber[:])

	// Safely write HolderRptFormat
	if holderRecord.HolderRptFormat != 0 {
		buffer.WriteByte(holderRecord.HolderRptFormat)
	} else {
		buffer.WriteByte(' ') // Write space if not set
	}

	buffer.Write(holderRecord.HolderSICCode[:])
	buffer.Write(holderRecord.HolderIncState[:])
	buffer.Write(holderRecord.HolderIncDateCCYY[:])
	buffer.Write(holderRecord.HolderIncDateMM[:])
	buffer.Write(holderRecord.HolderIncDateDD[:])
	buffer.Write(holderRecord.HolderName[:])
	buffer.Write(holderRecord.HolderCity[:])
	buffer.Write(holderRecord.HolderCounty[:])
	buffer.Write(holderRecord.HolderState[:])

	// Add Contact1 and Contact2
	formatContact(&buffer, holderRecord.HolderContact1) // Format Contact1
	formatContact(&buffer, holderRecord.HolderContact2) // Format Contact2

	buffer.Write(holderRecord.HolderFaxAC[:])
	buffer.Write(holderRecord.HolderFaxNbr[:])
	buffer.Write(holderRecord.HolderNAICSCode[:])
	buffer.Write(holderRecord.Filler[:])

	// Append CR/LF (Carriage Return + Line Feed) to the buffer
	buffer.Write([]byte{'\r', '\n'})

	// Convert the buffer to a byte slice and ensure it's 625 bytes by appending spaces if necessary
	result := buffer.Bytes()
	return result
}

// Format the PropertyRecord by looping through its bytes
func FormatPropertyRecord(propertyRecord models.PropertyRecord) []byte {
	// Use a bytes buffer to collect all the bytes
	var buffer bytes.Buffer

	// Append all the fields to the buffer one by one
	buffer.WriteByte(propertyRecord.TRCode)
	buffer.Write(propertyRecord.PropSequenceNumber[:])
	buffer.WriteByte(propertyRecord.PropOwnerType)
	buffer.WriteByte(propertyRecord.PropNameID)
	buffer.Write(propertyRecord.PropOwnerNameLast[:])
	buffer.Write(propertyRecord.PropOwnerNameFirst[:])
	//fmt.Println("propertyRecord.PropOwnerNameFirst: ", buffer.Len(), propertyRecord.PropSequenceNumber, propertyRecord.PropOwnerNameFirst)

	buffer.Write(propertyRecord.PropOwnerNameMiddle[:])
	buffer.Write(propertyRecord.PropOwnerNamePrefix[:])
	buffer.Write(propertyRecord.PropOwnerNameSuffix[:])
	buffer.Write(propertyRecord.PropOwnerNameTitle[:])
	buffer.Write(propertyRecord.PropOwnerAddress1[:])
	buffer.Write(propertyRecord.PropOwnerAddress2[:])
	buffer.Write(propertyRecord.PropOwnerAddress3[:])
	buffer.Write(propertyRecord.PropOwnerCity[:])
	buffer.Write(propertyRecord.PropOwnerCounty[:])
	buffer.Write(propertyRecord.PropOwnerState[:])
	buffer.Write(propertyRecord.PropOwnerZIP[:])
	buffer.Write(propertyRecord.PropOwnerCountry[:])
	buffer.Write(propertyRecord.PropOwnerTaxID[:])
	buffer.Write(propertyRecord.PropOwnerTaxIDExt[:])
	buffer.Write(propertyRecord.PropOwnerDOBCCYY[:])
	buffer.Write(propertyRecord.PropOwnerDOBMM[:])
	buffer.Write(propertyRecord.PropOwnerDOBDD[:])
	buffer.Write(propertyRecord.PropStartTransCCYY[:])
	buffer.Write(propertyRecord.PropStartTransMM[:])
	buffer.Write(propertyRecord.PropStartTransDD[:])
	buffer.Write(propertyRecord.PropEndTransCCYY[:])
	buffer.Write(propertyRecord.PropEndTransMM[:])
	buffer.Write(propertyRecord.PropEndTransDD[:])
	buffer.Write(propertyRecord.PropType[:])
	buffer.Write(propertyRecord.PropAmountReported[:])
	buffer.Write(propertyRecord.PropDeductionType[:])
	buffer.Write(propertyRecord.PropDeductionAmount[:])
	buffer.Write(propertyRecord.PropAmountAdvertised[:])
	buffer.Write(propertyRecord.PropAdditionType[:])
	buffer.Write(propertyRecord.PropAdditionAmount[:])
	buffer.Write(propertyRecord.PropDeletionType[:])
	buffer.Write(propertyRecord.PropDeletionAmount[:])
	buffer.Write(propertyRecord.PropAmountRemitted[:])

	// Safely write PropInterestFlag (write space if not set)
	if propertyRecord.PropInterestFlag != 0 {
		buffer.WriteByte(propertyRecord.PropInterestFlag)
	} else {
		buffer.WriteByte(' ') // Write space if not set
	}

	buffer.Write(propertyRecord.PropInterestRate[:])
	buffer.Write(propertyRecord.PropStockIssueName[:])
	buffer.Write(propertyRecord.PropStockCUSIP[:])
	buffer.Write(propertyRecord.PropNumberOfShares[:])
	buffer.Write(propertyRecord.PropAddShares[:])
	buffer.Write(propertyRecord.PropDelShares[:])
	buffer.Write(propertyRecord.PropRemShares[:])
	buffer.Write(propertyRecord.PropUnexchangedIssueName[:])
	buffer.Write(propertyRecord.PropUnexchangedCUSIP[:])
	buffer.Write(propertyRecord.PropUnexchangedShares[:])
	buffer.Write(propertyRecord.PropAccountNumber[:])
	buffer.Write(propertyRecord.PropCheckNumber[:])
	buffer.Write(propertyRecord.PropDescription[:])
	buffer.Write(propertyRecord.PropRelationshipCode[:])
	buffer.Write(propertyRecord.PropOwnerTypeCode[:])
	buffer.Write(propertyRecord.Filler[:])
	// Append CR/LF (Carriage Return + Line Feed) to the buffer
	buffer.Write([]byte{'\r', '\n'})

	// Convert the buffer to a byte slice and ensure it's 625 bytes by appending spaces if necessary
	result := buffer.Bytes()

	// Ensure the result is 625 bytes long, append spaces if needed
	if len(result) < 625 {
		panic("Property record is not 625 bytes long")
		// padding := make([]byte, 625-len(result))
		// for i := range padding {
		// 	padding[i] = ' ' // fill with spaces
		// }
		// result = append(result, padding...)
	}

	return result
}

// Format the SummaryRecord by looping through its bytes
func FormatSummaryRecord(summaryRecord models.SummaryRecord) []byte {
	// Use a bytes buffer to collect all the bytes
	var buffer bytes.Buffer

	// Append all the fields to the buffer one by one
	buffer.WriteByte(summaryRecord.TRCode)              // Single byte for TRCode (position 1)
	buffer.Write(summaryRecord.SummNbrOfRecords[:])     // 6 bytes for number of records (positions 2-7)
	buffer.Write(summaryRecord.SummNbrOfProperties[:])  // 6 bytes for number of properties (positions 8-13)
	buffer.Write(summaryRecord.SummAmountReported[:])   // 12 bytes for amount reported (positions 14-25)
	buffer.Write(summaryRecord.SummDeductionAmount[:])  // 12 bytes for deduction amount (positions 26-37)
	buffer.Write(summaryRecord.SummAmountAdvertised[:]) // 12 bytes for amount advertised (positions 38-49)
	buffer.Write(summaryRecord.SummAdditionAmount[:])   // 12 bytes for addition amount (positions 50-61)
	buffer.Write(summaryRecord.SummDeletionAmount[:])   // 12 bytes for deletion amount (positions 62-73)
	buffer.Write(summaryRecord.SummAmountRemitted[:])   // 12 bytes for amount remitted (positions 74-85)
	buffer.Write(summaryRecord.SummNbrOfShares[:])      // 14 bytes for number of shares (positions 86-99)
	buffer.Write(summaryRecord.SummSharesAdd[:])        // 14 bytes for shares added (positions 100-113)
	buffer.Write(summaryRecord.SummSharesDel[:])        // 14 bytes for shares deleted (positions 114-127)
	buffer.Write(summaryRecord.SummSharesRemitted[:])   // 14 bytes for shares remitted (positions 128-141)
	buffer.WriteByte(summaryRecord.SummNegativeReport)  // Single byte for negative report (position 142)
	buffer.Write(summaryRecord.SummSoftwareVersion[:])  // 20 bytes for software version (positions 143-162)
	buffer.Write(summaryRecord.SummCreator[:])          // 20 bytes for creator (positions 163-182)
	buffer.Write(summaryRecord.SummCreatorContact[:])   // 70 bytes for creator contact (positions 183-252)
	buffer.Write(summaryRecord.Filler[:])               // 373 bytes for filler (positions 253-625)

	// Append CR/LF (Carriage Return + Line Feed) to the buffer
	buffer.Write([]byte{'\r', '\n'})

	// Convert the buffer to a byte slice and ensure it's 625 bytes by appending spaces if necessary
	result := buffer.Bytes()

	return result
}

// Define a map of field names to populate the corresponding HolderRecord fields
var fieldMap = map[string]func(*models.HolderRecord, string){
	"HolderTaxID": func(h *models.HolderRecord, value string) {
		copy(h.HolderTaxID[:], utils.PadOrTruncate(value, len(h.HolderTaxID)))
	},
	"HolderTaxIDExt": func(h *models.HolderRecord, value string) {
		copy(h.HolderTaxIDExt[:], utils.PadOrTruncate(value, len(h.HolderTaxIDExt)))
	},
	"HolderRptYear": func(h *models.HolderRecord, value string) {
		copy(h.HolderRptYear[:], utils.PadOrTruncate(value, len(h.HolderRptYear)))
	},
	"HolderRptType": func(h *models.HolderRecord, value string) {
		if len(value) > 0 {
			h.HolderRptType = value[0]
		}
	},
	"HolderRptNumber": func(h *models.HolderRecord, value string) {
		copy(h.HolderRptNumber[:], utils.PadOrTruncate(value, len(h.HolderRptNumber)))
	},
	"HolderRptFormat": func(h *models.HolderRecord, value string) {
		if len(value) > 0 {
			h.HolderRptFormat = value[0]
		}
	},
	"HolderSICCode": func(h *models.HolderRecord, value string) {
		copy(h.HolderSICCode[:], utils.PadOrTruncate(value, len(h.HolderSICCode)))
	},
	"HolderIncState": func(h *models.HolderRecord, value string) {
		copy(h.HolderIncState[:], utils.PadOrTruncate(value, len(h.HolderIncState)))
	},
	"HolderIncDateCCYY": func(h *models.HolderRecord, value string) {
		copy(h.HolderIncDateCCYY[:], utils.PadOrTruncate(value, len(h.HolderIncDateCCYY)))
	},
	"HolderIncDateMM": func(h *models.HolderRecord, value string) {
		copy(h.HolderIncDateMM[:], utils.PadOrTruncate(value, len(h.HolderIncDateMM)))
	},
	"HolderIncDateDD": func(h *models.HolderRecord, value string) {
		copy(h.HolderIncDateDD[:], utils.PadOrTruncate(value, len(h.HolderIncDateDD)))
	},
	"HolderName": func(h *models.HolderRecord, value string) {
		copy(h.HolderName[:], utils.PadOrTruncate(value, len(h.HolderName)))
	},
	"HolderCity": func(h *models.HolderRecord, value string) {
		copy(h.HolderCity[:], utils.PadOrTruncate(value, len(h.HolderCity)))
	},
	"HolderCounty": func(h *models.HolderRecord, value string) {
		copy(h.HolderCounty[:], utils.PadOrTruncate(value, len(h.HolderCounty)))
	},
	"HolderState": func(h *models.HolderRecord, value string) {
		copy(h.HolderState[:], utils.PadOrTruncate(value, len(h.HolderState)))
	},
	"HolderFaxAC": func(h *models.HolderRecord, value string) {
		copy(h.HolderFaxAC[:], utils.PadOrTruncate(value, len(h.HolderFaxAC)))
	},
	"HolderFaxNbr": func(h *models.HolderRecord, value string) {
		copy(h.HolderFaxNbr[:], utils.PadOrTruncate(value, len(h.HolderFaxNbr)))
	},
	"HolderNAICSCode": func(h *models.HolderRecord, value string) {
		copy(h.HolderNAICSCode[:], utils.PadOrTruncate(value, len(h.HolderNAICSCode)))
	},

	// Updated HolderContact1 fields
	"HOLDER-CONTACT1-NAME": func(h *models.HolderRecord, value string) {
		copy(h.HolderContact1.Name[:], utils.PadOrTruncate(value, len(h.HolderContact1.Name)))
	},
	"HOLDER-CONTACT1-ADDR1": func(h *models.HolderRecord, value string) {
		copy(h.HolderContact1.Addr1[:], utils.PadOrTruncate(value, len(h.HolderContact1.Addr1)))
	},
	"HOLDER-CONTACT1-ADDR2": func(h *models.HolderRecord, value string) {
		copy(h.HolderContact1.Addr2[:], utils.PadOrTruncate(value, len(h.HolderContact1.Addr2)))
	},
	"HOLDER-CONTACT1-ADDR3": func(h *models.HolderRecord, value string) {
		copy(h.HolderContact1.Addr3[:], utils.PadOrTruncate(value, len(h.HolderContact1.Addr3)))
	},
	"HOLDER-CONTACT1-CITY": func(h *models.HolderRecord, value string) {
		copy(h.HolderContact1.City[:], utils.PadOrTruncate(value, len(h.HolderContact1.City)))
	},
	"HOLDER-CONTACT1-STATE": func(h *models.HolderRecord, value string) {
		copy(h.HolderContact1.State[:], utils.PadOrTruncate(value, len(h.HolderContact1.State)))
	},
	"HOLDER-CONTACT1-ZIP": func(h *models.HolderRecord, value string) {
		copy(h.HolderContact1.ZIP[:], utils.PadOrTruncate(value, len(h.HolderContact1.ZIP)))
	},
	"HOLDER-CONTACT1-COUNTRY": func(h *models.HolderRecord, value string) {
		copy(h.HolderContact1.Country[:], utils.PadOrTruncate(value, len(h.HolderContact1.Country)))
	},
	"HOLDER-CONTACT1-TEL-AC": func(h *models.HolderRecord, value string) {
		copy(h.HolderContact1.TelAC[:], utils.PadOrTruncate(value, len(h.HolderContact1.TelAC)))
	},
	"HOLDER-CONTACT1-TEL-NBR": func(h *models.HolderRecord, value string) {
		copy(h.HolderContact1.TelNbr[:], utils.PadOrTruncate(value, len(h.HolderContact1.TelNbr)))
	},
	"HOLDER-CONTACT1-TEL-EXTENSION": func(h *models.HolderRecord, value string) {
		copy(h.HolderContact1.Ext[:], utils.PadOrTruncate(value, len(h.HolderContact1.Ext)))
	},
	"HOLDER-CONTACT1-EMAIL": func(h *models.HolderRecord, value string) {
		copy(h.HolderContact1.Email[:], utils.PadOrTruncate(value, len(h.HolderContact1.Email)))
	},

	// Updated HolderContact2 fields
	"HOLDER-CONTACT2-NAME": func(h *models.HolderRecord, value string) {
		copy(h.HolderContact2.Name[:], utils.PadOrTruncate(value, len(h.HolderContact2.Name)))
	},
	"HOLDER-CONTACT2-ADDR1": func(h *models.HolderRecord, value string) {
		copy(h.HolderContact2.Addr1[:], utils.PadOrTruncate(value, len(h.HolderContact2.Addr1)))
	},
	"HOLDER-CONTACT2-ADDR2": func(h *models.HolderRecord, value string) {
		copy(h.HolderContact2.Addr2[:], utils.PadOrTruncate(value, len(h.HolderContact2.Addr2)))
	},
	"HOLDER-CONTACT2-ADDR3": func(h *models.HolderRecord, value string) {
		copy(h.HolderContact2.Addr3[:], utils.PadOrTruncate(value, len(h.HolderContact2.Addr3)))
	},
	"HOLDER-CONTACT2-CITY": func(h *models.HolderRecord, value string) {
		copy(h.HolderContact2.City[:], utils.PadOrTruncate(value, len(h.HolderContact2.City)))
	},
	"HOLDER-CONTACT2-STATE": func(h *models.HolderRecord, value string) {
		copy(h.HolderContact2.State[:], utils.PadOrTruncate(value, len(h.HolderContact2.State)))
	},
	"HOLDER-CONTACT2-ZIP": func(h *models.HolderRecord, value string) {
		copy(h.HolderContact2.ZIP[:], utils.PadOrTruncate(value, len(h.HolderContact2.ZIP)))
	},
	"HOLDER-CONTACT2-COUNTRY": func(h *models.HolderRecord, value string) {
		copy(h.HolderContact2.Country[:], utils.PadOrTruncate(value, len(h.HolderContact2.Country)))
	},
	"HOLDER-CONTACT2-TEL-AC": func(h *models.HolderRecord, value string) {
		copy(h.HolderContact2.TelAC[:], utils.PadOrTruncate(value, len(h.HolderContact2.TelAC)))
	},
	"HOLDER-CONTACT2-TEL-NBR": func(h *models.HolderRecord, value string) {
		copy(h.HolderContact2.TelNbr[:], utils.PadOrTruncate(value, len(h.HolderContact2.TelNbr)))
	},
	"HOLDER-CONTACT2-TEL-EXTENSION": func(h *models.HolderRecord, value string) {
		copy(h.HolderContact2.Ext[:], utils.PadOrTruncate(value, len(h.HolderContact2.Ext)))
	},
	"HOLDER-CONTACT2-EMAIL": func(h *models.HolderRecord, value string) {
		copy(h.HolderContact2.Email[:], utils.PadOrTruncate(value, len(h.HolderContact2.Email)))
	},
}

// Helper function to format a Contact struct
func formatContact(buffer *bytes.Buffer, contact models.Contact) {
	buffer.Write(contact.Name[:])
	buffer.Write(contact.Addr1[:])
	buffer.Write(contact.Addr2[:])
	buffer.Write(contact.Addr3[:])
	buffer.Write(contact.City[:])
	buffer.Write(contact.State[:])
	buffer.Write(contact.ZIP[:])
	buffer.Write(contact.Country[:])
	buffer.Write(contact.TelAC[:])
	buffer.Write(contact.TelNbr[:])
	buffer.Write(contact.Ext[:])
	buffer.Write(contact.Email[:])
}
