package country

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Map to store country names and their corresponding 3-digit Alpha-3 codes
var CountryCodes = map[string]string{
	"Afghanistan":                       "AFG",
	"Albania":                           "ALB",
	"Algeria":                           "DZA",
	"American Samoa":                    "ASM",
	"Andorra":                           "AND",
	"Angola":                            "AGO",
	"Anguilla":                          "AIA",
	"Antarctica":                        "ATA",
	"Antigua and Barbuda":               "ATG",
	"Argentina":                         "ARG",
	"Armenia":                           "ARM",
	"Aruba":                             "ABW",
	"Australia":                         "AUS",
	"Austria":                           "AUT",
	"Azerbaijan":                        "AZE",
	"Bahamas":                           "BHS",
	"Bahrain":                           "BHR",
	"Bangladesh":                        "BGD",
	"Barbados":                          "BRB",
	"Belarus":                           "BLR",
	"Belgium":                           "BEL",
	"Belize":                            "BLZ",
	"Benin":                             "BEN",
	"Bermuda":                           "BMU",
	"Bhutan":                            "BTN",
	"Bolivia":                           "BOL",
	"Bonaire, Sint Eustatius and Saba":  "BES",
	"Bosnia and Herzegovina":            "BIH",
	"Botswana":                          "BWA",
	"Brazil":                            "BRA",
	"British Indian Ocean Territory":    "IOT",
	"Brunei Darussalam":                 "BRN",
	"Bulgaria":                          "BGR",
	"Burkina Faso":                      "BFA",
	"Burundi":                           "BDI",
	"Cabo Verde":                        "CPV",
	"Cambodia":                          "KHM",
	"Cameroon":                          "CMR",
	"Canada":                            "CAN",
	"Cayman Islands":                    "CYM",
	"Central African Republic":          "CAF",
	"Chad":                              "TCD",
	"Chile":                             "CHL",
	"China":                             "CHN",
	"Christmas Island":                  "CXR",
	"Cocos (Keeling) Islands":           "CCK",
	"Colombia":                          "COL",
	"Comoros":                           "COM",
	"Congo (Democratic Republic)":       "COD",
	"Congo":                             "COG",
	"Cook Islands":                      "COK",
	"Costa Rica":                        "CRI",
	"Croatia":                           "HRV",
	"Cuba":                              "CUB",
	"Curaçao":                           "CUW",
	"Cyprus":                            "CYP",
	"Czechia":                           "CZE",
	"Côte d'Ivoire":                     "CIV",
	"Denmark":                           "DNK",
	"Djibouti":                          "DJI",
	"Dominica":                          "DMA",
	"Dominican Republic":                "DOM",
	"Ecuador":                           "ECU",
	"Egypt":                             "EGY",
	"El Salvador":                       "SLV",
	"Equatorial Guinea":                 "GNQ",
	"Eritrea":                           "ERI",
	"Estonia":                           "EST",
	"Eswatini":                          "SWZ",
	"Ethiopia":                          "ETH",
	"Falkland Islands [Malvinas]":       "FLK",
	"Faroe Islands":                     "FRO",
	"Fiji":                              "FJI",
	"Finland":                           "FIN",
	"France":                            "FRA",
	"French Guiana":                     "GUF",
	"French Polynesia":                  "PYF",
	"French Southern Territories":       "ATF",
	"Gabon":                             "GAB",
	"Gambia":                            "GMB",
	"Georgia":                           "GEO",
	"Germany":                           "DEU",
	"Ghana":                             "GHA",
	"Gibraltar":                         "GIB",
	"Greece":                            "GRC",
	"Greenland":                         "GRL",
	"Grenada":                           "GRD",
	"Guadeloupe":                        "GLP",
	"Guam":                              "GUM",
	"Guatemala":                         "GTM",
	"Guernsey":                          "GGY",
	"Guinea":                            "GIN",
	"Guinea-Bissau":                     "GNB",
	"Guyana":                            "GUY",
	"Haiti":                             "HTI",
	"Heard Island and McDonald Islands": "HMD",
	"Holy See":                          "VAT",
	"Honduras":                          "HND",
	"Hong Kong":                         "HKG",
	"Hungary":                           "HUN",
	"Iceland":                           "ISL",
	"India":                             "IND",
	"Indonesia":                         "IDN",
	"Iran":                              "IRN",
	"Iraq":                              "IRQ",
	"Ireland":                           "IRL",
	"Isle of Man":                       "IMN",
	"Israel":                            "ISR",
	"Italy":                             "ITA",
	"Jamaica":                           "JAM",
	"Japan":                             "JPN",
	"Jersey":                            "JEY",
	"Jordan":                            "JOR",
	"Kazakhstan":                        "KAZ",
	"Kenya":                             "KEN",
	"Kiribati":                          "KIR",
	"Korea (North)":                     "PRK",
	"Korea, Republic Of":                "KOR",
	"Korea (South)":                     "KOR",
	"Kuwait":                            "KWT",
	"Kyrgyzstan":                        "KGZ",
	"Lao People's Democratic Republic":  "LAO",
	"Latvia":                            "LVA",
	"Lebanon":                           "LBN",
	"Lesotho":                           "LSO",
	"Liberia":                           "LBR",
	"Libya":                             "LBY",
	"Liechtenstein":                     "LIE",
	"Lithuania":                         "LTU",
	"Luxembourg":                        "LUX",
	"Macao":                             "MAC",
	"Madagascar":                        "MDG",
	"Malawi":                            "MWI",
	"Malaysia":                          "MYS",
	"Maldives":                          "MDV",
	"Mali":                              "MLI",
	"Malta":                             "MLT",
	"Marshall Islands":                  "MHL",
	"Martinique":                        "MTQ",
	"Mauritania":                        "MRT",
	"Mauritius":                         "MUS",
	"Mayotte":                           "MYT",
	"Mexico":                            "MEX",
	"Micronesia":                        "FSM",
	"Moldova":                           "MDA",
	"Monaco":                            "MCO",
	"Mongolia":                          "MNG",
	"Montenegro":                        "MNE",
	"Montserrat":                        "MSR",
	"Morocco":                           "MAR",
	"Mozambique":                        "MOZ",
	"Myanmar":                           "MMR",
	"Namibia":                           "NAM",
	"Nauru":                             "NRU",
	"Nepal":                             "NPL",
	"Netherlands":                       "NLD",
	"New Caledonia":                     "NCL",
	"New Zealand":                       "NZL",
	"Nicaragua":                         "NIC",
	"Niger":                             "NER",
	"Nigeria":                           "NGA",
	"Niue":                              "NIU",
	"Norfolk Island":                    "NFK",
	"Northern Mariana Islands":          "MNP",
	"Norway":                            "NOR",
	"Oman":                              "OMN",
	"Pakistan":                          "PAK",
	"Palau":                             "PLW",
	"Palestine, State of":               "PSE",
	"Panama":                            "PAN",
	"Papua New Guinea":                  "PNG",
	"Paraguay":                          "PRY",
	"Peru":                              "PER",
	"Philippines":                       "PHL",
	"Pitcairn":                          "PCN",
	"Poland":                            "POL",
	"Portugal":                          "PRT",
	"Puerto Rico":                       "PRI",
	"Qatar":                             "QAT",
	"Romania":                           "ROU",
	"Russia":                            "RUS",
	"Russian Federation":                "RUS",
	"Rwanda":                            "RWA",
	"Réunion":                           "REU",
	"Saint Barthélemy":                  "BLM",
	"Saint Helena":                      "SHN",
	"Saint Kitts and Nevis":             "KNA",
	"Saint Lucia":                       "LCA",
	"Saint Martin (French part)":        "MAF",
	"Saint Pierre and Miquelon":         "SPM",
	"Saint Vincent and the Grenadines":  "VCT",
	"Samoa":                             "WSM",
	"San Marino":                        "SMR",
	"Sao Tome and Principe":             "STP",
	"Saudi Arabia":                      "SAU",
	"Senegal":                           "SEN",
	"Serbia":                            "SRB",
	"Seychelles":                        "SYC",
	"Sierra Leone":                      "SLE",
	"Singapore":                         "SGP",
	"Sint Maarten (Dutch part)":         "SXM",
	"Slovakia":                          "SVK",
	"Slovenia":                          "SVN",
	"Solomon Islands":                   "SLB",
	"Somalia":                           "SOM",
	"South Africa":                      "ZAF",
	"South Georgia and the South Sandwich Islands": "SGS",
	"South Sudan":              "SSD",
	"Spain":                    "ESP",
	"Sri Lanka":                "LKA",
	"Sudan":                    "SDN",
	"Suriname":                 "SUR",
	"Svalbard and Jan Mayen":   "SJM",
	"Sweden":                   "SWE",
	"Switzerland":              "CHE",
	"Syria":                    "SYR",
	"Taiwan":                   "TWN",
	"Tajikistan":               "TJK",
	"Tanzania":                 "TZA",
	"Thailand":                 "THA",
	"Timor-Leste":              "TLS",
	"Togo":                     "TGO",
	"Tokelau":                  "TKL",
	"Tonga":                    "TON",
	"Trinidad and Tobago":      "TTO",
	"Tunisia":                  "TUN",
	"Turkey":                   "TUR",
	"Turkmenistan":             "TKM",
	"Tuvalu":                   "TUV",
	"Uganda":                   "UGA",
	"Ukraine":                  "UKR",
	"United Arab Emirates":     "ARE",
	"United Kingdom":           "GBR",
	"United States":            "USA",
	"Uruguay":                  "URY",
	"Uzbekistan":               "UZB",
	"Vanuatu":                  "VUT",
	"Venezuela":                "VEN",
	"Vietnam":                  "VNM",
	"Virgin Islands (British)": "VGB",
	"Virgin Islands (U.S.)":    "VIR",
	"Wallis and Futuna":        "WLF",
	"Western Sahara":           "ESH",
	"Yemen":                    "YEM",
	"Zambia":                   "ZMB",
	"Zimbabwe":                 "ZWE",
	"Åland Islands":            "ALA",
}

// Function to look up the 3-digit country code
func LookupCountryCode(country string) string {
	// Normalize the country name (trim spaces and convert to title case using cases.Title)
	titleCaser := cases.Title(language.English) // Create a caser for the English language
	countryTrimmed := strings.TrimSpace(titleCaser.String(strings.ToLower(country)))

	if code, exists := CountryCodes[countryTrimmed]; exists {
		return code
	}

	fmt.Println("Country not found:", "START", country, "MID", countryTrimmed, "END")
	// If the country is not found, return "XXX" (could also return an error or a default value)
	return "XXX"
}
