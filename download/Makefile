# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)


test: fmt
	@echo "\033[92mTest Reading Stories ...\033[0m"
	@go test -v

dl:
	@echo "\033[92mRun forever ...\033[0m"
	@go run example/main.go

fmt:
	@echo "\033[92mGo fmt source code...\033[0m"
	@go fmt *.go

install:
	@echo "\033[92mInstalling github.com/siongui/goigstorylink ...\033[0m"
	go get -u github.com/siongui/goigstorylink
	@echo "\033[92mInstalling github.com/fatih/color ...\033[0m"
	go get -u github.com/fatih/color
