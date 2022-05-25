FROM golang:alpine as builder

WORKDIR /notifyer-build

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o notifyer cmd/main.go

FROM alpine:latest

RUN apk add ca-certificates

WORKDIR /notifyer

COPY --from=builder /notifyer-build/notifyer .

CMD [ "./notifyer" ]



