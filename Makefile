.PHONY: install chapter2section4-1 chapter2section4-2

install:
	go install github.com/rakyll/gotest@latest

chapter2section4-1:
	@gotest -v -short Chapter-2-Basics\Testing\section4_1_test.go

chapter2section4-2:
	@gotest -v -short Chapter-2-Basics\Testing\section4_2_test.go

chapter2section4-3-1:
	@gotest -v -short Chapter-2-Basics\Testing\section4_3_1_test.go

chapter2section4-3-2:
	@gotest -v -short Chapter-2-Basics\Testing\section4_3_2_test.go