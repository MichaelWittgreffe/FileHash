
all:
	@$(MAKE) clean create-dir install test build success || $(MAKE) failure

ide-build:
	@$(MAKE) clean build success || $(MAKE) failure

build:
	go build -o bin/filehash -v

clean:
	go clean
	if [ -d ./bin/filehash ]; then rm ./bin/filehash; fi;

test:
	go test ./... -coverprofile=coverage.out

coverage:
	go tool cover -html=coverage.out

install:
	go get -u github.com/spf13/cobra/cobra

success:
	printf "\n\e[1;32mBuild Successful\e[0m\n"

failure:
	printf "\n\e[1;31mBuild Failure\e[0m\n" 
	exit 1

create-dir:
	if ! [ -d ./bin ]; then mkdir bin; fi;
