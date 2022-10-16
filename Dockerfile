# goバージョン
FROM golang:latest

# COPY go.mod /go.mod
# COPY go.sum /go.sum 
# RUN go mod download

WORKDIR /go/src
# COPY ./api/ /go/src/api

# COPY go/go.mod /go.mod
# COPY go/go.sum /go.sum 
# RUN go mod download

# WORKDIR /go/src
# COPY go/api /go/src

ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64
EXPOSE 1324

# CMD ["go", "run", "api/main.go"]
