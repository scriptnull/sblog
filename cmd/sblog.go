package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/scriptnull/sblog/util"
)

func printHelp() {
	fmt.Printf(`
sblog [title] "tag1,tag2...tagn"

Starts a new blog post.
`)
}

func openInBrowser(url string) error {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	return err
}

func main() {
	if len(os.Args) <= 1 {
		os.Stderr.WriteString("missing arguments\n")
		printHelp()
		os.Exit(1)
	}

	title := os.Args[1]
	if len(title) == 0 {
		os.Stderr.WriteString("title needs to be a non-empty string\n")
		printHelp()
		os.Exit(1)
	}

	fileName := util.GenerateMarkdownFileName(title)

	var tags []string
	if len(os.Args) == 3 {
		tags = strings.Split(os.Args[2], ",")
	}

	baseURL := "https://github.com/scriptnull/vishnubharathi.codes/new/master/source/_posts"

	frontMatter := util.GenerateFrontMatter(title, tags)
	url, err := util.GenerateURL(baseURL, fileName, frontMatter)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Error while generating URL: %s", err))
		os.Exit(1)
	}

	if err := openInBrowser(url); err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Error while opening URL in browser: %s", err))
		os.Exit(1)
	}
}
