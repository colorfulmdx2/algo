package __array

import (
	"strconv"
	"strings"
)

type Codec struct {
	b strings.Builder
}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	defer codec.b.Reset()

	for _, word := range strs {
		codec.b.WriteString(strconv.Itoa(len(word)))
		codec.b.WriteRune('|')
		codec.b.WriteString(word)
	}

	return codec.b.String()
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
	var result []string

	for i := 0; i < len(strs); {
		wordStartIndex := i
		wordEndIndex := i

		for strs[wordEndIndex] != '|' {
			wordEndIndex++
		}

		var nextWordLen int
		dec := 1

		for j := wordEndIndex - 1; j >= wordStartIndex; j-- {
			nextWordLen += (int(strs[j]) - 48) * dec
			dec *= 10
		}

		start := wordEndIndex + 1
		end := start + nextWordLen

		result = append(result, string(strs[start:end]))

		i = end
	}

	return result
}

func (codec *Codec) Decode1(strs string) []string {
	var words []string
	i := 0

	for i < len(strs) {
		// Find the position of the delimiter '|'
		delimiterIndex := strings.IndexByte(strs[i:], '|') + i
		// Extract the length as an integer
		length, _ := strconv.Atoi(strs[i:delimiterIndex])
		// Calculate the start and end positions of the word
		start := delimiterIndex + 1
		end := start + length
		// Append the word to the result slice
		words = append(words, strs[start:end])
		// Move the index to the end of the current word
		i = end
	}

	return words
}

// Your Codec object will be instantiated and called as such:
// var codec Codec
// codec.Decode(codec.Encode(strs));
