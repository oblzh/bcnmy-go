package metax

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type MetaAPIResponse struct {
	Log     string        `json:"log"`
	Flag    int           `json:"flag"`
	Total   int           `json:"total"`
	ListAPI []MetaAPIInfo `json:"listApis"`
}

type MetaAPIInfo struct {
	/// need to filter non contractAdress
	ContractAddress   string      `json:"contractAddress"`
	ID                string      `json:"id"`
	Name              string      `json:"name"`
	URL               string      `json:"url"`
	Version           int         `json:"version"`
	Method            string      `json:"method"`
	MethodType        string      `json:"methodType"`
	APIType           string      `json:"apiType"`
	MetaTxLimitStatus int         `json:"metaTxLimitStatus"`
	MetaTxLimit       MetaTxLimit `json:"metaTxLimit"`
}

type MetaTxLimit struct {
	Type              int     `json:"type"`
	Value             float32 `json:"value"`
	DurationValue     int     `json:"durationValue"`
	Day               string  `json:"day"`
	LimitStartTime    int64   `json:"limitStartTime"`
	LimitDurationInMs int64   `json:"limitDurationInMs"`
}

func (b *Bcnmy) GetMetaAPI(ctx context.Context) (*MetaAPIResponse, error) {
	bodyCh := make(chan []byte)
	errorCh := make(chan error)
	defer close(bodyCh)
	defer close(errorCh)

	req, err := http.NewRequest(http.MethodGet, MetaAPIURL, nil)
	if err != nil {
		b.logger.WithError(err).Error("MetaAPI NewRequest failed")
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("x-api-key", b.apiKey)
	var resp MetaAPIResponse
	b.asyncHttpx(req, errorCh, bodyCh)
	select {
	case ret := <-bodyCh:
		err = json.Unmarshal(ret, &resp)
		if err != nil {
			return nil, fmt.Errorf("AddDestinationAddresses unmarshal failed, %v", err)
		}
		if resp.Flag != 143 {
			err := fmt.Errorf("%v", resp)
			b.logger.WithError(err).Error(resp.Log)
			return nil, err
		}
		return &resp, nil
	case err := <-errorCh:
		b.logger.Error(err.Error())
		return nil, err
	}
}
