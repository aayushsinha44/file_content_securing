package main

import (
	"fmt"
	"log"

	"github.com/aayushsinha44/file_content_securing/blockchain/src/blockchain"
)

func main() {
	privateKey := "5afecc0e14aed67e2999ecfddc34b8054e8ccb48268f961ee04485c6e4e303ae"

	userInfo, err := blockchain.GetUserInfo(privateKey)
	must(err)
	blockChainInfo := &blockchain.BlockChainInfo{
		ClientAddress: "localhost",
		ClientPort:    7545,
		IsHttps:       false,
	}

	bc, err := userInfo.MakeConnection(nil, blockChainInfo)
	must(err)
	address, err := bc.StoreFile("TestFile")
	must(err)
	fmt.Println(address)
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
