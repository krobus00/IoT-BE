export GO111MODULE ?= on

mod:
	go mod download

tidy:
	go mod tidy
