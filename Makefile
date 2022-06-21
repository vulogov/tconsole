all: compile
pre:
	go get github.com/lrita/cmap
	go get github.com/gammazero/deque
	go get github.com/pterm/pterm
	go get github.com/stretchr/testify/assert
	go get github.com/pterm/pterm/putils
c:
	go build  -v ./...
test:
	go test -v -coverprofile=coverage.out ./...
rebuild: pre c test
compile: c test
