FROM golang:1.13-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

RUN git clone https://github.com/ethereum/go-ethereum.git

RUN cd /go/go-ethereum/ && make geth


# Pull Geth into a second stage deploy alpine container
FROM ubuntu:18.04

RUN apt-get update
RUN apt-get install -y ca-certificates
COPY --from=builder /go/go-ethereum/build/bin/geth /usr/local/bin/
# COPY --from=builder /go /go
# RUN apt-get update
# RUN apt-get install make

EXPOSE 8545 8546 8547 30303 30303/udp
RUN geth --rpc &