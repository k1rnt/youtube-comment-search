package search

import (
	"regexp"
	"strings"
)

func Keyword(comments []string, keyword string) []string {
	var result []string
	for _, comment := range comments {
		if strings.Contains(comment, keyword) {
			result = append(result, comment)
		}
	}
	return result
}

func Regex(comments []string, regex string) []string {
	var result []string
	for _, comment := range comments {
		re := regexp.MustCompile(regex)
		if re.MatchString(comment) {
			result = append(result, comment)
		}
	}
	return result
}
