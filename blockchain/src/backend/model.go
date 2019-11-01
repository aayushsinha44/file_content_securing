package backend

type Owner struct {
	PrivateKey   string `json:privatekey`
	OwnerAddress string `json:owneraddress`
	FileAddress  string `json:fileaddress`
	OwnerPower   int64  `json:ownerpower`
}
type Message struct {
	Mess string `json:message`
}
type File struct {
	FileName   string `json:filename`
	PrivateKey string `json:privatekey`
}
type AddressResponse struct {
	Address string `json:address`
}
type AddressList struct {
	Address []string
}
type OwnerList struct {
	ListOfOwner []string
}
type FileAddress struct {
	PrivateKey string `json:privatekey`
	Address    string `json:address`
}
type FileName struct {
	Name string `json:name`
}
type FileText struct {
	PrivateKey string `json:privatekey`
	Text       string `json:text`
}
