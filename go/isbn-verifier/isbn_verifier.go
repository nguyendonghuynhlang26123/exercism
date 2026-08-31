package isbnverifier

import (
	"strings"
	"unicode"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn,"-", "")

	if len(isbn) != 10 {
		return false
	}
	println(isbn)

	total := 0
	for i, r := range isbn {
		var d int
		if r == 'X' {
			if i != len(isbn) - 1 {
				return false
			}
			d = 10
		} else {
			if !unicode.IsDigit(r) {
				return false
			}
			d = int(r - '0')
		}

		total += d * (10 - i)
	}

	println()
	return total % 11 == 0
}
