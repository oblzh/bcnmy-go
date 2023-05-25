package metax

import (
	"fmt"
	"io"
	"net/http"
)

func (b *Bcnmy) asyncHttpx(req *http.Request, errorCh chan error, bodyCh chan []byte) {
	go func() {
		res, err := b.httpClient.Do(req)
		if err != nil {
			b.logger.WithError(err).Error("HttpClient request to failed")
			errorCh <- err
			return
		}
		defer res.Body.Close()
		replyData, err := io.ReadAll(res.Body)
		if err != nil {
			b.logger.WithError(err).Error("io read request body failed")
			errorCh <- err
			return
		}
		bodyCh <- replyData
	}()
}

func (b *Bcnmy) backendAsyncHttpx(req *http.Request, errorCh chan error, bodyCh chan []byte) {
	go func() {
		res, err := b.backendHttpClient.Do(req)
		if err != nil {
			b.logger.WithError(err).Error("HttpClient request to failed")
			errorCh <- err
			return
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			errorCh <- fmt.Errorf("%v", res.StatusCode)
			return
		}
		replyData, err := io.ReadAll(res.Body)
		if err != nil {
			b.logger.WithError(err).Error("io read request body failed")
			errorCh <- err
			return
		}
		bodyCh <- replyData
	}()
}
