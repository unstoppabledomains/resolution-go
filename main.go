package main

/*
#include <android/log.h>
*/

import "C"
import (
	"fmt"
	"os"

	"github.com/unstoppabledomains/resolution-go/v3/resolution"
)

func getAuthKey() string {
	if os.Getenv("UD_AUTH_KEY") != "" {
		return os.Getenv("UD_AUTH_KEY")
	}

	panic("UD_AUTH_KEY is not set!")
}

//export Addr
func Addr(domainName string, ticker string) *C.char {
	fmt.Println("input domainName", domainName)
	fmt.Println("input ticker", ticker)
	uns, _ := resolution.NewUnsBuilder().SetUdClient(getAuthKey()).Build()
	ethAddress, _ := uns.Addr(domainName, ticker)
	return C.CString(ethAddress)
}

//export Hello
func Hello() {
	fmt.Println("hello world")
}

func main() {
	// Addr("brad.crypto", "ETH")
}
