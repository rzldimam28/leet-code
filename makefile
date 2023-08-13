test-all:
	@go test -v ./...

test:
	@go test ${DIR} -v -run ${NAME} # not recommended:(