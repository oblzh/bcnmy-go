package test

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/oblzh/bcnmy-go/abi/demo"
	"github.com/oblzh/bcnmy-go/metax"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestSendTransactionStatus(t *testing.T) {
	b, _ := metax.NewBcnmy(os.Getenv("httpRpc"), os.Getenv("apiKey"), time.Second*10)
	b.WithDapp(demo.TransferDemoMetaData.ABI, common.HexToAddress("0x56b71565f6e7f9de4c3217a6e5d4133bc7fc67eb"))
	resp, err := b.GetTransactionStatus("0xfe7db1e9b66adf9aff979062f0475505194a027a29757401d3f69cb9b5cfdf1b")
	assert.Nil(t, err)
	fmt.Println(resp)
}
