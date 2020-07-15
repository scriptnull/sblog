package util

import (
	"strings"
	"testing"
)

func TestGenerateURL(t *testing.T) {
	base := "https://github.com/scriptnull/vishnubharathi.codes/new/master"
	fileName := "source/_posts/example.md"
	value := "example content"

	url, err := GenerateURL(base, fileName, value)
	if err != nil {
		t.Errorf("Expects to not error out, err: %s", err)
	}

	if url != "https://github.com/scriptnull/vishnubharathi.codes/new/master?filename=source%2F_posts%2Fexample.md&value=example+content" {
		t.Errorf("Expects correct URL to be generated")
	}
}

func TestGenerateFrontMatter(t *testing.T) {
	tags := []string{"example", "blog"}
	fm := GenerateFrontMatter("example", tags)

	if !strings.HasPrefix(fm, "---\n") {
		t.Error("Expects --- as first line of front matter")
	}

	if !strings.HasSuffix(fm, "---\n") {
		t.Error("Expects --- as last line of front matter")
	}

	if !strings.Contains(fm, "title: example") {
		t.Error("Expects title to be present in front matter")
	}

	if !strings.Contains(fm, `tags: ["example","blog"]`) {
		t.Error("Expects tags to be present in front matter")
	}
}
