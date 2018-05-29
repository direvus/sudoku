GO=go
GOFMT=gofmt


.PHONY: test bench format build


test:
	${GO} test .


bench:
	${GO} test -bench .


format:
	${GOFMT} -d .


cover:
	${GO} test -coverprofile=coverage.txt -covermode=atomic


build: sudoku


sudoku: cmd/sudoku.go
	${GO} build cmd/sudoku.go
