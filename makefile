build:
	mkdir golang-restapi-tasks
	cp ./data.json ./golang-restapi-tasks
	go build -o ./golang-restapi-tasks main.go

run:
	./golang-restapi-tasks/main

dev:
	go run middleware.go handler.go tasks.go server.go main.go