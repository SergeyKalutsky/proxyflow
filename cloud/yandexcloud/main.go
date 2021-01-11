package main

import (
	"fmt"

	"github.com/proxyflow"
)

func main() {
	fmt.Println(proxyflow.ReadToken("yandexOauthToken.txt"))
}
