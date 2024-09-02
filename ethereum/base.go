package ethereum

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/hashicorp/go-retryablehttp"
	"net/http"
	"time"
)

type ethereumClient interface {
	ethereum.ChainReader
	ethereum.ChainStateReader
	ethereum.ContractCaller
	ethereum.GasEstimator
	ethereum.GasPricer
	ethereum.LogFilterer
	ethereum.TransactionReader
	ethereum.TransactionSender

	bind.DeployBackend
}

// newRetryClient creates a new retry Client with default settings.
func newRetryClient() *http.Client {
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 10
	retryClient.CheckRetry = CustomRetryPolicy

	return retryClient.StandardClient()
}

// Dial connects a client to the given URL.
func Dial(rawUrl string) (*ethclient.Client, error) {
	return DialContext(context.Background(), rawUrl)
}

// DialContext connects a client to the given URL with context.
func DialContext(ctx context.Context, rawUrl string) (*ethclient.Client, error) {
	standClient := newRetryClient()
	standClient.Timeout = 1 * time.Minute

	rc, err := rpc.DialOptions(ctx, rawUrl, rpc.WithHTTPClient(standClient))
	if err != nil {
		return nil, err
	}
	return ethclient.NewClient(rc), nil
}
