package ipfs

import (
	"strings"

	shell "github.com/ipfs/go-ipfs-api"
)

func UploadFile(text string) (string, error) {
	// Where your local node is running on localhost:5001
	sh := shell.NewShell("localhost:5001")
	cid, err := sh.Add(strings.NewReader(text))
	if err != nil {
		return "", err
	}
	return cid, nil
}
