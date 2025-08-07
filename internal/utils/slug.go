package utils

import (
	"regexp"
	"strings"
)

func GenerateSlug(input string) string {
	slug := strings.ToLower(input)

	re := regexp.MustCompile(`[^a-z0-9\s-]`)
	slug = re.ReplaceAllString(slug, "")

	slug = strings.ReplaceAll(slug, " ", "-")

	slug = regexp.MustCompile(`-+`).ReplaceAllString(slug, "-")
	return strings.Trim(slug, "-")
}
