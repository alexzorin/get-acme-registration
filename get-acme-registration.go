package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"context"

	"golang.org/x/crypto/acme"
)

var (
	acmeDirectory = "https://acme-v01.api.letsencrypt.org/directory"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: get-user-registration [keyfile.pem]")
		return
	}

	keyBuf, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Failed to read key file %s: %v\n", os.Args[1], err)
		return
	}

	decoded, _ := pem.Decode(keyBuf)
	key, err := x509.ParsePKCS1PrivateKey(decoded.Bytes)
	if err != nil {
		fmt.Printf("Failed to parse key file: %v\n", err)
		return
	}

	if dir := os.Getenv("ACME_DIRECTORY"); dir != "" {
		acmeDirectory = dir
	}

	cl := acme.Client{
		Key:          key,
		DirectoryURL: acmeDirectory,
	}

	if _, err = cl.Register(context.Background(), &acme.Account{AgreedTerms: "intentionally_failing"}, func(tos string) bool {
		return false
	}); err != nil {
		switch aErr := err.(type) {
		case *acme.Error:
			if aErr.StatusCode == 409 {
				fmt.Printf("Found existing registration: %s\n", aErr.Header.Get("Boulder-Requester"))
				return
			}
			fmt.Printf("Couldn't find key, probably not registered: (%v)\n", aErr.Detail)
			return
		default:
			fmt.Printf("Failed to create registration: %v (%v)\n", err, reflect.TypeOf(err))
		}
	}

	panic("This shouldn't happen")

}
