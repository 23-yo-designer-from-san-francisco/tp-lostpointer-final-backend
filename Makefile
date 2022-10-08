.PHONY: server
server:
	go build -tags=jsoniter -o bin/api/autfinal -v ./cmd/server 

.PHONY: swagger
swagger:
	swag init -g ./cmd/server/main.go
