FROM golang:1.20.7-alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN go mod verify
RUN go build -o /go/bin/server ./cmd

FROM alpine
COPY --from=builder /go/bin/server /app/server
COPY --from=builder /go/src/app/.env /app

WORKDIR /app
EXPOSE 5050
CMD ["./server"]