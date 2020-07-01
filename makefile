build:
	mkdir bin
	cp data.json bin
	go build -o bin/golang-restapi-tasks -v .

run:
	bin/golang-restapi-tasks

dev:
	go run -v .