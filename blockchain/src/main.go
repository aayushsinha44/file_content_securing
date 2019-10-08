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

	// Creating new file
	address, err := bc.StoreFile("TestFile")
	must(err)
	fmt.Println("Address of File:", address)

	// Getting file name
	fa := &blockchain.FileAddress{Address: address}
	fileName, err := bc.GetFileName(fa)
	must(err)
	fmt.Println("File Name:", fileName)

	// Updating file Name
	status, err := bc.UpdateFileName(fa, "TestFileChanged")
	must(err)
	fmt.Println("File Name Changing: ", status)

	fileName, err = bc.GetFileName(fa)
	must(err)
	fmt.Println("New File Name:", fileName)

	// Getting owner list
	addressList, err := bc.GetOwnerList(fa)
	must(err)
	fmt.Println(addressList)
	// for index, element := range addressList {
	// 	fmt.Println(index, element)
	// }

}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
