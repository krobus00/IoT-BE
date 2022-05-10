# builder image
FROM golang:1.18.1-alpine as base

ARG github_id
ARG github_token

WORKDIR /builder
RUN apk add upx
ENV GO111MODULE=on CGO_ENABLED=0
RUN apk add git
RUN git config --global url."https://${github_id}:${github_token}@github.com".insteadOf "https://github.com"
COPY go.mod go.sum /builder/
RUN go mod download
COPY . .
RUN go build \
  -ldflags "-s -w" \
  -o /builder/main /builder/cmd/app/main.go
RUN upx -9 /builder/main

# runner image
FROM alpine:3.8
WORKDIR /app
COPY --from=base /builder/main main
COPY --from=base /builder/.env .env
CMD ["/app/main"]