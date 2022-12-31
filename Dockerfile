FROM golang:1.19-alpine as builder

ADD . /src
WORKDIR /src


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o movie-catalog .

FROM alpine
COPY --from=builder /src/movie-catalog /usr/local/bin/movie-catalog
WORKDIR /usr/local/bin
EXPOSE 8081

ENTRYPOINT ["./movie-catalog"]