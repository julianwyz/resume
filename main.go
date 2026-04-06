package main

import (
	"fmt"
	"os"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

const htmlTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Julian Wyzykowski - Resume</title>
<link rel="stylesheet" href="resume.css">
</head>
<body>
<div class="page">
%s
</div>
</body>
</html>`

func main() {
	src, err := os.ReadFile("resume.md")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading resume.md: %v\n", err)
		os.Exit(1)
	}

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)

	opts := html.RendererOptions{
		Flags: html.CommonFlags | html.HrefTargetBlank,
	}
	renderer := html.NewRenderer(opts)

	body := markdown.ToHTML(src, p, renderer)

	out := fmt.Sprintf(htmlTemplate, string(body))
	if err := os.WriteFile("resume.html", []byte(out), 0o644); err != nil {
		fmt.Fprintf(os.Stderr, "error writing resume.html: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("resume.html generated")
}
