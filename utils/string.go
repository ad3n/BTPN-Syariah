package utils

import (
	"regexp"
	"strings"
)

func ToUnderScore(text string) string {
	camel := regexp.MustCompile("(^[^A-Z0-9]*|[A-Z0-9]*)([A-Z0-9][^A-Z]+|$)")
	result := []string{}

	for _, sub := range camel.FindAllStringSubmatch(text, -1) {
		if sub[1] != "" {
			result = append(result, sub[1])
		}

		if sub[2] != "" {
			result = append(result, sub[2])
		}
	}

	return strings.ToLower(strings.Join(result, "_"))
}
