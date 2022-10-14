gen:
	protoc --go_out=. --proto_path=proto proto/*.proto

clean:
	rm pb/*.go

run:
	go run cmd/main.go

test:
	go test -cover -race ./...