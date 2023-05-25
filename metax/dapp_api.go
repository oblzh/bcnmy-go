package metax

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

/*
https://docs-gasless.biconomy.io/api/check-limits
code

	200
	150	when DApp limits are exhausted
	151	when User limits are exhausted
	152	when API/User limits are exhausted

limit.type

	0	the limit was applied on gas usage
	1	the limit was applied on the number of meta transaction
*/
type CheckLimitResponse struct {
	Code         int       `json:"code"`
	Message      string    `json:"message"`
	ResponseCode int       `json:"responseCode"`
	Allowed      bool      `json:"allowed"`
	Limit        LimitInfo `json:"limit"`
}

type LimitInfo struct {
	Allowed   bool    `json:"allowed"`
	Type      int     `json:"type"`
	ResetTime int64   `json:"resetTime"`
	LimitLeft float32 `json:"limitLeft"`
}

func (b *Bcnmy) CheckLimits(from string, method string) (*CheckLimitResponse, error) {
	apiId, ok := b.apiID[fmt.Sprintf("%s-%s", b.address.Hex(), method)]
	if !ok {
		err := fmt.Errorf("ApiId not found for %s", method)
		b.logger.Error(err.Error())
		return nil, err
	}

	bodyCh := make(chan []byte)
	errorCh := make(chan error)
	defer close(bodyCh)
	defer close(errorCh)

	values := url.Values{
		"userAddress": {from},
		"apiId":       {apiId.ID},
	}
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s?%s", CheckLimitURL, values.Encode()),
		nil,
	)
	if err != nil {
		b.logger.WithError(err).Error("CheckLimits NewRequest failed")
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("x-api-key", b.apiKey)
	var resp CheckLimitResponse
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
