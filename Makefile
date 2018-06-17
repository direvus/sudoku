GO=go
GOFMT=gofmt


.PHONY: test bench format install


test:
	${GO} test .


bench:
	${GO} test -bench .


format:
	${GOFMT} -d .


cover:
	${GO} test -coverprofile=coverage.txt -covermode=atomic


install:
	${GO} install ./...
