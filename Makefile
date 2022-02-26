check_install:
	which swagger || GO111MODULE=off && apt install -y apt-transport-https gnupg curl \
	curl -1sLf 'https://dl.cloudsmith.io/public/go-swagger/go-swagger/gpg.2F8CB673971B5C9E.key' | apt-key add - \
	curl -1sLf 'https://dl.cloudsmith.io/public/go-swagger/go-swagger/config.deb.txt?distro=debian&codename=any-version' > /etc/apt/sources.list.d/go-swagger-go-swagger.list \
	apt update \
	apt install swagger

swagger: check_install
	swagger generate spec -o ./swagger.yaml --scan-models

# This test runs only in the CI workflow, if you run it in local maybe fail 
test: swagger
	go test */*_test.go
