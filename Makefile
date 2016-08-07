PROGRAM_NAME=toy-dns-nameserver
default: lint test

lint:
	golint ./... | egrep -v -f lint.ignores || true

test:
	go test -cover -v ./...

deps:
	go get github.com/golang/lint/golint
	go get gopkg.in/check.v1
	go get golang.org/x/tools/cmd/cover

ci: lint test

run: build run_program

build:
	go build

run_program:
	./$(PROGRAM_NAME)

clean:
	rm $(PROGRAM_NAME)
