FROM golang:1.19 as builder

WORKDIR /go/src/phatchips.net/transcoder

COPY go.mod go.sum ./
RUN go mod download

COPY src ./
RUN CGO_ENABLED=0 go build -o app

FROM alpine:3.17

RUN apk --no-cache add ca-certificates ffmpeg

WORKDIR /root/

RUN mkdir -p /opt/output

COPY --from=builder /go/src/phatchips.net/transcoder/app .

CMD ["./app"]