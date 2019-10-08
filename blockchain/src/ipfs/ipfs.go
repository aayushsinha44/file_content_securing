package ipfs

import (
	"fmt"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"
)

type IPFSInfo struct {
	IPFSUrl string
}

func UploadFile(text string, ipfsinfo *IPFSInfo) (string, error) {
	if ipfsinfo == nil {
		return "", fmt.Errorf("IPFS url required")
	}
	sh := shell.NewShell(ipfsinfo.IPFSUrl)
	cid, err := sh.Add(strings.NewReader(text))
	if err != nil {
		return "", err
	}
	return cid, nil
}
