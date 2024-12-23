get:
	go get ./...

buildserver:
	echo "Build API"
	CGO_ENABLED=0 GOOS=linux go build -o ./server/kittenserver ./server/main.go

runserver:
	echo "Run API"
	docker compose -p kittenserver -f ./docker-compose.yaml up --build	