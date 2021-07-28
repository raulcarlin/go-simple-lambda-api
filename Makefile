dependencies:
	go mod tidy

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/simple-get-lambda cmd/get/main.go
	env GOOS=linux go build -ldflags="-s -w -X main.ApiID=${LOCALSTACK_API_ID}" -o bin/simple-post-lambda cmd/post/main.go

clean:
	rm -rf ./bin 

remove:
	sls remove --verbose

deploy: clean build
	serverless deploy --stage local
