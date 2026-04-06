.DEFAULT_GOAL := build

build: resume.html
	wkhtmltopdf --enable-local-file-access --print-media-type resume.html resume.pdf

resume.html: resume.md resume.css main.go
	go run main.go