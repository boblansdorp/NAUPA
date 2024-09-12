package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/tealeg/xlsx"
)

const softwareVersion = "0.0.1"
const softwareCreator = "Bob Lansdorp"
const softwareCreatorContact = "Bob Lansdorp Consulting LLC, https://github.com/boblansdorp/NAUPA"

type HolderRecord struct {
	TRCode            byte     // 1 byte, position 1
	HolderTaxID       [9]byte  // 9 bytes, positions 2-10
	HolderTaxIDExt    [4]byte  // 4 bytes, positions 11-14
	HolderRptYear     [4]byte  // 4 bytes, positions 15-18
	HolderRptType     byte     // 1 byte, position 19
	HolderRptNumber   [2]byte  // 2 bytes, positions 20-21
	HolderRptFormat   byte     // 1 byte, position 22
	HolderSICCode     [4]byte  // 4 bytes, positions 23-26
	HolderIncState    [2]byte  // 2 bytes, positions 27-28
	HolderIncDateCCYY [4]byte  // 4 bytes, positions 29-32
	HolderIncDateMM   [2]byte  // 2 bytes, positions 33-34
	HolderIncDateDD   [2]byte  // 2 bytes, positions 35-36
	HolderName        [40]byte // 40 bytes, positions 37-76
	HolderCity        [30]byte // 30 bytes, positions 77-106
	HolderCounty      [20]byte // 20 bytes, positions 107-126
	HolderState       [2]byte  // 2 bytes, positions 127-128
	HolderContact1    Contact  // Positions 129-316 for Contact1 details
	HolderContact2    Contact  // Positions 367-554 for Contact2 details
	HolderFaxAC       [3]byte  // 3 bytes, positions 605-607
	HolderFaxNbr      [7]byte  // 7 bytes, positions 608-614
	HolderNAICSCode   [6]byte  // 6 bytes, positions 615-620
	Filler            [5]byte  // 5 bytes, positions 621-625
}

type Contact struct {
	Name    [40]byte // 40 bytes, Name
	Addr1   [30]byte // 30 bytes, Address Line 1
	Addr2   [30]byte // 30 bytes, Address Line 2
	Addr3   [30]byte // 30 bytes, Address Line 3
	City    [30]byte // 30 bytes, City
	State   [2]byte  // 2 bytes, State
	ZIP     [9]byte  // 9 bytes, ZIP
	Country [3]byte  // 3 bytes, Country
	TelAC   [3]byte  // 3 bytes, Area Code
	TelNbr  [7]byte  // 7 bytes, Telephone Number
	Ext     [4]byte  // 4 bytes, Extension
	Email   [50]byte // 50 bytes, Email
}

type PropertyRecord struct {
	TRCode                   byte     // 1 byte, position 1
	PropSequenceNumber       [6]byte  // 6 bytes, positions 2-7
	PropOwnerType            byte     // 1 byte, position 8
	PropNameID               byte     // 1 byte, position 9
	PropOwnerNameLast        [40]byte // 40 bytes, positions 10-49
	PropOwnerNameFirst       [30]byte // 30 bytes, positions 50-79
	PropOwnerNameMiddle      [10]byte // 10 bytes, positions 80-89
	PropOwnerNamePrefix      [10]byte // 10 bytes, positions 90-99
	PropOwnerNameSuffix      [10]byte // 10 bytes, positions 100-109
	PropOwnerNameTitle       [6]byte  // 6 bytes, positions 110-115
	PropOwnerAddress1        [30]byte // 30 bytes, positions 116-145
	PropOwnerAddress2        [30]byte // 30 bytes, positions 146-175
	PropOwnerAddress3        [30]byte // 30 bytes, positions 176-205
	PropOwnerCity            [30]byte // 30 bytes, positions 206-235
	PropOwnerCounty          [20]byte // 20 bytes, positions 236-255
	PropOwnerState           [2]byte  // 2 bytes, positions 256-257
	PropOwnerZIP             [9]byte  // 9 bytes, positions 258-266
	PropOwnerCountry         [3]byte  // 3 bytes, positions 267-269
	PropOwnerTaxID           [9]byte  // 9 bytes, positions 270-278
	PropOwnerTaxIDExt        [2]byte  // 2 bytes, positions 279-280
	PropOwnerDOBCCYY         [4]byte  // 4 bytes, positions 281-284
	PropOwnerDOBMM           [2]byte  // 2 bytes, positions 285-286
	PropOwnerDOBDD           [2]byte  // 2 bytes, positions 287-288
	PropStartTransCCYY       [4]byte  // 4 bytes, positions 289-292
	PropStartTransMM         [2]byte  // 2 bytes, positions 293-294
	PropStartTransDD         [2]byte  // 2 bytes, positions 295-296
	PropEndTransCCYY         [4]byte  // 4 bytes, positions 297-300
	PropEndTransMM           [2]byte  // 2 bytes, positions 301-302
	PropEndTransDD           [2]byte  // 2 bytes, positions 303-304
	PropType                 [4]byte  // 4 bytes, positions 305-308
	PropAmountReported       [10]byte // 10 bytes, positions 309-318
	PropDeductionType        [2]byte  // 2 bytes, positions 319-320
	PropDeductionAmount      [10]byte // 10 bytes, positions 321-330
	PropAmountAdvertised     [10]byte // 10 bytes, positions 331-340
	PropAdditionType         [2]byte  // 2 bytes, positions 341-342
	PropAdditionAmount       [10]byte // 10 bytes, positions 343-352
	PropDeletionType         [2]byte  // 2 bytes, positions 353-354
	PropDeletionAmount       [10]byte // 10 bytes, positions 355-364
	PropAmountRemitted       [10]byte // 10 bytes, positions 365-374
	PropInterestFlag         byte     // 1 byte, position 375
	PropInterestRate         [7]byte  // 7 bytes, positions 376-382
	PropStockIssueName       [25]byte // 25 bytes, positions 383-407
	PropStockCUSIP           [9]byte  // 9 bytes, positions 408-416
	PropNumberOfShares       [12]byte // 12 bytes, positions 417-428
	PropAddShares            [12]byte // 12 bytes, positions 429-440
	PropDelShares            [12]byte // 12 bytes, positions 441-452
	PropRemShares            [12]byte // 12 bytes, positions 453-464
	PropUnexchangedIssueName [25]byte // 25 bytes, positions 465-489
	PropUnexchangedCUSIP     [9]byte  // 9 bytes, positions 490-498
	PropUnexchangedShares    [12]byte // 12 bytes, positions 499-510
	PropAccountNumber        [20]byte // 20 bytes, positions 511-530
	PropCheckNumber          [20]byte // 20 bytes, positions 531-550
	PropDescription          [50]byte // 50 bytes, positions 551-600
	PropRelationshipCode     [2]byte  // 2 bytes, positions 601-602
	PropOwnerTypeCode        [2]byte  // 2 bytes, positions 603-604
	Filler                   [21]byte // 21 bytes, positions 605-625
}

