package romanNumerials

import "strings"

func ConvertToRomans(arabicNum int) string {
	var finalNum strings.Builder
	for arabicNum > 0 {
		switch arabicNum {
		case 4:
			finalNum.WriteString("IV")
			arabicNum -= 4
			break
		case 9:
			finalNum.WriteString("IX")
			arabicNum -= 9
			break
		case 10:
			finalNum.WriteString("X")
			arabicNum -= 10
			break
		case 5, 6, 7, 8:
			finalNum.WriteString("V")
			arabicNum -= 5
			break
		default:
			finalNum.WriteString("I")
			arabicNum--
		}
	}
	return finalNum.String()
}
