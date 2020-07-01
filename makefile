build:
	mkdir bin
	cp data.json bin
	go build -o bin/go-tasks middleware.go handler.go tasks.go server.go main.go

run:
	bin/go-tasks

dev:
	go run middleware.go handler.go tasks.go server.go main.go