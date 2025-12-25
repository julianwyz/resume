default: build

build:
	go tool easypdf convert -f resume.md --css resume.css -o resume.pdf