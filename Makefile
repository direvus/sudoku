GO=go
GOFMT=gofmt


.PHONY: test bench format


test:
	${GO} test .


bench:
	${GO} test -bench .


format:
	${GOFMT} -d .


sudoku: cmd/sudoku.go
	${GO} build cmd/sudoku.go
