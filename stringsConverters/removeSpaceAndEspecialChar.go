package stringsconverters

import (
	"fmt"
	"regexp"
	"strings"
)

func RemoveSpaceAndEpecialChar() {

	str := "4029234093840 3984924 39482934 3.4.5.5.6.34.34.343.4"

	semEspaço := strings.ReplaceAll(str, " ", "")
	formatada := strings.ReplaceAll(semEspaço, ".", "")
	fmt.Printf("Replace: %s \n", formatada)

	modified := func(r rune) rune {
		if r == ' ' || r == '.' {
			return -1
		}
		return r
	}

	result := strings.Map(modified, str)

	fmt.Printf("Map: %s \n", result)

	fmt.Println(result == strings.TrimSpace(formatada))

	fmt.Printf("%s\n%s", result, formatada)

}

func RemoveEspcifiedChars(toCovert string, chars ...rune) {

	filter := func(r rune) rune {
		for _, char := range chars {
			if r == char {
				return -1
			}
		}
		return r
	}

	result := strings.Map(filter, toCovert)

	fmt.Printf("Map: %s \n", result)

}

func RemoveNonNumber(toConvert string) string {
	NumbersOnlyRegexp := regexp.MustCompile(`[^0-9]+`)

	return NumbersOnlyRegexp.ReplaceAllString(toConvert, "")

}
