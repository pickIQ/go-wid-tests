package iteration

import "strings"

const defaultRepeatCount = 5

// Repeat takes char string(character which we want to repeat) and repeatTime(number of how many time we want repeat char. Default value: 5)
func Repeat(char string, repeatTime int) string {
	repeatCount := defaultRepeatCount

	if repeatTime >= 1 {
		repeatCount = repeatTime
	}

	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(char)
	}

	return repeated.String()
}