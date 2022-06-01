export GO111MODULE ?= on

mod:
	go mod download

tidy:
	go mod tidy

mock:
	rm -rf mocks/* && \
	mockery --dir=./api/service/ --case=underscore --all --disable-version-string && \
	mockery --dir=./api/repository/ --case=underscore --all --disable-version-string && \
	mockery --dir=./api/requester/ --case=underscore --all --disable-version-string && \
	mockery --dir=./infrastructure --case=underscore --all --disable-version-string && \
	mockery --dir=./util/ --case=underscore --all --disable-version-string

test:
	go test ./... -coverprofile cover.out && go tool cover -func cover.out

cover:
	go test -cover -coverprofile=cover.out `go list ./...` && go tool cover -html=cover.out