type SummaryRecord struct {
	TRCode               byte      // 1 byte, position 1, "9" as per the spec
	SummNbrOfRecords     [6]byte   // 6 bytes, positions 2-7
	SummNbrOfProperties  [6]byte   // 6 bytes, positions 8-13
	SummAmountReported   [12]byte  // 12 bytes, positions 14-25
	SummDeductionAmount  [12]byte  // 12 bytes, positions 26-37
	SummAmountAdvertised [12]byte  // 12 bytes, positions 38-49
	SummAdditionAmount   [12]byte  // 12 bytes, positions 50-61
	SummDeletionAmount   [12]byte  // 12 bytes, positions 62-73
	SummAmountRemitted   [12]byte  // 12 bytes, positions 74-85
	SummNbrOfShares      [14]byte  // 14 bytes, positions 86-99
	SummSharesAdd        [14]byte  // 14 bytes, positions 100-113
	SummSharesDel        [14]byte  // 14 bytes, positions 114-127
	SummSharesRemitted   [14]byte  // 14 bytes, positions 128-141
	SummNegativeReport   byte      // 1 byte, position 142, "Y" or space
	SummSoftwareVersion  [20]byte  // 20 bytes, positions 143-162
	SummCreator          [20]byte  // 20 bytes, positions 163-182
	SummCreatorContact   [70]byte  // 70 bytes, positions 183-252
	Filler               [373]byte // 373 bytes, positions 253-625, space-filled
}

// Create a HolderRecord initialized with spaces
func NewHolderRecord() HolderRecord {
	// Create a new HolderRecord
	record := HolderRecord{
		TRCode:         '1',          // Always set TRCode to "1" according to the PDF spec
		HolderContact1: NewContact(), // Initialize Contact1 with spaces
		HolderContact2: NewContact(), // Initialize Contact2 with spaces
	}

	// Fill each byte array field with spaces
	fillByteArrayWithSpaces(record.HolderTaxID[:], len(record.HolderTaxID))
	fillByteArrayWithSpaces(record.HolderTaxIDExt[:], len(record.HolderTaxIDExt))
	fillByteArrayWithSpaces(record.HolderRptYear[:], len(record.HolderRptYear))
	fillByteArrayWithSpaces(record.HolderRptNumber[:], len(record.HolderRptNumber))
	fillByteArrayWithSpaces(record.HolderSICCode[:], len(record.HolderSICCode))
	fillByteArrayWithSpaces(record.HolderIncState[:], len(record.HolderIncState))
	fillByteArrayWithSpaces(record.HolderIncDateCCYY[:], len(record.HolderIncDateCCYY))
	fillByteArrayWithSpaces(record.HolderIncDateMM[:], len(record.HolderIncDateMM))
	fillByteArrayWithSpaces(record.HolderIncDateDD[:], len(record.HolderIncDateDD))
	fillByteArrayWithSpaces(record.HolderName[:], len(record.HolderName))
	fillByteArrayWithSpaces(record.HolderCity[:], len(record.HolderCity))
	fillByteArrayWithSpaces(record.HolderCounty[:], len(record.HolderCounty))
	fillByteArrayWithSpaces(record.HolderState[:], len(record.HolderState))
	fillByteArrayWithSpaces(record.HolderFaxAC[:], len(record.HolderFaxAC))
	fillByteArrayWithSpaces(record.HolderFaxNbr[:], len(record.HolderFaxNbr))
	fillByteArrayWithSpaces(record.HolderNAICSCode[:], len(record.HolderNAICSCode))
	fillByteArrayWithSpaces(record.Filler[:], len(record.Filler))

	return record
}

