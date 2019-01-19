.PHONY: 
	update
	generate


update:
	go mod tidy

generate:
	go generate ./...
