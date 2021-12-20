package utils

import "strings"

func ExtractContext(input string) [][]string {
	context := make([][]string, 0)

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		context = append(context, strings.Split(line, ","))
	}

	return context
}

func CompressContext(context [][]string) string {
	lines := make([]string, 0)

	for _, line := range context {
		lines = append(lines, strings.Join(line, ","))
	}

	return strings.Join(lines, "\n")
}
