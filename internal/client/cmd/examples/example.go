// Generated by LIBLAB | https://liblab.com

package main

import (
	"fmt"

	"github.com/liblab-sdk/pkg/client"
	"github.com/liblab-sdk/pkg/clientconfig"
)

func main() {
	config := clientconfig.NewConfig()

	client := client.NewClient(config)

	fmt.Printf("%#v", client)
}
