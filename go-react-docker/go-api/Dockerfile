FROM golang:1.21-alpine

WORKDIR /api

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /go-api

EXPOSE 8080

CMD ["/go-api"]