package udclient

import (
	"context"
	"net/http"
	"runtime"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type UdClient struct {
	L1ContractBackend *ethclient.Client
	L2ContractBackend *ethclient.Client
}

var libVersion = "resolution-go/v3.0.0"
var proxyBaseUrl = "https://resolve.unstoppabledomains.com"
var l1RpcProxyPath = "/chains/eth/rpc"
var l2RpcProxyPath = "/chains/matic/rpc"

// Dial connects a client to the a proxy service with a authentication key
func Dial(apiKey string) (*UdClient, error) {
	l1ProxyUrl := proxyBaseUrl + l1RpcProxyPath
	l2ProxyUrl := proxyBaseUrl + l2RpcProxyPath

	tokenHeader := rpc.WithHeader("Authorization", "Bearer "+apiKey)

	libClientHeaderString := libVersion + "/" + runtime.Version()
	agentHeader := rpc.WithHeader("X-Lib-Client", libClientHeaderString)

	httpClient := rpc.WithHTTPClient(&http.Client{
		Timeout: 3 * time.Second,
	})

	ctx := context.Background()

	l1RpcClient, err := rpc.DialOptions(ctx, l1ProxyUrl, httpClient, tokenHeader, agentHeader)

	if err != nil {
		return nil, err
	}

	l2RpcClient, err := rpc.DialOptions(ctx, l2ProxyUrl, httpClient, tokenHeader, agentHeader)

	if err != nil {
		return nil, err
	}

	l1ContractBackend := ethclient.NewClient(l1RpcClient)
	l2ContractBackend := ethclient.NewClient(l2RpcClient)

	return &UdClient{
		L1ContractBackend: l1ContractBackend,
		L2ContractBackend: l2ContractBackend,
	}, nil
}
