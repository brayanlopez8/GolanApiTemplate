FROM golang:1.8-alpine
RUN apk add --update go git
RUN go get github.com/lib/pq/...
ADD . /go/src/main
RUN go install hello-app
ENV USER=username \
    PASSWORD=password \
    DB=dbname \
    HOST=hostname \
    PORT=5432

FROM alpine:latest
COPY --from=0 /go/bin/hello-app/ .
COPY --from=0 /go/src/hello-app/templates ./templates
ENV PORT 4040
CMD ["./hello-app"]