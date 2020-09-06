package commands

import (
	"regexp"
	"strings"
)

func filterReplaceStrings(lines []string, old, new string) []string {
	var newLines []string

	for _, line := range lines {
		newLine := strings.Replace(line, old, new, -1)
		newLines = append(newLines, newLine)

	}
	return newLines
}

func filterReplacePrefixString(lines []string, old, new string) []string {
	var newLines []string

	for _, line := range lines {
		newLine := line
		if strings.HasPrefix(line, old) {
			newLine = strings.Replace(line, old, new, 1)
		}
		newLines = append(newLines, newLine)
	}
	return newLines
}

func filterLintList(lines []string) []string {
	var newLines []string
	for i, line := range lines {
		trimedLine := strings.TrimSpace(line)
		if strings.HasPrefix(trimedLine, "- ") {
			previousTrimedLine := strings.TrimSpace(lines[i-1])
			if !strings.HasPrefix(previousTrimedLine, "- ") && previousTrimedLine != "" {
				newLines = append(newLines, "")
			}
		}
		newLines = append(newLines, line)
	}
	return newLines
}

func filterAllRegexReplace(lines []string, regexPattern *regexp.Regexp, to string) []string {
	var newLines []string

	for _, line := range lines {
		newLine := regexPattern.ReplaceAllString(line, to)
		newLines = append(newLines, newLine)

	}
	return newLines
}
