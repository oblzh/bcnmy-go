package metax

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"log"
	"math/big"
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

func GetDomainSeparator(forwarderAddress common.Address, chainId *big.Int) (common.Hash, error) {
	typedData := apitypes.TypedData{
		Types:       SignedTypes,
		PrimaryType: EIP712DomainType,
		Domain: apitypes.TypedDataDomain{
			Name:              ForwardRequestName,
			Version:           Version,
			VerifyingContract: forwarderAddress.Hex(),
			Salt:              hexutil.Encode(common.LeftPadBytes(chainId.Bytes(), 32)),
		},
		Message: make(map[string]interface{}),
	}
	domainSeparator, err := typedData.HashStruct(EIP712DomainType, typedData.Domain.Map())
	return common.HexToHash(hexutil.Encode(domainSeparator)), err
}
