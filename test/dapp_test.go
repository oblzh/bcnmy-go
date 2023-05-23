package test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/ethereum/go-ethereum/common"

	metax "github.com/oblzh/bcnmy-go/metax"
	demo "github.com/oblzh/bcnmy-go/abi/demo"
)

func TestCheckLimits(t *testing.T) {
	b, _ := metax.NewBcnmy(os.Getenv("httpRpc"), os.Getenv("apiKey"), time.Second*10)
    b.WithDapp(demo.TransferDemoABI, common.HexToAddress("0x56b71565f6e7f9de4c3217a6e5d4133bc7fc67eb"))
	resp, err := b.CheckLimits("0x96774c64dc3f46f64d17034ce6cf7b2ef31da56a", "transfer")
	assert.Nil(t, err)
	fmt.Println(resp)
}
