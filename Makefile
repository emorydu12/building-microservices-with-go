cucumber:
	docker compose up -d
	cd features && godog ./
	docker compose stop

run:
	docker compose up -d
	go run main.go
	docker compose stop

unit:
	go test -cover ./... -race -v ./...

test: unit cucumber

package:
	go get -t ./...
	go get github.com/cucumber/godog/cmd/godog