// Create a Contact initialized with spaces
func NewContact() Contact {
	contact := Contact{}

	// Fill each byte array field with spaces
	fillByteArrayWithSpaces(contact.Name[:], len(contact.Name))
	fillByteArrayWithSpaces(contact.Addr1[:], len(contact.Addr1))
	fillByteArrayWithSpaces(contact.Addr2[:], len(contact.Addr2))
	fillByteArrayWithSpaces(contact.Addr3[:], len(contact.Addr3))
	fillByteArrayWithSpaces(contact.City[:], len(contact.City))
	fillByteArrayWithSpaces(contact.State[:], len(contact.State))
	fillByteArrayWithSpaces(contact.ZIP[:], len(contact.ZIP))
	fillByteArrayWithSpaces(contact.Country[:], len(contact.Country))
	fillByteArrayWithSpaces(contact.TelAC[:], len(contact.TelAC))
	fillByteArrayWithSpaces(contact.TelNbr[:], len(contact.TelNbr))
	fillByteArrayWithSpaces(contact.Ext[:], len(contact.Ext))
	fillByteArrayWithSpaces(contact.Email[:], len(contact.Email))

	return contact
}

// Create a PropertyRecord initialized with spaces
func NewPropertyRecord() PropertyRecord {
	// Create a new PropertyRecord
	record := PropertyRecord{
		TRCode:        '2', // Set TRCode to "2" as per specification
		PropOwnerType: 'P', // Set PropOwnerType to "P" as per specification
		PropNameID:    'C', // Set PropNameID to "C" as per specification for a business
	}

	// Fill each byte array field with spaces
	fillByteArrayWithZeroes(record.PropSequenceNumber[:], len(record.PropSequenceNumber))
	fillByteArrayWithSpaces(record.PropOwnerNameLast[:], len(record.PropOwnerNameLast))
	fillByteArrayWithSpaces(record.PropOwnerNameFirst[:], len(record.PropOwnerNameFirst))
	fillByteArrayWithSpaces(record.PropOwnerNameMiddle[:], len(record.PropOwnerNameMiddle))
	fillByteArrayWithSpaces(record.PropOwnerNamePrefix[:], len(record.PropOwnerNamePrefix))
	fillByteArrayWithSpaces(record.PropOwnerNameSuffix[:], len(record.PropOwnerNameSuffix))
	fillByteArrayWithSpaces(record.PropOwnerNameTitle[:], len(record.PropOwnerNameTitle))
	fillByteArrayWithSpaces(record.PropOwnerAddress1[:], len(record.PropOwnerAddress1))
	fillByteArrayWithSpaces(record.PropOwnerAddress2[:], len(record.PropOwnerAddress2))
	fillByteArrayWithSpaces(record.PropOwnerAddress3[:], len(record.PropOwnerAddress3))
	fillByteArrayWithSpaces(record.PropOwnerCity[:], len(record.PropOwnerCity))
	fillByteArrayWithSpaces(record.PropOwnerCounty[:], len(record.PropOwnerCounty))
	fillByteArrayWithSpaces(record.PropOwnerState[:], len(record.PropOwnerState))
	fillByteArrayWithSpaces(record.PropOwnerZIP[:], len(record.PropOwnerZIP))
	fillByteArrayWithSpaces(record.PropOwnerCountry[:], len(record.PropOwnerCountry))
	fillByteArrayWithSpaces(record.PropOwnerTaxID[:], len(record.PropOwnerTaxID))
	fillByteArrayWithSpaces(record.PropOwnerTaxIDExt[:], len(record.PropOwnerTaxIDExt))
	fillByteArrayWithSpaces(record.PropOwnerDOBCCYY[:], len(record.PropOwnerDOBCCYY))
	fillByteArrayWithZeroes(record.PropOwnerDOBMM[:], len(record.PropOwnerDOBMM))
	fillByteArrayWithZeroes(record.PropOwnerDOBDD[:], len(record.PropOwnerDOBDD))
	fillByteArrayWithZeroes(record.PropStartTransCCYY[:], len(record.PropStartTransCCYY))
	fillByteArrayWithSpaces(record.PropStartTransMM[:], len(record.PropStartTransMM))
	fillByteArrayWithSpaces(record.PropStartTransDD[:], len(record.PropStartTransDD))
	fillByteArrayWithSpaces(record.PropEndTransCCYY[:], len(record.PropEndTransCCYY))
	fillByteArrayWithSpaces(record.PropEndTransMM[:], len(record.PropEndTransMM))
	fillByteArrayWithSpaces(record.PropEndTransDD[:], len(record.PropEndTransDD))
	fillByteArrayWithSpaces(record.PropType[:], len(record.PropType))
	fillByteArrayWithSpaces(record.PropAmountReported[:], len(record.PropAmountReported))
	fillByteArrayWithSpaces(record.PropDeductionType[:], len(record.PropDeductionType))
	fillByteArrayWithSpaces(record.PropDeductionAmount[:], len(record.PropDeductionAmount))
	fillByteArrayWithSpaces(record.PropAmountAdvertised[:], len(record.PropAmountAdvertised))
	fillByteArrayWithSpaces(record.PropAdditionType[:], len(record.PropAdditionType))
	fillByteArrayWithSpaces(record.PropAdditionAmount[:], len(record.PropAdditionAmount))
	fillByteArrayWithSpaces(record.PropDeletionType[:], len(record.PropDeletionType))
	fillByteArrayWithSpaces(record.PropDeletionAmount[:], len(record.PropDeletionAmount))
	fillByteArrayWithSpaces(record.PropAmountRemitted[:], len(record.PropAmountRemitted))
	fillByteArrayWithSpaces(record.PropInterestRate[:], len(record.PropInterestRate))
	fillByteArrayWithSpaces(record.PropStockIssueName[:], len(record.PropStockIssueName))
	fillByteArrayWithSpaces(record.PropStockCUSIP[:], len(record.PropStockCUSIP))
	fillByteArrayWithSpaces(record.PropNumberOfShares[:], len(record.PropNumberOfShares))
	fillByteArrayWithSpaces(record.PropAddShares[:], len(record.PropAddShares))
	fillByteArrayWithSpaces(record.PropDelShares[:], len(record.PropDelShares))
	fillByteArrayWithSpaces(record.PropRemShares[:], len(record.PropRemShares))
	fillByteArrayWithSpaces(record.PropUnexchangedIssueName[:], len(record.PropUnexchangedIssueName))
	fillByteArrayWithSpaces(record.PropUnexchangedCUSIP[:], len(record.PropUnexchangedCUSIP))
	fillByteArrayWithSpaces(record.PropUnexchangedShares[:], len(record.PropUnexchangedShares))
	fillByteArrayWithSpaces(record.PropAccountNumber[:], len(record.PropAccountNumber))
	fillByteArrayWithSpaces(record.PropCheckNumber[:], len(record.PropCheckNumber))
	fillByteArrayWithSpaces(record.PropDescription[:], len(record.PropDescription))
	fillByteArrayWithSpaces(record.PropRelationshipCode[:], len(record.PropRelationshipCode))
	fillByteArrayWithSpaces(record.PropOwnerTypeCode[:], len(record.PropOwnerTypeCode))
	fillByteArrayWithSpaces(record.Filler[:], len(record.Filler))

	return record
}

