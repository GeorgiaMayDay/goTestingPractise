package iteration

func Repeat(toBeRepeated string, timesRepeated int) string {
	var repeated string = ""
	for i := 0; i < timesRepeated; i++ {
		repeated = repeated + toBeRepeated
	}
	return repeated
}
