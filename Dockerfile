FROM golang:1.14.4-stretch

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main .

WORKDIR /dist

RUN cp /go/src/app/main .

RUN rm -rf /go/src/app

COPY config config

CMD ["/dist/main"]