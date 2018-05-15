GO=go
GOFMT=gofmt


.PHONY: test bench format build


test:
	${GO} test .


bench:
	${GO} test -bench .


format:
	${GOFMT} -d .


build: sudoku


sudoku: cmd/sudoku.go
	${GO} build cmd/sudoku.go
