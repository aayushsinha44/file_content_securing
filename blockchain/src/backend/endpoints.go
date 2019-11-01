package backend

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/context"
	"github.com/aayushsinha44/file_content_securing/blockchain/src/blockchain"
	"github.com/aayushsinha44/file_content_securing/blockchain/src/ipfs"
	"github.com/aayushsinha44/file_content_securing/blockchain/src/backend"
)

var blockChainInfo BlockChainInfo = &blockchain.BlockChainInfo{
 	ClientAddress: "192.168.114.27",
	ClientPort:   7545,
	IsHttps:       false,
} 

func makeConnection(privateKey string) (*blockchain.BlockChain, error) {
	userInfo, err := blockchain.GetUserInfo(privateKey)
	if err != nil {
		return nil, err
	}
	bc, err := userInfo.MakeConnection(nil, blockChainInfo)
	return bc, err
}
//Creating a new file and returning its address
func CreateNewFile(w http.ResponseWriter,r *http.Request){
	var file File
	err:=json.NewDecoder(r.Body).Decode(&file)
	if err!=nil{
		log.Fatalln(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Message{"Bad structure of request"})
		return
	}
	bc, err := makeConnection(file.PrivateKey)
	if err!=nil{
		log.Fatalln(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Message{"Bad structure of request"})
		return
	}
	address, err := bc.StoreFile(file.FileName)
	if err!=nil{
		log.Fatalln(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Message{err.Error()})
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncode(w).Encode(AddressResponse{address})
	return
} 
//Adding an owner to a file
func AddOwner(w http.ResponseWriter,r *http.Request){
	var user Owner
	err := json.NewDecoder(r.Body).Decode(&user)
	if err !=nil{
		log.Fatalln(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Message{"Bad structure of request"})
		return
	}
	// TODO :-Validation of request body
	bc, err := makeConnection(user.PrivateKey)
	if err!= nil {
		log.Fatalln(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Message{"Error in creating connection"})
		return
	}
	fa := &blockchain.FileAddress{Address: user.FileAddress}

	// Adding new Owner
	status, err := bc.AddOwner(fa ,user.OwnerAddress, user.OwnerPower)
	if err != nil {
		log.Fatalln(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Message{"Error in  Adding Owner"})
		return
	}
	if status == "success" {
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(Message{"Owner added Successfully"})
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Message{"Owner Nhi add hua"})
	}
}


//Getting the file name using file address

func GetFileName(w http.ResponseWriter,r *http.Request){
	var FileAdd FileAddress
	err=json.NewDecoder(r.Body).Decode(&FileAdd)
	if err!=nil{
		log.Fatalln(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Message{err.Error()})
		return
	}
	bc, err := makeConnection(FileAdd.PrivateKey)
	if err!= nil {
		log.Fatalln(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Message{"Error in creating connection"})
		return
	}
	filename,err:=bc.GetFileName(FileAdd.Address)
	if err!=nil{
		log.Fatalln(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Message{err.Error()})
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(FileName{filename})
	return
}

//Getting the list of owners
func GetOwnerList(w http.ResponseWriter,r *http.Request){
	var FileAdd FileAddress
	// err=json.NewDecoder(r.Body).Decode(&FileAdd)
	// if err!=nil{
	// 	log.Fatalln(err)
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	json.NewEncoder(w).Encode(Message{err.Error()})
	// 	return
	// }
	faddress string
	faddress:=r.Header.Get("FileAddress")
	privateki string
	privateki=r.Header.Get("PrivateKey")
	FileAdd.PrivateKey=privateki
	FileAdd.Address=faddress
	bc, err := makeConnection(FileAdd.PrivateKey)
	if err!= nil {
		log.Fatalln(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Message{"Error in creating connection"})
		return
	}
	var addressList []string
	fa:=AddressResponse{faddress}
	addressList,err=bc.GetOwnerList(fa)
	if err!= nil {
		log.Fatalln(err)
		json.NewEncoder(w).Encode(Message{err.Error()})
		return
	}
	json.NewEncoder(w).Encode(AddressList{addressList})
	return
	// var Addressr AddressResponse
	// Addressr.Address=fadress
	// var addressList AddressList
	// addressList,err=bc.GetOwnerList(Addressr)
	// if err!= nil {
	// 	log.Fatalln(err)
	// 	json.NewEncoder(w).Encode(Message{err.Error()})
	// 	return
	// }
	// json.NewEncoder(w).Encode(addressList)
	return
}

//Uploading file to ipfs

func UploadFile(w http.ResponseWriter,r *http.Request){
	var ftext FileText
	err=json.NewDecoder(r.Body).Decode(&ftext)
	if err!=nil{
		log.Fatalln(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Message{err.Error()})
		return
	}
	ipfsInfo := &ipfs.IPFSInfo{IPFSUrl: "localhost:5001"}
	ipfsHash, err := ipfs.UploadFile(ftext.Text, ipfsInfo)
	if err!=nil{
		log.Fatalln(err)
		json.NewEncoder(w).Encode(Message{err.Error()})
		return
	}
	log.Println("File successfully uploaded to IPFS")
	status, err = bc.UpdateKey(fa, "new-encrpytion-key")
	if err!=nil{
		log.Fatalln(err)
		json.NewEncoder(w).Encode(Message{err.Error()})
		return
	}
	if status=="success"{
		log.Println("Ipfs hash successfully updated on the blockchain")
		json.NewEncoder(w).Encode(Message{"Ipfs hash successfully updated on the blockchain"})
		return
	} else{
		log.Println("Ipfs hash unsuccessfully updated on the blockchain")
		json.NewEncoder(w).Encode(Message{"Ipfs hash unsuccessfully updated on the blockchain"})
		return
	}
}