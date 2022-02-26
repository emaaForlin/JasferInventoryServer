check_install:
	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger

swagger: check_install
	swagger generate spec -o ./swagger.yaml --scan-models

# This test runs only in the CI workflow, if you run it in local maybe fail 
test: swagger
	go test */*_test.go
