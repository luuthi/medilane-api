# base image
FROM golang:1.16.4-alpine3.13 as build

RUN apk add --no-cache --update git build-base openssh-client

WORKDIR /go/src/api

#ARG WEB_PRIVATE_KEY
#ARG GIT_DOMAIN
#
#RUN mkdir ~/.ssh && \
#    echo "$WEB_PRIVATE_KEY" | tr -d '\r' > ~/.ssh/id_rsa && \
#    chmod 600 ~/.ssh/id_rsa && \
#    ssh-keyscan -H $GIT_DOMAIN >> ~/.ssh/known_hosts

COPY . .

RUN ls -l

RUN git config --global http.sslVerify true &&\
    go get -v github.com/swaggo/swag/cmd/swag &&\
    $GOPATH/bin/swag init -g cmd/main.go &&\
    go mod download

RUN CGO_ENABLED=0 GOOS=linux go build cmd/main.go

RUN ls -l

FROM alpine:3.12
WORKDIR /app
COPY --from=build /go/src/api/main /app/account-api
RUN chmod +x /app/account-api
ENTRYPOINT ["/app/account-api"]

