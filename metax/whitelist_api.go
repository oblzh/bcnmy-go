package metax

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type AddDestinationRequest struct {
	DestinationAddresses []string `json:"destinationAddresses"`
}

type AddDestinationResponse struct {
	Code               int      `json:"code"`
	Message            string   `json:"message"`
	RegisteredCount    int      `json:"registeredCount"`
	DuplicateContracts []string `json:"duplicateContracts"`
	InvalidContracts   []string `json:"invalidContracts"`
}

type AddProxyContractsRequest struct {
	Addresses []string `json:"addresses"`
}

type ProxyContractsResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type GetProxyContractsResponse struct {
	ProxyContractsResponse
	Total     int `json:"total"`
	Addresses []struct {
		Address string `json:"address"`
		Status  bool   `json:"status"`
	} `json:"addresses"`
}

type PatchProxyContractsRequest struct {
	Status  int    `json:"status"` // 0 => inactive, 1 => active
	Address string `json:"address"`
}

func (b *Bcnmy) AddDestinationAddresses(data *AddDestinationRequest) (*AddDestinationResponse, error) {
	bodyCh := make(chan []byte)
	errorCh := make(chan error)
	defer close(bodyCh)
	defer close(errorCh)
	body, err := json.Marshal(data)
	if err != nil {
		b.logger.WithError(err).Error("json marshal `AddDestinationRequest` data failed")
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, AddDestinationAddressesURL, bytes.NewBuffer(body))
	if err != nil {
		b.logger.WithError(err).Error("AddDestinationAddresses NewRequest failed")
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", b.GetAuthorization())
	var resp AddDestinationResponse
	b.asyncHttpx(req, errorCh, bodyCh)
	select {
	case ret := <-bodyCh:
		err = json.Unmarshal(ret, &resp)
		if err != nil {
			return nil, fmt.Errorf("AddDestinationAddresses unmarshal failed, %v", err)
		}
		return &resp, nil
	case err := <-errorCh:
		b.logger.Error(err.Error())
		return nil, err
	}
}

func (b *Bcnmy) AddProxyContracts(data *AddProxyContractsRequest) (*ProxyContractsResponse, error) {
	bodyCh := make(chan []byte)
	errorCh := make(chan error)
	defer close(bodyCh)
	defer close(errorCh)

	body, err := json.Marshal(data)
	if err != nil {
		b.logger.WithError(err).Error("json marshal `AddProxyContractsRequest` data failed")
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, ProxyContractsURL, bytes.NewBuffer(body))
	if err != nil {
		b.logger.WithError(err).Error("AddProxyContracts NewRequest failed")
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", b.GetAuthorization())
	var resp ProxyContractsResponse
	b.asyncHttpx(req, errorCh, bodyCh)
	select {
	case ret := <-bodyCh:
		err = json.Unmarshal(ret, &resp)
		if err != nil {
			return nil, fmt.Errorf("AddProxyContracts unmarshal failed, %v", err)
		}
		return &resp, nil
	case err := <-errorCh:
		b.logger.Error(err.Error())
		return nil, err
	}
}

func (b *Bcnmy) PatchProxyContracts(data *PatchProxyContractsRequest) (*ProxyContractsResponse, error) {
	bodyCh := make(chan []byte)
	errorCh := make(chan error)
	defer close(bodyCh)
	defer close(errorCh)

	body, err := json.Marshal(data)
	if err != nil {
		b.logger.WithError(err).Error("json marshal `PatchProxyContractsRequest` data failed")
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPatch, ProxyContractsURL, bytes.NewBuffer(body))
	if err != nil {
		b.logger.WithError(err).Error("PatchProxyContracts NewRequest failed")
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", b.GetAuthorization())
	var resp ProxyContractsResponse
	b.asyncHttpx(req, errorCh, bodyCh)
	select {
	case ret := <-bodyCh:
		err = json.Unmarshal(ret, &resp)
		if err != nil {
			return nil, fmt.Errorf("PatchProxyContracts unmarshal failed, %v", err)
		}
		return &resp, nil
	case err := <-errorCh:
		b.logger.Error(err.Error())
		return nil, err
	}
}

func (b *Bcnmy) GetProxyContracts() (*GetProxyContractsResponse, error) {
	bodyCh := make(chan []byte)
	errorCh := make(chan error)
	defer close(bodyCh)
	defer close(errorCh)
	req, err := http.NewRequest(http.MethodGet, ProxyContractsURL, nil)
	if err != nil {
		b.logger.WithError(err).Error("GetProxyContracts NewRequest failed")
		return nil, err
	}
	req.Header.Set("Authorization", b.GetAuthorization())
	var resp GetProxyContractsResponse
	b.asyncHttpx(req, errorCh, bodyCh)
	select {
	case ret := <-bodyCh:
		err = json.Unmarshal(ret, &resp)
		if err != nil {
			return nil, fmt.Errorf("GetProxyContracts unmarshal failed, %v", err)
		}
		return &resp, nil
	case err := <-errorCh:
		b.logger.Error(err.Error())
		return nil, err
	}
}
