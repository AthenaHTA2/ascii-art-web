package AsciiArt

import "strings"

func BubbleConv(inputword string, chars map[int]string) string {

	result := ""
	var splitbyNewLine []string

	if len(inputword) == 0 {
		return ""
	}

	if inputword == "\\n" {
		return "\n"
	}

	for i, j := 0, 0; j < len(inputword); i++ {
		if i+1 <= len(inputword)-1 {
			if inputword[i] == '\\' && inputword[i+1] == 'n' {
				splitbyNewLine = append(splitbyNewLine, inputword[j:i])
				j = i + 2
			}
		} else if i == len(inputword)-1 {
			splitbyNewLine = append(splitbyNewLine, inputword[j:])
		}
	}

	for _, b := range splitbyNewLine {
		var bubbleWriting []string

		for _, turntoBubbleWriting := range b {
			bubbleWriting = append(bubbleWriting, chars[int(turntoBubbleWriting)])
		}

		for linePos := 0; linePos < 8; linePos++ {
			for _, indCharacter := range bubbleWriting {
				splitLine := strings.Split(indCharacter, "\n")
				result += splitLine[linePos]

			}
			result += "\n"
		}
	}

	return result

}
