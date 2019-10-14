# Build Geth in a stock Go builder container
FROM golang:1.13-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

RUN git clone https://github.com/ethereum/go-ethereum.git

RUN cd /go/go-ethereum && make all

# Pull Geth into a second stage deploy alpine container
FROM ubuntu:18.04

COPY --from=builder /go/go-ethereum/build/bin/geth /usr/local/bin/
COPY --from=builder /go/go-ethereum/build/bin/abigen /usr/local/bin/

EXPOSE 8545 8546 8547 30303 30303/udp
RUN geth --rpc &

ENV INITRD No
ENV LANG en_US.UTF-8
ENV GOVERSION 1.13
ENV GOROOT /opt/go
ENV GOPATH /root/.go

RUN apt-get update
RUN apt-get install -y software-properties-common
RUN add-apt-repository ppa:ethereum/ethereum
RUN apt-get -y update
RUN apt-get install -y solc
RUN apt-get install -y wget make gcc musl-dev git

RUN cd /opt && wget https://storage.googleapis.com/golang/go${GOVERSION}.linux-amd64.tar.gz && \
    tar zxf go${GOVERSION}.linux-amd64.tar.gz && rm go${GOVERSION}.linux-amd64.tar.gz && \
    ln -s /opt/go/bin/go /usr/bin/ && \
    mkdir $GOPATH

COPY --from=builder /go/go-ethereum/ /opt/go/src/github.com/ethereum/go-ethereum/
RUN go get github.com/ipfs/go-ipfs-api 
RUN go get github.com/allegro/bigcache
RUN apt clean
COPY ipfs ipfs
RUN chmod +x ipfs
RUN ./ipfs init
RUN ./ipfs daemon &