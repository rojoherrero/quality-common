.PHONY: 
	update
	generate
	generate-security

update:
	go mod tidy

generate:
	go generate ./...

generate-security:
	protoc -I security/ security/security_service.proto --go_out=plugins=grpc:security
