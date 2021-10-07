package main

import (
    "os"
    "fmt"
    "time"
    "github.com/cyberark/conjur-api-go/conjurapi"
)

func main() {
    variableIdentifier := "secrets/username"

    config, err := conjurapi.LoadConfig()
    if err != nil {
        panic(err)
    }
	
	authnTokenFile := os.Getenv("CONJUR_AUTHN_TOKEN_FILE")

    conjur, err := conjurapi.NewClientFromTokenFile(config, authnTokenFile)
    if err != nil {
        panic(err)
    }
    for true {
        // Retrieve a secret into []byte.
        secretValue, err := conjur.RetrieveSecret(variableIdentifier)
       if err != nil {
        panic(err)
       }
       fmt.Println("The secret value is: ", string(secretValue))

       // Retrieve a secret into io.ReadCloser, then read into []byte.
       // Alternatively, you can transfer the secret directly into secure memory,
       // vault, keychain, etc.
       secretResponse, err := conjur.RetrieveSecretReader(variableIdentifier)
       if err != nil {
           panic(err)
       }

       secretValue, err = conjurapi.ReadResponseBody(secretResponse)
       if err != nil {   
           panic(err)
       }
       fmt.Println("The secret value is: ", string(secretValue))
       time.Sleep(10 * time.Second)
    }
}