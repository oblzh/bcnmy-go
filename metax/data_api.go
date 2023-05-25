package metax

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"net/url"
	"strings"
)

type UniqueUserDataRequest struct {
	StartDate string `json:"startDate"` /// Format (“MM-DD-YYYY”) example: 21st Jan 2022 would be 01-21-2022
	EndDate   string `json:"endDate"`   /// Format (“MM-DD-YYYY”)
}

type UniqueUserDataResponse struct {
	GeneralResponse
	UniqueUserData []struct {
		Date      string   `json:"date"`
		Count     int      `json:"count"`
		Addresses []string `json:"addresses"`
	}
}

type UserLimitRequest struct {
	SignerAddress string `json:"signerAddress"`
	ApiId         string `json:"apiId"`
}

type UserLimitResponse struct {
	GeneralResponse
	UserLimitData struct {
		LimitLeft struct {
			SignerAddress        string  `json:"signerAddress"`
			TransactionLimitLeft float32 `json:"transactionLimitLeft"`
			TransactionCount     int     `json:"transactionCount"`
			AreLimitsConsumed    bool    `json:"areLimitsConsumed"`
			UserTransactionLimit int     `json:"userTransactionLimit"`
		} `json:"limitLeft"`
		LimitType        string   `json:"limitType"`
		LimitStartTime   *big.Int `json:"limitStartTime"`
		LimitEndTime     *big.Int `json:"limitEndTime"`
		TimePeriodInDays int      `json:"timePeriodInDays"`
	} `json:"userLimitData"`
}

func (b *Bcnmy) GetUniqueUserData(data *UniqueUserDataRequest) (*UniqueUserDataResponse, error) {
	bodyCh := make(chan []byte)
	errorCh := make(chan error)
	defer close(errorCh)
	defer close(bodyCh)

	body := url.Values{
		"startDate": {data.StartDate},
		"endDate":   {data.EndDate},
	}
	req, err := http.NewRequest(http.MethodGet, UniqueUserDataURL, strings.NewReader(body.Encode()))
	if err != nil {
		b.logger.WithError(err).Error("GetUniqueUserData NewRequest failed")
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("authToken", b.authToken)
	req.Header.Set("apiKey", b.apiKey)
	var resp UniqueUserDataResponse
	b.asyncHttpx(req, errorCh, bodyCh)
	select {
	case ret := <-bodyCh:
		err = json.Unmarshal(ret, &resp)
		if err != nil {
			return nil, fmt.Errorf("GetUniqueUserData unmarshal failed, %v", err)
		}
		return &resp, nil
	case err := <-errorCh:
		b.logger.Error(err.Error())
		return nil, err
	}
}

func (b *Bcnmy) GetUserLimit(data *UserLimitRequest) (*UserLimitResponse, error) {
	bodyCh := make(chan []byte)
	errorCh := make(chan error)
	defer close(errorCh)
	defer close(bodyCh)

	body := url.Values{
		"signerAddress": {data.SignerAddress},
		"apiId":         {data.ApiId},
	}
	req, err := http.NewRequest(http.MethodGet, UserLimitURL, strings.NewReader(body.Encode()))
	if err != nil {
		b.logger.WithError(err).Error("GetUserLimit NewRequest failed")
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("authToken", b.authToken)
	req.Header.Set("apiKey", b.apiKey)
	var resp UserLimitResponse
	b.asyncHttpx(req, errorCh, bodyCh)
	select {
	case ret := <-bodyCh:
		err = json.Unmarshal(ret, &resp)
		if err != nil {
			return nil, fmt.Errorf("GetUserLimit unmarshal failed, %v", err)
		}
		return &resp, nil
	case err := <-errorCh:
		b.logger.Error(err.Error())
		return nil, err
	}
}
