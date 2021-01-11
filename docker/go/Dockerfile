FROM golang:latest

ENV GOPATH /go
ADD . /go
WORKDIR /go
RUN go get github.com/gorilla/mux && go get github.com/carlescere/scheduler && go get github.com/go-sql-driver/mysql  && go get github.com/gorilla/handlers && go get github.com/rs/cors
RUN go install github.com/gorilla/mux && go install github.com/carlescere/scheduler &&  go install github.com/go-sql-driver/mysql && go install github.com/gorilla/handlers && go install github.com/rs/cors