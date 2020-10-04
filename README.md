# File Content Securing

Uses Blockhain technology & IPFS protocol for securing the contents of files.


## Tech Stack
1. Ethereum
2. Golang
3. Solidity
4. Reactjs

## Execution

First start ipfs with this command
```
docker run --network=host ipfs/go-ipfs
```

if you change some code execute this command first or if you are running the code for first time
```
docker build -f DockerfileExec -t <dockerhub:username>/fcsexec .
```

for running code use this command. This will execute blockchain/src/main.go
```
docker run --network=host -p 8081:8081 -it <dockerhub:username>/fcsexec
```
