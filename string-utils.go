package simpleutils

import (
	"fmt"
	"strings"
	"unicode"
)

//Charget returns the rune at the logical index position
func CharGet(input string, position int) rune {
	chars := []rune(input)
	if position < 0 || position >= len(chars) {
		return 0
	} else {
		return chars[position]
	}
}

//PrintableStr converts non-printable characters in input to a dot
func PrintableStr(input string) string {
	output := ""
	for _, c := range input {
		if unicode.IsPrint(c) {
			output += string(c)
		} else {
			output += "."
		}
	}

	return output
}

//Substr returns a string based on logical character start and length,
//instead of bytes like a slice
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

//IsTokenCharFirst returns true if ch is a letter or underscore
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

//IsTokenCharNext returns true if ch is a letter, number or underscore
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

//IsTokenName returns true if s contains only letters, numbers or underscores,
//and does not start with a number, and has at least one letter
func IsTokenName(s string) bool {
	first := true
	hasLetter := false
	for _, ch := range s {
		if first {
			if !IsTokenCharFirst(ch) {
				return false
			}
			first = false
		} else if !IsTokenCharNext(ch) {
			return false
		}

		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')  {
			hasLetter = true
		}
	}
	return hasLetter
}

//IsTokenNameWithMiddleChars returns true if s contains only letters, numbers or underscores,
//or has allowed characters that are not first or last, does not start with a number, and has
//at least one letter. A common middleChars string is "-".
func IsTokenNameWithMiddleChars(s string, middleChars string) bool {
	first := true
	hasLetter := false
	lastCh := rune(0)

	for _, ch := range s {
		lastCh = ch
		if first {
			if !IsTokenCharFirst(ch) {
				return false
			}
			first = false
		} else if !IsTokenCharNext(ch) {
			if strings.ContainsAny(string(ch), middleChars) {
				continue
			}
			return false
		}

		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')  {
			hasLetter = true
		}
	}

	if strings.ContainsAny(string(lastCh), middleChars) {
		return false
	}

	return hasLetter
}

//Escape translates control characters to backslash escape sequence;
//e.g., '\r' becomes `\r`
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

//WhichSuffix returns a pointer to the first suffix matching s, or
//nil if none of the suffixes match
func WhichSuffix(s string, suffixes ...string) *string {
	for _, suffix := range suffixes {
		if strings.HasSuffix(s, suffix) {
			return &suffix
		}
	}
	return nil
}

func patternMatchWorker(pattern []rune, testString []rune) bool {
	pos := 0

	for i, c := range pattern {
		if c == '*' {
			subpattern := pattern[i+1:]
			if len(subpattern) == 0 {
				return true
			}

			for pos < len(testString) {
				if patternMatchWorker(subpattern, testString[pos:]) {
					return true
				}
				pos++
			}

			return false
		} else if pos >= len(testString) || c != testString[pos] {
			return false
		}
		pos++
	}

	return pos == len(testString)
}

//PatternMatch performs a simple asterisk-based pattern match
func PatternMatch(pattern string, testString string) bool {
	p := []rune(pattern)
	ts := []rune(testString)
	return patternMatchWorker(p, ts)
}

func _getByte(data []byte, pos int) (byte, bool) {
	if pos >= len(data) {
		return 0, false
	}
	return data[pos], true
}

//Utf8len returns the number of bytes in the logical UTF-8 character,
//0 if the character is incomplete, or -1 if the character is invalid,
func Utf8len(data []byte, offset int) int {
	b, valid := _getByte(data, offset)
	if !valid {
		return 0		// incomplete
	}

	if b < 0x80 {
		return 1		// common 0-127 case
	}

	// Multi-byte detected.
	// The head byte indicates the length.
	// 110xxxxx = 2 bytes, 1110xxxx = 3 bytes, 11110xxx = 4 bytes
	var len int
	if b & 0b11100000 == 0b11000000 {
		len = 2
	} else if b & 0b11110000 == 0b11100000 {
		len = 3
	} else if b & 0b11111000 == 0b11110000 {
		len = 4
	} else {
		return -1		// invalid
	}

	// each trail byte must be in 10xxxxxx form
	for i := 1 ; i < len ; i++ {
		b, valid = _getByte(data, offset + i)
		if !valid {
			return 0		// incomplete
		}
		if b & 0b11000000 != 0b10000000 {
			return -1
		}
	}

	return len
}

//StringArrayToStrings converts an array of strings to a single
//string, placing the delimiter string between each
func StringArrayToString(strs []string, delimiter string) string {
	var sb strings.Builder

	for _,str := range strs {
		if sb.Len() > 0 {
			sb.WriteString(delimiter)
		}
		sb.WriteString(str)
	}

	return sb.String()
}


// IndexOf is like strings.Index with a starting index
func IndexOf(testString, substring string, startingIndex int) int {
	return strings.Index(testString[startingIndex:], substring)
}

// IndexOfAny is like strings.Index with a starting index
func IndexOfAny(testString, chars string, startingIndex int) int {
	return strings.IndexAny(testString[startingIndex:], chars)
}