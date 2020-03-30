FROM golang:latest

WORKDIR /go/src/api

COPY . .

RUN go get -d -t  -v ./...

RUN go install -i -v ./...

EXPOSE 8081
CMD ["api"]