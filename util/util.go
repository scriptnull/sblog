package util

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

// GenerateFrontMatter creates the YAML front matter for a new blog post
func GenerateFrontMatter(title string, tags []string) string {
	t := time.Now()
	timeString := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	for i := range tags {
		tags[i] = fmt.Sprintf("\"%s\"", tags[i])

		if i != len(tags)-1 {
			tags[i] = fmt.Sprintf("%s,", tags[i])
		}
	}

	generatedFrontMatter := fmt.Sprintf(`---
title: %s
date: %s
tags: [%s]
---
`, title, timeString, strings.Join(tags, ""))

	return generatedFrontMatter
}

// GenerateMarkdownFileName generates a markdown filename
func GenerateMarkdownFileName(title string) string {
	words := strings.Split(title, " ")
	for i := range words {
		words[i] = strings.ToLower(strings.Trim(words[i], " "))
	}
	hyphenatedTitle := strings.Join(words, "-")
	return fmt.Sprintf("%s.md", hyphenatedTitle)
}

// GenerateURL generates a Github URL that could be opened for writing a new blog post
func GenerateURL(base, fileName, value string) (string, error) {
	u, err := url.Parse(base)
	if err != nil {
		return "", err
	}

	q := u.Query()
	q.Set("filename", fileName)
	q.Set("value", value)
	u.RawQuery = q.Encode()

	return u.String(), nil
}
