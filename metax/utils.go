package metax

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// WaitMined waits for tx to be mined on the blockchain.
// It stops waiting when the context is canceled.
func WaitMined(ctx context.Context, b bind.DeployBackend, txHash common.Hash) (*types.Receipt, error) {
	queryTicker := time.NewTicker(time.Second)
	defer queryTicker.Stop()

	for {
		receipt, err := b.TransactionReceipt(ctx, txHash)
		if err == nil {
			return receipt, nil
		}

		if errors.Is(err, ethereum.NotFound) {
			log.Printf("Transaction %s not yet mined\n", txHash)
		} else {
			log.Printf("Receipt retrieval failed: %v\n", err)
		}

		// Wait for the next round.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-queryTicker.C:
		}
	}
}

func ConvertToJsonStr(obj interface{}) string {
	jsonStr, err := json.MarshalIndent(obj, "", " ")
	if err != nil {
		return ""
	} else {
		return string(jsonStr)
	}
}
