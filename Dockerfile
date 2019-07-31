FROM golang:1.12
LABEL maintainer="andrej.jovanovic01@gmail.com"

WORKDIR /go/src/app
COPY main.go .

RUN go get -d -v github.com/gorilla/mux
RUN go install github.com/gorilla/mux

EXPOSE 10000/tcp

ENTRYPOINT go run main.go