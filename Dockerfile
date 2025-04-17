FROM golang:latest as builder

WORKDIR /mimicdb

COPY . .

RUN go mod download

RUN go build -o ./bin/mimicdb ./core/cmd

FROM debian:12

COPY --from=builder /mimicdb/bin/mimicdb /

ENTRYPOINT ["/mimicdb"]
