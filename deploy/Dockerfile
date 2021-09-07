# base image
FROM golang:1.16.4-alpine3.13 as build

RUN apk add --no-cache --update git build-base openssh-client

WORKDIR /go/src/api

COPY . .

RUN ls -l

RUN git config --global http.sslVerify true &&\
    go get -v github.com/swaggo/swag/cmd/swag &&\
    go mod download &&\
    $GOPATH/bin/swag init -g cmd/main.go

RUN CGO_ENABLED=0 GOOS=linux go build cmd/main.go

RUN ls -l

FROM alpine:3.12
WORKDIR /app
COPY --from=build /go/src/api/main /app/medilane-api
RUN chmod +x /app/medilane-api
ENTRYPOINT ["/app/medilane-api"]

