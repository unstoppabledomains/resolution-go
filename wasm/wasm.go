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

//export Addr2
func Addr2(this js.Value, inputs []js.Value) interface{} {
	var domain string = inputs[0].String()
	var ticker string = inputs[1].String()
	var apiKey string = inputs[2].String()
	fmt.Println("input domainName", domain, len(domain))
	fmt.Println("input ticker", ticker, len(ticker))

	uns, _ := resolution.NewUnsBuilder().SetUdClient(apiKey).Build()
	ethAddress, _ := uns.Addr(domain, ticker)
	fmt.Println("ethAddress", ethAddress)
	return ethAddress
}

//export Hello
func Hello() {
	fmt.Println("hello world")
}

func main() {
	fmt.Println(Addr("brad.crypto", "ETH"))
}
