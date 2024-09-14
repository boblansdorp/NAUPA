package utils

import (
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// Function that fills a byte array with spaces
func FillByteArrayWithSpaces(arr []byte, length int) {
	for i := 0; i < length; i++ {
		arr[i] = ' ' // Fill each position in the array with a space character
	}
}

// Function that fills a byte array with spaces
func FillByteArrayWithZeroes(arr []byte, length int) {
	for i := 0; i < length; i++ {
		arr[i] = '0' // Fill each position in the array with a space character
	}
}

// Replace special non-ASCII characters with ASCII equivalents
func ReplaceSpecialChars(value string) string {
	replacements := map[rune]rune{
		'Ø': 'O', // Replace Ø with O
		'ø': 'o', // Replace ø with o
		'Æ': 'A', // Replace Æ with A (you could replace with AE if needed)
		'æ': 'a', // Replace æ with a
		'Å': 'A', // Replace Å with A
		'å': 'a', // Replace å with a
		'Ñ': 'N', // Replace Ñ with N
		'ñ': 'n', // Replace ñ with n
		'ß': 'B', // Replace ß with s
		// Add other replacements as needed
	}

	var output strings.Builder
	for _, r := range value {
		if replacement, exists := replacements[r]; exists {
			output.WriteRune(replacement)
		} else {
			output.WriteRune(r)
		}
	}
	return output.String()
}

// Function to check if the rune is a non-spacing mark (Mn category)
func IsMn(r rune) bool {
	return unicode.Is(unicode.Mn, r)
}

// Helper function to convert non-ASCII characters to their nearest ASCII equivalent
func ToASCII(value string) string {
	// Normalize the string to decompose characters (NFD form)
	t := transform.Chain(norm.NFD, transform.RemoveFunc(IsMn), norm.NFC)
	asciiStr, _, err := transform.String(t, value)
	if err != nil {
		return value // Return the original string if there's an error
	}

	// Replace specific non-ASCII characters with their ASCII equivalents
	asciiStr = ReplaceSpecialChars(asciiStr)

	return asciiStr
}

// Helper function to pad or truncate a string to the exact length (handling UTF-8 properly)
func PadOrTruncate(value string, length int) string {
	// Convert to nearest ASCII equivalent
	asciiValue := ToASCII(value)

	runes := []rune(asciiValue) // Convert to runes to handle UTF-8 characters correctly

	// If the string is longer than the required length, truncate it
	if len(runes) > length {
		return string(runes[:length])
	}

	// If the string is shorter, pad with spaces
	return fmt.Sprintf("%-*s", length, asciiValue)
}
