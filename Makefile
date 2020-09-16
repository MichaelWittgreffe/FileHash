APPLICATION_NAME = filehash
APPLICATION_PUBLISHER = michaelwittgreffe

all:
	@$(MAKE) clean create-dir install test build success || $(MAKE) failure

build:
	GOOS=linux go build -o bin/${APPLICATION_NAME} -v cmd/${APPLICATION_NAME}.go

ide-build:
	@$(MAKE) build success || $(MAKE) failure

clean:
	go clean
	if [ -f ./bin/${APPLICATION_NAME} ]; then rm ./bin/${APPLICATION_NAME}; fi;
	@$(MAKE) clean-test-data

clean-test-data:
	if [ -f ./coverage.html ]; then rm ./coverage.html; fi;
	if [ -f ./coverage.out ]; then rm ./coverage.out; fi;

test:
	@$(MAKE) clean-test-data
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

install:
	go mod tidy
	go mod download

success:
	printf "\n\e[1;32mBuild Successful\e[0m\n"

failure:
	printf "\n\e[1;31mBuild Failure\e[0m\n"
	exit 1

create-dir:
	if ! [ -d ./bin ]; then mkdir bin; fi;