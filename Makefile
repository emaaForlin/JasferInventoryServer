check_install:
	which ./swagger || curl -o swagger -L'#' "https://github.com/go-swagger/go-swagger/releases/download/v0.29.0/swagger_linux_amd64" && chmod +x ./swagger

swagger: check_install
	./swagger generate spec -o ./swagger.yaml --scan-models

# This test runs only in the CI workflow, if you run it in local maybe fail 
test:
	go test */*_test.go

build: swagger
	go build -o ./bin/app