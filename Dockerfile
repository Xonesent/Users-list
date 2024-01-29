FROM golang:latest as builder

WORKDIR /build
COPY go.mod . 
COPY go.sum . 
COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait-for-postgres.sh
RUN go build -o user-app ./cmd/main.go

CMD ["./user-app"]