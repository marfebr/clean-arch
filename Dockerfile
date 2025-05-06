FROM golang:latest as base


COPY .  /app
WORKDIR /app
RUN   curl -LO https://github.com/golang-migrate/migrate/releases/download/v4.18.3/migrate.linux-amd64.tar.gz; \
      tar zxvf migrate.linux-amd64.tar.gz

RUN go mod tidy

FROM base as builder

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o server ./cmd/ordersystem/main.go ./cmd/ordersystem/wire_gen.go

EXPOSE 8000


FROM scratch as prod

WORKDIR /app
COPY --from=builder /app/server .
COPY --from=builder /app/migrations migrations
COPY --from=builder /app/migrate .
COPY --from=base /app/cmd/ordersystem/.env /app/.env

CMD [ "./server" ]




