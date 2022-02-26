check_install:
	which swagger || download_url="https://github.com/go-swagger/go-swagger/releases/download/v0.29.0/swagger_linux_amd64" | .browser_download_url') | curl -o /usr/local/bin/swagger -L'#' ${download_url} | chmod +x /usr/local/bin/swagger

swagger: check_install
	swagger generate spec -o ./swagger.yaml --scan-models

# This test runs only in the CI workflow, if you run it in local maybe fail 
test: swagger
	go test */*_test.go
