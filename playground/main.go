package main

import (
	cli2 "github.com/pesakit/pesakit/cli"
	"github.com/pesakit/pesalib/vat"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"sort"
)

//func main() {
//	// Use the best keyring implementation for your operating system
//	kr, err := keyring.Open(keyring.Config{
//		AllowedBackends:                nil,
//		ServiceName:                    "pesakit",
//		KeychainName:                   "",
//		KeychainTrustApplication:       false,
//		KeychainSynchronizable:         false,
//		KeychainAccessibleWhenUnlocked: false,
//		KeychainPasswordFunc: func(s string) (string, error) {
//			log.Println("Enter password for keychain:", s)
//			return "", nil
//		},
//		FilePasswordFunc: func(s string) (string, error) {
//			log.Println("Enter password for file:", s)
//			return "", nil
//		},
//		FileDir:                 "",
//		KWalletAppID:            "",
//		KWalletFolder:           "",
//		LibSecretCollectionName: "",
//		PassDir:                 "",
//		PassCmd:                 "",
//		PassPrefix:              "",
//		WinCredPrefix:           "",
//	})
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	err = kr.Set(keyring.Item{
//		Key:                         "username",
//		Data:                        []byte("password"),
//		Label:                       "passwords",
//		Description:                 "this stores passwords",
//		KeychainNotTrustApplication: false,
//		KeychainNotSynchronizable:   false,
//	})
//	if err != nil {
//		return
//	}
//
//	v, err := kr.Get("username")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	log.Printf("llamas was %v", v)
//}

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "lang",
				Aliases: []string{"l"},
				Value:   "english",
				Usage:   "Language for the greeting",
			},
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Load configuration from `FILE`",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "complete",
				Aliases: []string{"c"},

				Usage: "complete a task on the list",
				//UsageText:   "Usage Text Here Is my input",
				Description: "Commands Desc here is My Input",
				Category:    "push category here is my input",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a task to the list",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			cli2.VatCommand(os.Stderr, vat.NewCalculator()),
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
