FROM golang:1.18-alpine as builder


WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy application data into image
COPY . .

RUN CGO_ENABLED=0 go build -o ./server .

# second build (multi-stage build), build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /go/src/app/server .

CMD ["./server"]