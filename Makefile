.PHONY: run debug

run:
	go run main.go
debug:
	dlv debug .
