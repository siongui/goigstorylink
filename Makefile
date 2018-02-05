# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)


test: fmt
	@echo "\033[92mTest Reading Stories ...\033[0m"
	@go test -v getlinks.go getlinks_test.go config.go allstories.go userstories.go debug.go

userinfo: fmt
	@echo "\033[92mTest userinfo.go ...\033[0m"
	@go test -v userinfo.go userinfo_test.go

fmt:
	@echo "\033[92mGo fmt source code...\033[0m"
	@go fmt *.go
