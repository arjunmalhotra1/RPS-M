.PHONY: test start

test:
	go test -v ./
  
start:
	go run main.go
