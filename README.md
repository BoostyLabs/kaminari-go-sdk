# Kaminari Go SDK

A Golang package for the using Kaminari API.

### Installation

    go get -u github.com/BoostyLabs/kaminari-go-sdk

### Basic Usage

Grab your [`API token` and `Secret Key`](https://app.kaminari.cloud/developers/api-keys).
With that creds and Kaminari's API url you could initialize client.

```go
package main

import (
    "log"

    kaminari "github.com/BoostyLabs/kaminari-go-sdk"
    "github.com/BoostyLabs/kaminari-go-sdk/client"
)

func main() {
    cl, err := client.DefaultClient(&client.Config{
        ApiKey:     "[API_KEY]",
        SecretKey:  "[SECRET_KEY]",
        ApiUrl:     "[API_URL]",
    })
    if err != nil {
    	log.Println(err)
        return	
    }

    balance, err := cl.GetBalance("1")
    if err != nil {
        log.Println(err)
        return
    }

    log.Println("your balance", balance)
}
```

More examples you could find in `client_test.go` file.
