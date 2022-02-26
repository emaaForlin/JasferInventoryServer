check_install:
	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger

swagger: check_install
	swagger generate spec -o ./swagger.yaml --scan-models

# This test runs only in the CI workflow, if you run it in local maybe fail 
test:
	go test */*_test.go DB_HOST=${TEST_DB_HOST} DB_PORT=${TEST_DB_PORT} DB_USER=${TEST_DB_USER} DB_PASS=${TEST_DB_PASS} DB_NAME=${DB_NAME}
