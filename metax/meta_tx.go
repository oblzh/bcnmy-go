package metax

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

type MetaTxMessage struct {
	From          common.Address `json:"from"`
	To            common.Address `json:"to"`
	Token         common.Address `json:"token"`
	TxGas         uint64         `json:"txGas"`
	TokenGasPrice string         `json:"tokenGasPrice"`
	BatchId       *big.Int       `json:"batchId"`
	BatchNonce    *big.Int       `json:"batchNonce"`
	Deadline      *big.Int       `json:"deadline"`
	Data          string         `json:"data"`
}

type MetaTxRequest struct {
	From          string        `json:"from"`
	To            string        `json:"to"`
	ApiID         string        `json:"apiId"`
	Params        []interface{} `json:"params"`
	SignatureType string        `json:"signatureType"`
}

type MetaTxResponse struct {
	TxHash  common.Hash `json:"txHash"`
	Log     string      `json:"log"`
	Flag    int         `json:"flag"`
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Code    int         `json:"code"`
	Limit   struct {
		Type      int     `json:"type"`
		LimitLeft float32 `json:"limitLeft"`
		ResetTime int64   `json:"resetTime"`
	} `json:"limit"`
	Allowed bool `json:"allowed"`
}

func (m *MetaTxMessage) TypedData() apitypes.TypedDataMessage {
	return apitypes.TypedDataMessage{
		"from":          m.From.Hex(),
		"to":            m.To.Hex(),
		"token":         m.Token.Hex(),
		"txGas":         hexutil.EncodeUint64(m.TxGas),
		"tokenGasPrice": m.TokenGasPrice,
		"batchId":       m.BatchId.String(),
		"batchNonce":    m.BatchNonce.String(),
		"deadline":      m.Deadline.String(),
		"data":          m.Data,
	}
}

func (b *Bcnmy) SendMetaNativeTx(data *MetaTxRequest) (*MetaTxResponse, error) {
	bodyCh := make(chan []byte)
	errorCh := make(chan error)
	defer close(bodyCh)
	defer close(errorCh)

	body, err := json.Marshal(data)
	if err != nil {
		b.logger.WithError(err).Error("json marshal `MetaTxRequest` data failed")
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, MetaTxNativeURL, bytes.NewBuffer(body))
	if err != nil {
		b.logger.Error("SendMetaNativeTx NewRequest failed")
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("x-api-key", b.apiKey)
	var resp MetaTxResponse
	b.asyncHttpx(req, errorCh, bodyCh)
	select {
	case ret := <-bodyCh:
		err = json.Unmarshal(ret, &resp)
		if err != nil {
			return nil, fmt.Errorf("SendMetaNativeTx unmarshal failed, %v", err)
		}
		if resp.TxHash == common.HexToHash("0x0") {
			err := fmt.Errorf("Message: %s, Code: %v, Limit: %v", resp.Message, resp.Code, resp.Limit)
			return &resp, err
		}
		return &resp, nil
	case err := <-errorCh:
		b.logger.WithError(err).Error(err.Error())
		return nil, err
	}
}

func (b *Bcnmy) RawTransact(signer *Signer, method string, params ...interface{}) (*MetaTxResponse, *types.Transaction, *types.Receipt, error) {
	apiId, ok := b.apiID[fmt.Sprintf("%s-%s", b.address.Hex(), method)]
	if !ok {
		err := fmt.Errorf("ApiId %s not found for %s", apiId.ID, method)
		b.logger.Error(err.Error())
		return nil, nil, nil, err
	}
	funcSig, err := b.abi.Pack(method, params...)
	if err != nil {
		b.logger.WithError(err).Error("Abi Pack failed")
		return nil, nil, nil, err
	}

	callMsg := ethereum.CallMsg{
		From: signer.Address,
		To:   &b.address,
		Data: funcSig,
	}
	callOpts := bind.CallOpts{
		Context: b.ctx,
		From:    signer.Address,
	}
	estimateGas, err := b.ethClient.EstimateGas(b.ctx, callMsg)
	if err != nil {
		b.logger.WithError(err).Error("EstimateGas failed")
		return nil, nil, nil, err
	}
	batchNonce, err := b.trustedForwarder.Contract.GetNonce(&callOpts, signer.Address, b.batchId)
	if err != nil {
		b.logger.WithError(err).Errorf("GetNonce from %s failed", b.batchId)
		return nil, nil, nil, err
	}

	metaTxMessage := &MetaTxMessage{
		From:          signer.Address,
		To:            b.address,
		Token:         common.HexToAddress("0x0"),
		TxGas:         estimateGas,
		TokenGasPrice: "0",
		BatchId:       b.batchId,
		BatchNonce:    batchNonce,
		Deadline:      big.NewInt(time.Now().Add(time.Hour).Unix()),
		Data:          hexutil.Encode(funcSig),
	}

	typedData := apitypes.TypedData{
		Types:       SignedTypes,
		PrimaryType: ForwardRequestType,
		Domain: apitypes.TypedDataDomain{
			Name:              ForwardRequestName,
			Version:           Version,
			VerifyingContract: b.trustedForwarder.Address.Hex(),
			Salt:              hexutil.Encode(common.LeftPadBytes(b.chainId.Bytes(), 32)),
		},
		Message: metaTxMessage.TypedData(),
	}
	signature, err := signer.SignTypedData(typedData)
	if err != nil {
		b.logger.WithError(err).Error("Signer signTypeData failed")
		return nil, nil, nil, err
	}

	domainSeparator, err := typedData.HashStruct(EIP712DomainType, typedData.Domain.Map())
	if err != nil {
		b.logger.WithError(err).Error("EIP712Domain Separator hash failed")
		return nil, nil, nil, err
	}

	req := &MetaTxRequest{
		From:  signer.Address.Hex(),
		To:    b.address.Hex(),
		ApiID: apiId.ID,
		Params: []interface{}{
			metaTxMessage,
			hexutil.Encode(domainSeparator),
			hexutil.Encode(signature),
		},
		SignatureType: SignatureEIP712Type,
	}

	b.logger.Debugf("MetaTxRequest: %s", ConvertToJsonStr(req))
	b.logger.Debugf("MetaTxMessage: %s", ConvertToJsonStr(metaTxMessage))

	resp, err := b.SendMetaNativeTx(req)
	if err != nil {
		b.logger.Errorf("Transaction failed: %v", err)
		return resp, nil, nil, err
	}

	receipt, err := WaitMined(context.Background(), b.ethClient, resp.TxHash)
	if err != nil {
		b.logger.Errorf("WaitMined failed: %v", err)
		return resp, nil, nil, err
	}

	var tx *types.Transaction
	retries := 5
	for {
		var err error
		retries -= 1
		tx, _, err = b.ethClient.TransactionByHash(b.ctx, resp.TxHash)
		if err != nil {
			b.logger.Errorf("Checking TransactionByHash failed: %v, retries: %v", err, retries)
			time.Sleep(time.Second * b.sleepTimeSec)
		} else {
			break
		}
		if retries < 0 && err != nil {
			return resp, nil, nil, err
		}
	}
	return resp, tx, receipt, err
}

func (b *Bcnmy) BuildTransactParams(metaTxMessage *MetaTxMessage, typedDataHash string) ([]byte, error) {
	typedData := apitypes.TypedData{
		Types:       SignedTypes,
		PrimaryType: ForwardRequestType,
		Domain: apitypes.TypedDataDomain{
			Name:              ForwardRequestName,
			Version:           Version,
			VerifyingContract: b.trustedForwarder.Address.Hex(),
			Salt:              hexutil.Encode(common.LeftPadBytes(b.chainId.Bytes(), 32)),
		},
		Message: metaTxMessage.TypedData(),
	}
	hash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		b.logger.Errorf("HashStruct failed to hash typedData, %v", err)
		return nil, err
	}
	if hash.String() != typedDataHash {
		err := fmt.Errorf("Hash string not match parameter hash: %s typedDataHash %s", hash.String(), typedDataHash)
		b.logger.Errorf("%v", err)
		return nil, err
	}

	return typedData.HashStruct(EIP712DomainType, typedData.Domain.Map())
}

