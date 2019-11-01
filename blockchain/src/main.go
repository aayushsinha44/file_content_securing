package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/aayushsinha44/file_content_securing/blockchain/src/blockchain"
	"github.com/aayushsinha44/file_content_securing/blockchain/src/ipfs"
	"github.com/aayushsinha44/file_content_securing/blockchain/src/backend"
)
	

func main1() {

	privateKey := "4f06fea6b6248b94a6cf64d3cbcb3e6823ed1058dbecebf3915d91ba79751894"

	userInfo, err := blockchain.GetUserInfo(privateKey)
	must(err)
	blockChainInfo := &blockchain.BlockChainInfo{
		ClientAddress: "192.168.114.27",
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
	ipfsInfo := &ipfs.IPFSInfo{IPFSUrl: "localhost:5001"}
	ipfsHash, err := ipfs.UploadFile(text, ipfsInfo)
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

	// Unprivilaged access
	hackerPrivateKey := "729eb80a38ca9b5366402352b0956f28d74463c0b1018cf00330a2c1d4834fcd"
	userInfo, err = blockchain.GetUserInfo(hackerPrivateKey)
	must(err)

	bc, err = userInfo.MakeConnection(nil, blockChainInfo)
	must(err)

	key, err = bc.GetKey(fa)
	must(err)
	fmt.Println("Key: ", key)
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main(){
	router := backend.NewRouter() // create routes
	// These two lines are important if you're designing a front-end to utilise this API methods
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	// Launch server with CORS validations
	log.Fatal(http.ListenAndServe(":"+"8081", handlers.CORS(allowedOrigins, allowedMethods)(router)))

}