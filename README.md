# gogreynoise [![](https://godoc.org/github.com/dutchcoders/gogreynoise?status.svg)](http://godoc.org/github.com/dutchcoders/gogreynoise) [![Go Report Card](https://goreportcard.com/badge/dutchcoders/gogreynoise)](https://goreportcard.com/report/dutchcoders/gogreynoise) 

Golang client for Greynoise.io. The client currently only implements /context, but it is easy to extend other features.

## Usage

```
package greynoise

import (
        "fmt"
        "net"

        greynoise "github.com/dutchcoders/gogreynoise"
)

func ExampleExamples_output() {
        client, err := greynoise.New(
                greynoise.WithKey("{key}"),
        )
        if err != nil {
                panic(err.Error)
        }

        ip := net.ParseIP("8.8.8.8")
        result, err := client.Context(ip)
        if err != nil {
                panic(err.Error)
        }

        fmt.Printf("Actor: %s\n", result.Actor)
}
```

## Contributors

* [Remco Verhoef](https://twitter.com/remco_verhoef)

## Copyright and license

Code released under [Apache License 2.0](LICENSE).

