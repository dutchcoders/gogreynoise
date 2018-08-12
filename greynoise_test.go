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
