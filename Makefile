all: compile
pre:
	go get github.com/pterm/pterm
	go get github.com/pterm/pterm/putils
c:
	go build  -v ./...
test:
	go test -v
rebuild: pre c test
compile: c test
