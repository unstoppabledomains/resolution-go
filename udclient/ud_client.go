package udclient

import (
	"context"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type UdClient struct {
	L1ContractBackend *ethclient.Client
	L2ContractBackend *ethclient.Client
}

var l1RpcProxyPath = "/rpcproxy/l1"
var l2RpcProxyPath = "/rpcproxy/l2"

func Dial(apiKey string, proxyBaseUrl string) (*UdClient, error) {
	l1ProxyUrl := proxyBaseUrl + l1RpcProxyPath
	l2ProxyUrl := proxyBaseUrl + l2RpcProxyPath

	tokenHeader := rpc.WithHeader("authorization", "Bearer "+apiKey)
	httpClient := rpc.WithHTTPClient(&http.Client{
		Timeout: 3 * time.Second,
	})

	ctx := context.Background()

	l1RpcClient, err := rpc.DialOptions(ctx, l1ProxyUrl, httpClient, tokenHeader)

	if err != nil {
		return nil, err
	}

	l2RpcClient, err := rpc.DialOptions(ctx, l2ProxyUrl, httpClient, tokenHeader)

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
