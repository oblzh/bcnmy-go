package metax

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"net/http"
	"net/url"
	"time"
)

type MetaTxResponseV1 struct {
	Flag int         `json:"flag"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type MetaTxV1SuccessData struct {
	TransactionId string `json:"transactionId"`
	ConnectionUrl string `json:"connectionUrl"`
}

type BiconomyTransaction struct {
	Flag int    `json:"flag"`
	Log  string `json:"log"`
	Code int    `json:"code"`
	Data struct {
		Status  string `json:"status"`
		Receipt struct {
			TxHash string `json:"transactionHash"`
		} `json:"receipt"`
	} `json:"data"`
}

func (b *Bcnmy) SendMetaNativeTxV1(data *MetaTxRequest) (*MetaTxResponse, error) {
	bodyCh := make(chan []byte)
	errorCh := make(chan error)
	defer close(bodyCh)
	defer close(errorCh)

	body, err := json.Marshal(data)
	if err != nil {
		b.logger.WithError(err).Error("json marshal `MetaTxRequest` data failed")
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, MetaTxNativeURLV1, bytes.NewBuffer(body))
	if err != nil {
		b.logger.Error("SendMetaNativeTxV1 NewRequest failed")
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("x-api-key", b.apiKey)
	req.Header.Set("version", PACKAGE_VERSION)
	var resp MetaTxResponseV1
	b.asyncHttpx(req, errorCh, bodyCh)
	var successData MetaTxV1SuccessData
	select {
	case ret := <-bodyCh:
		err = json.Unmarshal(ret, &resp)
		if err != nil {
			return nil, fmt.Errorf("SendMetaNativeTxV1 unmarshal failed, %v", err)
		}
		if resp.Flag == 200 {
			dataBytes, err := json.Marshal(resp.Data)
			if err != nil {
				return nil, fmt.Errorf("Error marshaling response data:", err)
			}
			if err = json.Unmarshal(dataBytes, &successData); err != nil {
				return nil, fmt.Errorf("Error unmarshaling response data:", err)
			}
			break
		} else if resp.Flag == 400 {
			switch data := resp.Data.(type) {
			case string:
				return &MetaTxResponse{
					Flag:    resp.Flag,
					TxHash:  common.HexToHash("0x0"),
					Error:   data,
					Message: data,
					Code:    resp.Flag,
				}, nil
			case map[string]interface{}:
				return &MetaTxResponse{
					Flag:    resp.Flag,
					TxHash:  common.HexToHash("0x0"),
					Error:   data["error"].(string),
					Message: data["error"].(string),
					Code:    data["code"].(int),
				}, nil
			}
		}
	case err := <-errorCh:
		b.logger.WithError(err).Error(err.Error())
		return nil, err
	}
	bcnmyTxn, err := b.GetTransactionStatus(successData.TransactionId)
	if err != nil {
		b.logger.Error("SendTransactionStatus failed, check error logs")
		return nil, err
	} else {
		return &MetaTxResponse{
			Flag:   bcnmyTxn.Flag,
			Code:   bcnmyTxn.Code,
			TxHash: common.HexToHash(bcnmyTxn.Data.Receipt.TxHash),
		}, nil
	}
}

func (b *Bcnmy) GetTransactionStatus(transactionId string) (*BiconomyTransaction, error) {
	bodyCh := make(chan []byte, 1)
	errorCh := make(chan error)
	defer close(bodyCh)
	defer close(errorCh)
	queryParams := url.Values{}
	queryParams.Set("transactionId", transactionId)
	urlWithParams := MetaTransactionStatusURL + "?" + queryParams.Encode()
	req, err := http.NewRequest(http.MethodGet, urlWithParams, nil)
	if err != nil {
		b.logger.WithError(err).Error("SendTransactionStatus NewRequest failed")
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("x-api-key", b.apiKey)
	req.Header.Set("version", PACKAGE_VERSION)
	var resp BiconomyTransaction
	b.asyncHttpx(req, errorCh, bodyCh)
	retries := 5
LOOP:
	for {
		retries -= 1
		if retries <= 0 {
			return nil, fmt.Errorf("GetBiconomyTransactionStatus %s reach maxiumum retries", transactionId)
		}
		select {
		case ret := <-bodyCh:
			err = json.Unmarshal(ret, &resp)
			if err != nil {
				return nil, fmt.Errorf("GetTransactionStatus unmarshal failed, %v", err)
			}
			if resp.Code != 200 {
				b.logger.Infof("BiconomyTransaction %v Code is empty", resp)
				time.Sleep(b.sleepTimeSec * time.Second)
				continue LOOP
			}
			if resp.Data.Receipt.TxHash == "" || resp.Data.Receipt.TxHash == common.HexToHash("0x0").Hex() {
				b.logger.Infof("BiconomyTransaction %v txHash is empty", resp)
				time.Sleep(b.sleepTimeSec * time.Second)
				continue LOOP
			}
			return &resp, nil
		case err := <-errorCh:
			b.logger.Error(err.Error())
			time.Sleep(b.sleepTimeSec * time.Second)
			continue LOOP
		}
	}
}