// / Backend using this method, handle frontend passing signature, MetaTxMessage and
// / ForwardRequestType data Hash value
func (b *Bcnmy) EnhanceTransact(from string, method string, signature []byte, metaTxMessage *MetaTxMessage, typedDataHash string) (*MetaTxResponse, *types.Transaction, *types.Receipt, error) {
	apiId, ok := b.apiID[fmt.Sprintf("%s-%s", b.address.Hex(), method)]
	if !ok {
		err := fmt.Errorf("ApiId %s not found for %s", apiId.ID, method)
		b.logger.Error(err.Error())
		return nil, nil, nil, err
	}
	domainSeparator, err := b.BuildTransactParams(metaTxMessage, typedDataHash)
	if err != nil {
		b.logger.WithError(err).Error("EIP712Domain Separator hash failed")
		return nil, nil, nil, err
	}
	req := &MetaTxRequest{
		From:  from,
		To:    b.address.Hex(),
		ApiID: apiId.ID,
		Params: []interface{}{
			metaTxMessage,
			hexutil.Encode(domainSeparator),
			hexutil.Encode(signature),
		},
		SignatureType: SignatureEIP712Type,
	}
	resp, err := b.SendMetaNativeTx(req)
	if err != nil {
		b.logger.Errorf("Transaction failed: %v", err)
		return resp, nil, nil, err
	}

	b.logger.Debugf("MetaTxRequest: %s", ConvertToJsonStr(req))
	b.logger.Debugf("MetaTxMessage: %s", ConvertToJsonStr(metaTxMessage))

	receipt, err := WaitMined(context.Background(), b.ethClient, resp.TxHash)
	if err != nil {
		b.logger.Errorf("WaitMined failed: %v", err)
		return resp, nil, nil, err
	}

	var tx *types.Transaction
	retries := 5
	for {
		var err error
		retries -= 1
		tx, _, err = b.ethClient.TransactionByHash(b.ctx, resp.TxHash)
		if err != nil {
			b.logger.Errorf("Checking TransactionByHash failed: %v, retries: %v", err, retries)
			time.Sleep(time.Second * b.sleepTimeSec)
		} else {
			break
		}
		if retries < 0 && err != nil {
			return resp, nil, nil, err
		}
	}
	return resp, tx, receipt, err
}

func (b *Bcnmy) Pack(method string, params ...interface{}) ([]byte, error) {
	data, err := b.abi.Pack(method, params...)
	if err != nil {
		b.logger.WithError(err).Error("Abi Pack failed")
		return nil, err
	}
	return data, err
}
