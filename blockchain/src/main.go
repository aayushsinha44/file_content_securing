package main

import (
	"fmt"
	"log"

	"github.com/aayushsinha44/file_content_securing/blockchain/src/blockchain"
	"github.com/aayushsinha44/file_content_securing/blockchain/src/ipfs"
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

	// Adding new Owner
	status, err = bc.AddOwner(fa, "0x761367D4b30eAcFaA250A03CdD2c78e949f6Bb1a", 3)
	must(err)
	fmt.Println("New Owner Added: ", status)

	// Change Power of owner
	status, err = bc.ChangeOwnerPower(fa, "0x761367D4b30eAcFaA250A03CdD2c78e949f6Bb1a", 1)
	must(err)
	fmt.Println("Owner Power Changed: ", status)

	// Getting owner list
	addressList, err := bc.GetOwnerList(fa)
	must(err)
	fmt.Println("Owner List")
	fmt.Println(addressList)

	// Owner Power
	for _, element := range addressList {
		status, err := bc.GetOwnerPower(fa, element)
		must(err)
		fmt.Println("Owner: ", element)
		fmt.Println("Owner Power: ", status)
	}

	// Upload File to IPFS
	text := "Hello World"
	ipfsHash, err := ipfs.UploadFile(text)
	must(err)
	fmt.Println("File succesfully updated to ipfs")

	// Update IPFS Hash
	status, err = bc.UpdateIPFSHash(fa, ipfsHash)
	must(err)
	fmt.Println("IPFS Hash updated")

	// Get IPFS Hash
	hash, err := bc.GetIPFSHash(fa)
	must(err)
	fmt.Println("IPFS Hash: ", hash)

	// Update key
	status, err = bc.UpdateKey(fa, "new-encrpytion-key")
	must(err)
	fmt.Println("key updated")

	// Get IPFS Hash
	key, err := bc.GetKey(fa)
	must(err)
	fmt.Println("Key: ", key)

}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
