package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	file "github.com/aayushsinha44/file_content_securing/blockchain/src/file"
)

type BlockChain struct {
	Client *ethclient.Client
	Auth   *bind.TransactOpts
}

type BlockChainInfo struct {
	ClientAddress string
	ClientPort    int
	IsHttps       bool
}

type UserInfo struct {
	PrivateKey  *ecdsa.PrivateKey
	FromAddress common.Address
}

func GetUserInfo(privateKey string) (*UserInfo, error) {
	pk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}
	publicKey := pk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	userInfo := &UserInfo{
		PrivateKey: pk, FromAddress: fromAddress}

	return userInfo, nil
}

func (ui *UserInfo) MakeConnection(ctx context.Context, info *BlockChainInfo) (*BlockChain, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	if info == nil {
		return nil, fmt.Errorf("Blockchain information is required")
	}
	// if info.ClientAddresss == nil {
	// 	return nil, fmt.Errorf("Blockchain client address is required")
	// }
	// if info.IsHttps == nil {
	// 	// info.IsHttps = false
	// }
	// if info.ClientPort == nil {
	// 	// info.ClientPort = 80
	// }
	var address string
	if info.IsHttps {
		address = fmt.Sprintf("%s%s:%d", "https://", info.ClientAddress, info.ClientPort)
	} else {
		address = fmt.Sprintf("%s%s:%d", "http://", info.ClientAddress, info.ClientPort)
	}

	client, err := ethclient.Dial(address)
	if err != nil {
		return nil, err
	}

	nonce, err := client.PendingNonceAt(ctx, ui.FromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	gasLimit := uint64(300000 * 10)
	auth := bind.NewKeyedTransactor(ui.PrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	blockchain := &BlockChain{Client: client, Auth: auth}

	return blockchain, nil

}

func (b *BlockChain) StoreFile(fileName string) (string, error) {

	address, _, _, err := file.DeployFile(b.Auth, b.Client, fileName)
	if err != nil {
		return "", err
	}
	return address.Hex(), nil
}
