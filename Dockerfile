FROM golang:1.16

WORKDIR /go/src/github.com/gazure/oauth
RUN go get -d -v golang.org/x/net/html

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/gazure/oauth/app .
COPY keys ./keys
COPY templates ./templates
CMD ["./app"]
