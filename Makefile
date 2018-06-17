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


build: build/sudoku-solve build/sudoku-gen


build/sudoku-%: cmd/sudoku-%.go
	@mkdir -p build
	${GO} build -o $@ $<


clean:
	-rm -vf build/*
