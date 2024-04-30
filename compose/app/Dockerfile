FROM golang:1.22.2-alpine as builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash gcc gettext musl-dev

COPY ["app/go.mod", "app/go.sum", "app/.env", "./"]
RUN go mod download

COPY ../../app ./
RUN go build -o ./bin/app cmd/app/main.go

FROM alpine as runner

COPY --from=builder /usr/local/src/bin/app /

RUN chmod -R 777 /app

CMD ["/app"]