package utils

import (
	"strings"
)

func DecodeValue(text string) string {
	text = strings.ReplaceAll(text, "2@c86cb3", "'")
	text = strings.ReplaceAll(text, "2#c86cb3", "`")
	return text
}

func EncodeValue(text string) string {
	text = strings.ReplaceAll(text, "'", "2@c86cb3")
	text = strings.ReplaceAll(text, "`", "2#c86cb3")
	return text
}
