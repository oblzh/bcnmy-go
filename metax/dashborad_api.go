package metax

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type GeneralResponse struct {
	Code         int    `json:"code"`
	Message      string `json:"message"`
	ResponseCode int    `json:"responseCode"`
}

type CreateDappRequest struct {
	DappName             string `json:"dappName"`
	NetworkId            string `json:"networkId"`
	EnableBiconomyWallet bool   `json:"enableBiconomyWallet"`
}

type CreateDappResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ApiKey     string   `json:"apiKey"`
		FundingKey *big.Int `json:"fundingKey"`
	} `json:"data"`
}

type AddContractRequest struct {
	ContractName        string `json:"contractName"`
	ContractAddress     string `json:"contractAddress"`
	ContractType        string `json:"contractType"`        // SCW for contract wallet or SC for contract
	WalletType          string `json:"walletType"`          // SCW or GNOSIS or blank
	MetaTransactionType string `json:"metaTransactionType"` // DEFAULT, TRUSTED_FORWARDER, ERC20_FORWARDER
	ABI                 string `json:"abi"`
}

type AddMethodRequest struct {
	ApiType         string `json:"apiType"`
	MethodType      string `json:"methodType"`
	Name            string `json:"name"`
	ContractAddress string `json:"contractAddress"`
	Method          string `json:"method"`
}

type AddMethodResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	ApiIds  []struct {
		ApiId  string `json:"apiId"`
		Method string `json:"method"`
		Name   string `json:"name"`
	} `json:"apiIds"`
}

type DeleteContractRequest struct {
	ContractAddress string `json:"contractAddress"`
	ContractType    string `json:"contractType"` // SCW for contract wallet or SC for contract
}

type DeleteMethodRequest struct {
	ContractAddress string `json:"contractAddress"`
	Method          string `json:"method"`
}

func (b *Bcnmy) CreateDapp(data *CreateDappRequest) (*CreateDappResponse, error) {
	bodyCh := make(chan []byte)
	errorCh := make(chan error)
	defer close(bodyCh)
	defer close(errorCh)
	body := url.Values{
		"dappName":             {data.DappName},
		"networkId":            {data.NetworkId},
		"enableBiconomyWallet": {strconv.FormatBool(data.EnableBiconomyWallet)},
	}
	req, err := http.NewRequest(http.MethodPost, CreateDappPublicURL, strings.NewReader(body.Encode()))
	if err != nil {
		b.logger.WithError(err).Error("CreateDapp NewRequest failed")
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("authToken", b.authToken)

	var resp CreateDappResponse
	b.asyncHttpx(req, errorCh, bodyCh)
	select {
	case ret := <-bodyCh:
		err = json.Unmarshal(ret, &resp)
		if err != nil {
			return nil, fmt.Errorf("CreateDapp unmarshal failed, %v", err)
		}
		return &resp, nil
	case err := <-errorCh:
		b.logger.Error(err.Error())
		return nil, err
	}
}

func (b *Bcnmy) AddContract(data *AddContractRequest) (*GeneralResponse, error) {
	bodyCh := make(chan []byte)
	errorCh := make(chan error)
	defer close(bodyCh)
	defer close(errorCh)

	body := url.Values{
		"contractName":        {data.ContractName},
		"contractAddress":     {data.ContractAddress},
		"contractType":        {data.ContractType},
		"walletType":          {data.WalletType},
		"metaTransactionType": {data.MetaTransactionType},
		"abi":                 {data.ABI},
	}
	req, err := http.NewRequest(http.MethodPost, AddContractURL, strings.NewReader(body.Encode()))
	if err != nil {
		b.logger.WithError(err).Error("AddContract NewRequest failed")
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("authToken", b.authToken)
	req.Header.Set("apiKey", b.apiKey)
	var resp GeneralResponse
	b.asyncHttpx(req, errorCh, bodyCh)
	select {
	case ret := <-bodyCh:
		err = json.Unmarshal(ret, &resp)
		if err != nil {
			return nil, fmt.Errorf("AddContract unmarshal failed, %v", err)
		}
		return &resp, nil
	case err := <-errorCh:
		b.logger.Error(err.Error())
		return nil, err
	}
}

func (b *Bcnmy) AddMethod(data *AddMethodRequest) (*AddMethodResponse, error) {
	bodyCh := make(chan []byte)
	errorCh := make(chan error)
	defer close(bodyCh)
	defer close(errorCh)

	body := url.Values{
		"apiType":         {data.ApiType},
		"methodType":      {data.MethodType},
		"name":            {data.Name},
		"contractAddress": {data.ContractAddress},
		"method":          {data.Method},
	}
	req, err := http.NewRequest(http.MethodPost, AddMethodURL, strings.NewReader(body.Encode()))
	if err != nil {
		b.logger.WithError(err).Error("AddMethod NewRequest failed")
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("authToken", b.authToken)
	req.Header.Set("apiKey", b.apiKey)
	var resp AddMethodResponse
	b.asyncHttpx(req, errorCh, bodyCh)
	select {
	case ret := <-bodyCh:
		err = json.Unmarshal(ret, &resp)
		if err != nil {
			return nil, fmt.Errorf("AddMethod unmarshal failed, %v", err)
		}
		return &resp, nil
	case err := <-errorCh:
		b.logger.Error(err.Error())
		return nil, err
	}
}

func (b *Bcnmy) DeleteContract(data *DeleteContractRequest) (*GeneralResponse, error) {
	body := url.Values{
		"contractAddress": {data.ContractAddress},
		"contractType":    {data.ContractType},
	}
	req, err := http.NewRequest(http.MethodDelete, DeleteContractURL, strings.NewReader(body.Encode()))
	if err != nil {
		b.logger.WithError(err).Error("DeleteContract NewRequest failed")
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("authToken", b.authToken)
	req.Header.Set("apiKey", b.apiKey)

	res, err := b.httpClient.Do(req)
	if err != nil {
		b.logger.WithError(err).Error("HttpClient request to DeleteContract failed")
		return nil, err
	}
	defer res.Body.Close()
	replyData, err := io.ReadAll(res.Body)
	var ret *GeneralResponse
	if err := json.Unmarshal(replyData, &ret); err != nil {
		b.logger.WithError(err).Error("json unmarshal body data failed")
		return nil, err
	}
	return ret, nil
}

func (b *Bcnmy) DeleteMethod(data *DeleteMethodRequest) (*GeneralResponse, error) {
	body := url.Values{
		"contractAddress": {data.ContractAddress},
		"method":          {data.Method},
	}
	req, err := http.NewRequest(http.MethodDelete, DeleteMethodURL, strings.NewReader(body.Encode()))
	if err != nil {
		b.logger.WithError(err).Error("DeleteMethod NewRequest failed")
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("authToken", b.authToken)
	req.Header.Set("apiKey", b.apiKey)

	res, err := b.httpClient.Do(req)
	if err != nil {
		b.logger.WithError(err).Error("HttpClient request to DeleteMethod failed")
		return nil, err
	}
	defer res.Body.Close()
	replyData, err := io.ReadAll(res.Body)
	var ret *GeneralResponse
	if err := json.Unmarshal(replyData, &ret); err != nil {
		b.logger.WithError(err).Error("json unmarshal body data failed")
		return nil, err
	}
	return ret, nil
}
