check_install:
	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger

swagger: check_install
	swagger generate spec -o ./swagger.yaml --scan-models

test:
	go test */*_test.go

build: test swagger
	go build -o bin/app