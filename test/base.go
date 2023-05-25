package test

import (
	"github.com/oblzh/bcnmy-go/metax"
	"os"
	"time"
)

func buildBcnmy() *metax.Bcnmy {
	b, _ := metax.NewBcnmy(os.Getenv("httpRpc"), os.Getenv("apiKey"), 10*time.Second)
	b = b.WithAuthToken(os.Getenv("authToken"))
	return b
}
