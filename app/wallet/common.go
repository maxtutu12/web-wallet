package wallet

import (
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func Decrypt(filename, password string) (*keystore.Key, error) {
	storeData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("read %s file err:%v\n", filename, err)
		return nil, err
	}

	key, err := keystore.DecryptKey(storeData, password)
	if err != nil {
		log.Printf("keystore decrypt key err:%v\n", err)
		return nil, err
	}

	return key, nil
}
