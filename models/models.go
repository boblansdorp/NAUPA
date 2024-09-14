package models

const SoftwareVersion = "0.0.1"
const SoftwareCreator = "Bob Lansdorp"
const SoftwareCreatorContact = "Bob Lansdorp Consulting LLC, https://github.com/boblansdorp/NAUPA"

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
	HolderIncDateCCYY [4]byte  // 4 bytes, positions 29-32Q
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
