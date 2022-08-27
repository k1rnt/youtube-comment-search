package search

import "strings"

func Keyword(comments []string, keyword string) []string {
	var result []string
	for _, comment := range comments {
		if strings.Contains(comment, keyword) {
			result = append(result, comment)
		}
	}
	return result
}
