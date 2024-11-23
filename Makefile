.PHONY: chapter2section4-1

install:
	go install github.com/rakyll/gotest@latest

chapter2section4-1:
	@gotest -v -short Chapter-2-Basics\Testing\section4_1_test.go