// Create a SummaryRecord initialized with spaces
func NewSummaryRecord() SummaryRecord {
	// Create a new SummaryRecord
	record := SummaryRecord{
		TRCode:             '9', // As per the spec
		SummNegativeReport: 'Y', // Default to 'Y', write as space as soon as a propery record is detected
	}

	// Fill each byte array field with spaces
	fillByteArrayWithZeroes(record.SummNbrOfRecords[:], len(record.SummNbrOfRecords))
	fillByteArrayWithZeroes(record.SummNbrOfProperties[:], len(record.SummNbrOfProperties))
	fillByteArrayWithZeroes(record.SummAmountReported[:], len(record.SummAmountReported))
	fillByteArrayWithZeroes(record.SummDeductionAmount[:], len(record.SummDeductionAmount))
	fillByteArrayWithZeroes(record.SummAmountAdvertised[:], len(record.SummAmountAdvertised))
	fillByteArrayWithZeroes(record.SummAdditionAmount[:], len(record.SummAdditionAmount))
	fillByteArrayWithZeroes(record.SummDeletionAmount[:], len(record.SummDeletionAmount))
	fillByteArrayWithZeroes(record.SummAmountRemitted[:], len(record.SummAmountRemitted))
	fillByteArrayWithZeroes(record.SummNbrOfShares[:], len(record.SummNbrOfShares))
	fillByteArrayWithZeroes(record.SummSharesAdd[:], len(record.SummSharesAdd))
	fillByteArrayWithZeroes(record.SummSharesDel[:], len(record.SummSharesDel))
	fillByteArrayWithZeroes(record.SummSharesRemitted[:], len(record.SummSharesRemitted))

	// Pad or truncate the software version, creator, and contact to fit the respective fields
	softwareVersionPadded := padOrTruncate(softwareVersion, len(record.SummSoftwareVersion))
	softwareCreatorPadded := padOrTruncate(softwareCreator, len(record.SummCreator))
	softwareCreatorContactPadded := padOrTruncate(softwareCreatorContact, len(record.SummCreatorContact))

	// Copy the padded or truncated values into the respective fields
	copy(record.SummSoftwareVersion[:], softwareVersionPadded)
	copy(record.SummCreator[:], softwareCreatorPadded)
	copy(record.SummCreatorContact[:], softwareCreatorContactPadded)

	fillByteArrayWithZeroes(record.Filler[:], len(record.Filler))
	return record
}

