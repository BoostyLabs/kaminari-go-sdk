# Kaminari Go SDK

A Golang package for the using Kaminari API.

### Installation

    go get -u github.com/BoostyLabs/kaminari-go-sdk

### Basic Usage

Grab your [`API token`](https://app.kaminari.cloud/developers/api-keys).
With that api token and Kaminari's API url you could initialize client.

```go
package main

import (
    "log"

    kaminari "github.com/BoostyLabs/kaminari-go-sdk"
    "github.com/BoostyLabs/kaminari-go-sdk/client"
)

func main() {
    cl, err := client.DefaultClient(&client.Config{
        ApiKey: "[API_KEY]",
        ApiUrl: "[API_URL]",
    })
    if err != nil {
    	log.Println(err)
        return	
    }

    balance, err := cl.GetBalance()
    if err != nil {
        log.Println(err)
        return
    }
	
    log.Println("your balance", balance)
}
```

More examples you could find in `client_test.go` file
