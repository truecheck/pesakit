package main

import (
	"log"

	keyring "github.com/99designs/keyring"
)

func main() {
	// Use the best keyring implementation for your operating system
	kr, err := keyring.Open(keyring.Config{
		AllowedBackends:                nil,
		ServiceName:                    "pesakit",
		KeychainName:                   "",
		KeychainTrustApplication:       false,
		KeychainSynchronizable:         false,
		KeychainAccessibleWhenUnlocked: false,
		KeychainPasswordFunc: func(s string) (string, error) {
			log.Println("Enter password for keychain:", s)
			return "", nil
		},
		FilePasswordFunc: func(s string) (string, error) {
			log.Println("Enter password for file:", s)
			return "", nil
		},
		FileDir:                 "",
		KWalletAppID:            "",
		KWalletFolder:           "",
		LibSecretCollectionName: "",
		PassDir:                 "",
		PassCmd:                 "",
		PassPrefix:              "",
		WinCredPrefix:           "",
	})

	if err != nil {
		log.Fatal(err)
	}

	err = kr.Set(keyring.Item{
		Key:                         "username",
		Data:                        []byte("password"),
		Label:                       "passwords",
		Description:                 "this stores passwords",
		KeychainNotTrustApplication: false,
		KeychainNotSynchronizable:   false,
	})
	if err != nil {
		return
	}

	v, err := kr.Get("username")
	if err != nil {
		log.Fatal(err)
	}
	
	log.Printf("llamas was %v", v)
}