// Function that fills a byte array with spaces
func fillByteArrayWithSpaces(arr []byte, length int) {
	for i := 0; i < length; i++ {
		arr[i] = ' ' // Fill each position in the array with a space character
	}
}

// Function that fills a byte array with spaces
func fillByteArrayWithZeroes(arr []byte, length int) {
	for i := 0; i < length; i++ {
		arr[i] = '0' // Fill each position in the array with a space character
	}
}

func main() {
	dir := "." // Directory to scan for .xlsx files

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error with finding .xlsx file: %s\n", path)
			fmt.Println("Error: ", err)
			return err
		}

		// Ignore temporary Excel files that start with "~$"
		if strings.HasPrefix(filepath.Base(path), "~$") {
			fmt.Printf("Ignoring temporary Excel file: %s\n", path)
			return nil // Skip processing for temporary files
		}

		if filepath.Ext(path) == ".xlsx" {
			fmt.Printf("Found .xlsx file: %s\n", path)
			holderRecord, propertyRecords, summaryRecord, err := writeRecords(path)
			if err != nil {
				log.Printf("Error importing .xlsx file: %v\n", err)
				return err
			}
			err = createNAUPATxtFile("output.txt", holderRecord, propertyRecords, summaryRecord)
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

// Define a map of field names to populate the corresponding HolderRecord fields
var fieldMap = map[string]func(*HolderRecord, string){
	"HolderTaxID": func(h *HolderRecord, value string) {
		copy(h.HolderTaxID[:], padOrTruncate(value, len(h.HolderTaxID)))
	},
	"HolderTaxIDExt": func(h *HolderRecord, value string) {
		copy(h.HolderTaxIDExt[:], padOrTruncate(value, len(h.HolderTaxIDExt)))
	},
	"HolderRptYear": func(h *HolderRecord, value string) {
		copy(h.HolderRptYear[:], padOrTruncate(value, len(h.HolderRptYear)))
	},
	"HolderRptType": func(h *HolderRecord, value string) {
		if len(value) > 0 {
			h.HolderRptType = value[0]
		}
	},
	"HolderRptNumber": func(h *HolderRecord, value string) {
		copy(h.HolderRptNumber[:], padOrTruncate(value, len(h.HolderRptNumber)))
	},
	"HolderRptFormat": func(h *HolderRecord, value string) {
		if len(value) > 0 {
			h.HolderRptFormat = value[0]
		}
	},
	"HolderSICCode": func(h *HolderRecord, value string) {
		copy(h.HolderSICCode[:], padOrTruncate(value, len(h.HolderSICCode)))
	},
	"HolderIncState": func(h *HolderRecord, value string) {
		copy(h.HolderIncState[:], padOrTruncate(value, len(h.HolderIncState)))
	},
	"HolderIncDateCCYY": func(h *HolderRecord, value string) {
		copy(h.HolderIncDateCCYY[:], padOrTruncate(value, len(h.HolderIncDateCCYY)))
	},
	"HolderIncDateMM": func(h *HolderRecord, value string) {
		copy(h.HolderIncDateMM[:], padOrTruncate(value, len(h.HolderIncDateMM)))
	},
	"HolderIncDateDD": func(h *HolderRecord, value string) {
		copy(h.HolderIncDateDD[:], padOrTruncate(value, len(h.HolderIncDateDD)))
	},
	"HolderName": func(h *HolderRecord, value string) {
		copy(h.HolderName[:], padOrTruncate(value, len(h.HolderName)))
	},
	"HolderCity": func(h *HolderRecord, value string) {
		copy(h.HolderCity[:], padOrTruncate(value, len(h.HolderCity)))
	},
	"HolderCounty": func(h *HolderRecord, value string) {
		copy(h.HolderCounty[:], padOrTruncate(value, len(h.HolderCounty)))
	},
	"HolderState": func(h *HolderRecord, value string) {
		copy(h.HolderState[:], padOrTruncate(value, len(h.HolderState)))
	},
	"HolderFaxAC": func(h *HolderRecord, value string) {
		copy(h.HolderFaxAC[:], padOrTruncate(value, len(h.HolderFaxAC)))
	},
	"HolderFaxNbr": func(h *HolderRecord, value string) {
		copy(h.HolderFaxNbr[:], padOrTruncate(value, len(h.HolderFaxNbr)))
	},
	"HolderNAICSCode": func(h *HolderRecord, value string) {
		copy(h.HolderNAICSCode[:], padOrTruncate(value, len(h.HolderNAICSCode)))
	},

	// Updated HolderContact1 fields
	"HOLDER-CONTACT1-NAME": func(h *HolderRecord, value string) {
		copy(h.HolderContact1.Name[:], padOrTruncate(value, len(h.HolderContact1.Name)))
	},
	"HOLDER-CONTACT1-ADDR1": func(h *HolderRecord, value string) {
		copy(h.HolderContact1.Addr1[:], padOrTruncate(value, len(h.HolderContact1.Addr1)))
	},
	"HOLDER-CONTACT1-ADDR2": func(h *HolderRecord, value string) {
		copy(h.HolderContact1.Addr2[:], padOrTruncate(value, len(h.HolderContact1.Addr2)))
	},
	"HOLDER-CONTACT1-ADDR3": func(h *HolderRecord, value string) {
		copy(h.HolderContact1.Addr3[:], padOrTruncate(value, len(h.HolderContact1.Addr3)))
	},
	"HOLDER-CONTACT1-CITY": func(h *HolderRecord, value string) {
		copy(h.HolderContact1.City[:], padOrTruncate(value, len(h.HolderContact1.City)))
	},
	"HOLDER-CONTACT1-STATE": func(h *HolderRecord, value string) {
		copy(h.HolderContact1.State[:], padOrTruncate(value, len(h.HolderContact1.State)))
	},
	"HOLDER-CONTACT1-ZIP": func(h *HolderRecord, value string) {
		copy(h.HolderContact1.ZIP[:], padOrTruncate(value, len(h.HolderContact1.ZIP)))
	},
	"HOLDER-CONTACT1-COUNTRY": func(h *HolderRecord, value string) {
		copy(h.HolderContact1.Country[:], padOrTruncate(value, len(h.HolderContact1.Country)))
	},
	"HOLDER-CONTACT1-TEL-AC": func(h *HolderRecord, value string) {
		copy(h.HolderContact1.TelAC[:], padOrTruncate(value, len(h.HolderContact1.TelAC)))
	},
	"HOLDER-CONTACT1-TEL-NBR": func(h *HolderRecord, value string) {
		copy(h.HolderContact1.TelNbr[:], padOrTruncate(value, len(h.HolderContact1.TelNbr)))
	},
	"HOLDER-CONTACT1-TEL-EXTENSION": func(h *HolderRecord, value string) {
		copy(h.HolderContact1.Ext[:], padOrTruncate(value, len(h.HolderContact1.Ext)))
	},
	"HOLDER-CONTACT1-EMAIL": func(h *HolderRecord, value string) {
		copy(h.HolderContact1.Email[:], padOrTruncate(value, len(h.HolderContact1.Email)))
	},

	// Updated HolderContact2 fields
	"HOLDER-CONTACT2-NAME": func(h *HolderRecord, value string) {
		copy(h.HolderContact2.Name[:], padOrTruncate(value, len(h.HolderContact2.Name)))
	},
	"HOLDER-CONTACT2-ADDR1": func(h *HolderRecord, value string) {
		copy(h.HolderContact2.Addr1[:], padOrTruncate(value, len(h.HolderContact2.Addr1)))
	},
	"HOLDER-CONTACT2-ADDR2": func(h *HolderRecord, value string) {
		copy(h.HolderContact2.Addr2[:], padOrTruncate(value, len(h.HolderContact2.Addr2)))
	},
	"HOLDER-CONTACT2-ADDR3": func(h *HolderRecord, value string) {
		copy(h.HolderContact2.Addr3[:], padOrTruncate(value, len(h.HolderContact2.Addr3)))
	},
	"HOLDER-CONTACT2-CITY": func(h *HolderRecord, value string) {
		copy(h.HolderContact2.City[:], padOrTruncate(value, len(h.HolderContact2.City)))
	},
	"HOLDER-CONTACT2-STATE": func(h *HolderRecord, value string) {
		copy(h.HolderContact2.State[:], padOrTruncate(value, len(h.HolderContact2.State)))
	},
	"HOLDER-CONTACT2-ZIP": func(h *HolderRecord, value string) {
		copy(h.HolderContact2.ZIP[:], padOrTruncate(value, len(h.HolderContact2.ZIP)))
	},
	"HOLDER-CONTACT2-COUNTRY": func(h *HolderRecord, value string) {
		copy(h.HolderContact2.Country[:], padOrTruncate(value, len(h.HolderContact2.Country)))
	},
	"HOLDER-CONTACT2-TEL-AC": func(h *HolderRecord, value string) {
		copy(h.HolderContact2.TelAC[:], padOrTruncate(value, len(h.HolderContact2.TelAC)))
	},
	"HOLDER-CONTACT2-TEL-NBR": func(h *HolderRecord, value string) {
		copy(h.HolderContact2.TelNbr[:], padOrTruncate(value, len(h.HolderContact2.TelNbr)))
	},
	"HOLDER-CONTACT2-TEL-EXTENSION": func(h *HolderRecord, value string) {
		copy(h.HolderContact2.Ext[:], padOrTruncate(value, len(h.HolderContact2.Ext)))
	},
	"HOLDER-CONTACT2-EMAIL": func(h *HolderRecord, value string) {
		copy(h.HolderContact2.Email[:], padOrTruncate(value, len(h.HolderContact2.Email)))
	},
}

// Helper function to pad or truncate a string to the exact length
func padOrTruncate(value string, length int) string {
	// If value is longer than the required length, truncate it
	if len(value) > length {
		return value[:length]
	}

	// If value is shorter, pad with spaces
	return fmt.Sprintf("%-*s", length, value)
}

func writeRecords(filePath string) (HolderRecord, []PropertyRecord, SummaryRecord, error) {
	holderRecord := NewHolderRecord()
	propertyRecords := []PropertyRecord{}
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
				"Last Transaction Dt":         0,  // Column A
				"Property Type":               1,  // Column B
				"Reporting to State":          2,  // Column C
				"Cash Reported":               3,  // Column D
				"Deduction Code":              4,  // Column E
				"Cash Deduction":              5,  // Column F
				"Addition Code":               6,  // Column G
				"Cash Addition":               7,  // Column H
				"Interest Rate":               8,  // Column I
				"Account Number":              9,  // Column J
				"Check Number":                10, // Column K
				"Prop. Comments":              11, // Column L
				"Stock Issue Name":            12, // Column M
				"Stock CUSIP":                 13, // Column N
				"Stock Ticker Symbol":         14, // Column O
				"Subissue Name":               15, // Column P
				"Stk. Delivery Method":        16, // Column Q
				"Stk. Delivery Acct. #":       17, // Column R
				"Reported Shares":             18, // Column S
				"Deleted Shares":              19, // Column T
				"Original Shares":             20, // Column U
				"Original Certificate #":      21, // Column V
				"Original Registration Name":  22, // Column W
				"Owner1 Corporate Status":     23, // Column X
				"Owner1 Relationship":         24, // Column Y
				"Owner1 Type":                 25, // Column Z
				"Owner1 Tax ID":               26, // Column AA
				"Owner1 Tax ID Extension":     27, // Column AB
				"Owner1 Date of Birth":        28, // Column AC
				"Owner1 Lastname":             29, // Column AD
				"Owner1 Firstname":            30, // Column AE
				"Owner1 Middlename":           31, // Column AF
				"Owner1 Prefix":               32, // Column AG
				"Owner1 Suffix":               33, // Column AH
				"Owner1 Title":                34, // Column AI
				"Owner1 Address line 1":       35, // Column AJ
				"Owner1 Address line 2":       36, // Column AK
				"Owner1 Address line 3":       37, // Column AL
				"Owner1 City":                 38, // Column AM
				"Owner1 County":               39, // Column AN
				"Owner1 State":                40, // Column AO
				"Owner1 Zipcode":              41, // Column AP
				"Owner1 Country":              42, // Column AQ
				"Owner1 Email":                43, // Column AR
				"Owner1 Driver License":       44, // Column AS
				"Owner1 Driver License State": 45, // Column AT
				// Add more mappings for Owner2 and Owner3 if necessary
			}

			for rowIndex, row := range sheet.Rows {
				if rowIndex < 4 {
					// Skip the first 4 rows
					continue
				}

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
				copy(propertyRecord.PropType[:], padOrTruncate(row.Cells[columnMap["Property Type"]].String(), len(propertyRecord.PropType)))
				copy(propertyRecord.PropOwnerTypeCode[:], padOrTruncate(row.Cells[columnMap["Owner1 Type"]].String(), len(propertyRecord.PropOwnerTypeCode)))
				copy(propertyRecord.PropRelationshipCode[:], padOrTruncate(row.Cells[columnMap["Owner1 Relationship"]].String(), len(propertyRecord.PropRelationshipCode)))
				copy(propertyRecord.PropOwnerNameLast[:], padOrTruncate(row.Cells[columnMap["Owner1 Lastname"]].String(), len(propertyRecord.PropOwnerNameLast)))
				copy(propertyRecord.PropOwnerNameFirst[:], padOrTruncate(row.Cells[columnMap["Owner1 Firstname"]].String(), len(propertyRecord.PropOwnerNameFirst)))
				copy(propertyRecord.PropOwnerNameMiddle[:], padOrTruncate(row.Cells[columnMap["Owner1 Middlename"]].String(), len(propertyRecord.PropOwnerNameMiddle)))
				copy(propertyRecord.PropOwnerAddress1[:], padOrTruncate(row.Cells[columnMap["Owner1 Address line 1"]].String(), len(propertyRecord.PropOwnerAddress1)))
				copy(propertyRecord.PropOwnerAddress2[:], padOrTruncate(row.Cells[columnMap["Owner1 Address line 2"]].String(), len(propertyRecord.PropOwnerAddress2)))
				copy(propertyRecord.PropOwnerCity[:], padOrTruncate(row.Cells[columnMap["Owner1 City"]].String(), len(propertyRecord.PropOwnerCity)))
				copy(propertyRecord.PropOwnerState[:], padOrTruncate(row.Cells[columnMap["Owner1 State"]].String(), len(propertyRecord.PropOwnerState)))
				copy(propertyRecord.PropOwnerZIP[:], padOrTruncate(row.Cells[columnMap["Owner1 Zipcode"]].String(), len(propertyRecord.PropOwnerZIP)))

				// Populate other string fields
				copy(propertyRecord.PropOwnerTaxID[:], padOrTruncate(row.Cells[columnMap["Owner1 Tax ID"]].String(), len(propertyRecord.PropOwnerTaxID)))
				copy(propertyRecord.PropOwnerTaxIDExt[:], padOrTruncate(row.Cells[columnMap["Owner1 Tax ID Extension"]].String(), len(propertyRecord.PropOwnerTaxIDExt)))
				copy(propertyRecord.PropOwnerCountry[:], padOrTruncate(row.Cells[columnMap["Owner1 Country"]].String(), len(propertyRecord.PropOwnerCountry)))

				// Dates and numeric fields
				copy(propertyRecord.PropStartTransCCYY[:], padOrTruncate(row.Cells[columnMap["Last Transaction Dt"]].String(), len(propertyRecord.PropStartTransCCYY)))

				cashReported, err := row.Cells[columnMap["Cash Reported"]].Float()
				if err != nil {
					// Handle the error here, for example, log it or return from the function
					fmt.Println("Error converting cell to float:", err)
					panic(err)
				}
				summPropertyAmount += cashReported

				copy(propertyRecord.PropAmountReported[:], padOrTruncate(fmt.Sprintf("%010d", int64(cashReported*100)), len(propertyRecord.PropAmountReported))) // use 10 digit padded string in cents not dollars

				copy(propertyRecord.PropAmountAdvertised[:], padOrTruncate(row.Cells[columnMap["Cash Deduction"]].String(), len(propertyRecord.PropAmountAdvertised)))

				// Stock and CUSIP fields
				copy(propertyRecord.PropStockIssueName[:], padOrTruncate(row.Cells[columnMap["Stock Issue Name"]].String(), len(propertyRecord.PropStockIssueName)))
				copy(propertyRecord.PropStockCUSIP[:], padOrTruncate(row.Cells[columnMap["Stock CUSIP"]].String(), len(propertyRecord.PropStockCUSIP)))
				copy(propertyRecord.PropStockIssueName[:], padOrTruncate(row.Cells[columnMap["Stock Ticker Symbol"]].String(), len(propertyRecord.PropStockIssueName)))

				// Check Number and Account Number fields
				copy(propertyRecord.PropAccountNumber[:], padOrTruncate(row.Cells[columnMap["Account Number"]].String(), len(propertyRecord.PropAccountNumber)))
				copy(propertyRecord.PropCheckNumber[:], padOrTruncate(row.Cells[columnMap["Check Number"]].String(), len(propertyRecord.PropCheckNumber)))

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

func createNAUPATxtFile(fileName string, holderRecord HolderRecord, propertyRecords []PropertyRecord, summaryRecord SummaryRecord) error {
	// Create a new file or overwrite if exists
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	// Write the holder record
	formattedHolderRecord := formatHolderRecord(holderRecord)

	_, err = file.Write(formattedHolderRecord)
	if err != nil {
		return fmt.Errorf("failed to write holder record: %v", err)
	}

	// Write each property record
	for _, propertyRecord := range propertyRecords {
		// fmt.Println("Writing property record to file...", propertyRecord)
		// fmt.Println("propertyRecord: ", propertyRecord)

		formattedPropertyRecord := formatPropertyRecord(propertyRecord)

		_, err = file.Write(formattedPropertyRecord)
		if err != nil {
			return fmt.Errorf("failed to write property record: %v", err)
		}
	}

	formattedSummaryRecord := formatSummaryRecord(summaryRecord)

	_, err = file.Write(formattedSummaryRecord)
	if err != nil {
		return fmt.Errorf("failed to write summary record: %v", err)
	}

	return nil
}

// Helper function to format a Contact struct
func formatContact(buffer *bytes.Buffer, contact Contact) {
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

// Format the HolderRecord by looping through its bytes
func formatHolderRecord(holderRecord HolderRecord) []byte {
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
func formatPropertyRecord(propertyRecord PropertyRecord) []byte {
	// Use a bytes buffer to collect all the bytes
	var buffer bytes.Buffer

	// Append all the fields to the buffer one by one
	buffer.WriteByte(propertyRecord.TRCode)
	buffer.Write(propertyRecord.PropSequenceNumber[:])
	buffer.WriteByte(propertyRecord.PropOwnerType)
	buffer.WriteByte(propertyRecord.PropNameID)
	buffer.Write(propertyRecord.PropOwnerNameLast[:])
	buffer.Write(propertyRecord.PropOwnerNameFirst[:])
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
func formatSummaryRecord(summaryRecord SummaryRecord) []byte {
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
