package iteration

import "strings"

const repeatCount = 5

// Repeat takes char string(character which we want to repeat) and repeatTime(number of how many time we want repeat char)
func Repeat(char string, repeatTime int) string {
	var repeated strings.Builder
	for i := 0; i < repeatTime; i++ {
		repeated.WriteString(char)
	}

	return repeated.String()
}