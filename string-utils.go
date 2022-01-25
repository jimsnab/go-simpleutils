package simpleutils

import (
	"fmt"
	"strings"
)

func Substr(input string, start int, length int) string {
	chars := []rune(input)
	inputLen := len(chars)

	if start >= inputLen {
		return ""
	}

	if start+length > inputLen {
		length = inputLen - start
	}

	return string(chars[start : start+length])
}

func IsTokenCharFirst(ch rune) bool {
	if ch >= 'a' && ch <= 'z' {
		return true
	}
	if ch >= 'A' && ch <= 'Z' {
		return true
	}
	if ch == '_' {
		return true
	}
	return false
}

func IsTokenCharNext(ch rune) bool {
	if ch >= 'a' && ch <= 'z' {
		return true
	}
	if ch >= 'A' && ch <= 'Z' {
		return true
	}
	if ch >= '0' && ch <= '9' {
		return true
	}
	if ch == '_' {
		return true
	}
	return false
}

func IsTokenName(s string) bool {
	if len(s) == 0 {
		return false
	}

	first := true
	for _, ch := range s {
		if first {
			if !IsTokenCharFirst(ch) {
				return false
			}
			first = false
		} else if !IsTokenCharNext(ch) {
			return false
		}
	}
	return true
}

func Escape(s string) string {
	const ctrlChars = "\r\n\t\b\""
	const ctrlLetters = "rntb\""
	var sb strings.Builder

	for _, ch := range s {
		ctrlIndex := strings.Index(ctrlChars, string(ch))
		if ctrlIndex >= 0 {
			sb.WriteRune('\\')
			sb.WriteRune(rune(ctrlLetters[ctrlIndex]))
		} else if ch < ' ' {
			sb.WriteString(fmt.Sprintf("\\x%02X", int(ch)))
		} else {
			sb.WriteRune(ch)
		}
	}

	return sb.String()
}

func WhichSuffix(s string, suffixes ...string) *string {
	for _, suffix := range suffixes {
		if strings.HasSuffix(s, suffix) {
			return &suffix
		}
	}
	return nil
}
