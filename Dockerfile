FROM golang:1.21-alpine as builder

WORKDIR /app
COPY . .

RUN go build -o dalle main.go

FROM alpine

COPY --from=builder /app/dalle /dalle

CMD ["/dalle"]