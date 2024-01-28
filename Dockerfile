FROM golang:latest as builder

WORKDIR /build
COPY go.mod . 
COPY go.sum . 
COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait-for-postgres.sh

CMD ["migrate -path ./schema -database 'postgres://postgres:test@localhost:5436/postgres?sslmode=disable' up"]

RUN go build -o user-app ./cmd/main.go
CMD ["./user-app"]