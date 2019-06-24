
#build stage
FROM go-base:latest AS builder

WORKDIR /go/src/app

COPY . .

RUN go get $(go list ./... | grep -v /vendor/)

RUN go build -o /bin/app

#final stage
FROM alpine:latest

RUN apk --no-cache add \
    ca-certificates \
    tzdata 

WORKDIR /bin/

COPY --from=builder /bin/app .

CMD ["/bin/app"]

LABEL Name=go-release Version=1.0.0