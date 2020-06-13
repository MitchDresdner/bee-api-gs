# Compile stage
#   see: https://github.com/dlsniper/dockerdev
#        https://blog.jetbrains.com/go/2020/05/04/go-development-with-docker-containers/
#
FROM golang:1.13.8 AS build-env

ADD . /dockerdev
WORKDIR /dockerdev

RUN go build -o /server

# Final stage
FROM debian:buster

EXPOSE 8080

WORKDIR /
COPY --from=build-env /server /

CMD ["/server"]