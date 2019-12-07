package iteration

import "strings"

func Repeat(character string, repeatCount int) string {
	var repeated string
	repeated = strings.Repeat(character, repeatCount)
	return repeated
}
