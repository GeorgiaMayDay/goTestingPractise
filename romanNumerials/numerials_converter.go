package romanNumerials

import (
	"fmt"
	"strings"
)

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ValueOf(symbols ...byte) int {
	symbol := string(symbols)
	for _, s := range allRomanNumerals {
		if s.Symbol == symbol {
			return s.Value
		}
	}

	return 0
}

func couldBeSubtractive(position int, currentChar byte, fullNum string) bool {
	return position+1 < len(fullNum) && (currentChar == 'I' || currentChar == 'X' || currentChar == 'C')
}

func ConvertToRomans(arabicNum int) string {
	var finalNum strings.Builder
	for _, numeral := range allRomanNumerals {
		for arabicNum >= numeral.Value {
			finalNum.WriteString(numeral.Symbol)
			arabicNum -= numeral.Value
		}
	}
	return finalNum.String()
}

func ConvertToArabic(roman string) int {
	total := 0
	if roman == "XXXIX" {
		fmt.Println(total)
	}

	for i := 0; i < len(roman); i++ {
		symbol := roman[i]

		//check if there is a number pair possibel for this one
		if couldBeSubtractive(i, symbol, roman) {
			nextSymbol := roman[i+1]
			//check if it is a number pair
			check := ValueOf(symbol, nextSymbol)

			if check != 0 {
				total += check
				i++
			} else {
				total += ValueOf(symbol)
			}

		} else {
			total += ValueOf(symbol)
		}
	}
	return total
}
