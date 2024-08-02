FROM golang:1.9-alpine

WORKDIR /go/src/github.com/aren55555/grpc-example/garage_server
COPY . .

RUN apk update
RUN apk add git

RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."

CMD ["go-wrapper", "run"] # ["app"]

EXPOSE 50051
