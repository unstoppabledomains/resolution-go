package main

/*
#include <android/log.h>
*/

import (
	"fmt"
	"syscall/js"

	"github.com/unstoppabledomains/resolution-go/v3/resolution"
)

//export Addr
func Addr(domainName string, ticker string) string {
	fmt.Println("input domainName", domainName)
	fmt.Println("input ticker", ticker)
	uns, _ := resolution.NewUnsBuilder().SetUdClient("<api_key>").Build()
	ethAddress, _ := uns.Addr(domainName, ticker)
	return ethAddress
}

func NameHash(domainName js.Value, inputs []js.Value) interface{} {
	var domain string = inputs[0].String()
	var apiKey string = inputs[1].String()
	uns, _ := resolution.NewUnsBuilder().SetUdClient(apiKey).Build()

	nameHash, _ := uns.Namehash(domain)

	return nameHash
}

//export Hello
func Hello() {
	fmt.Println("hello world")
}

func main() {
	fmt.Println(Addr("brad.crypto", "ETH"))
	wait := make(chan struct{})
	js.Global().Set("NameHash", js.FuncOf(NameHash))
	<-wait
}
