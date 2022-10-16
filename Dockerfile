# goバージョン
FROM golang:latest

WORKDIR /go/src

ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64
EXPOSE 1324

# CMD ["go", "run", "api/main.go"]
