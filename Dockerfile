FROM golang:1.16-alpine

WORKDIR /go/src/app

ADD src src

CMD ["go", "run", "src/main.go"]
