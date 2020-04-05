FROM golang:latest

WORKDIR /go/src/used_timer
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN ls
CMD go run main/main.go
