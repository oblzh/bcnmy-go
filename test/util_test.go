package test

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"

	metax "github.com/oblzh/bcnmy-go/metax"
)

func TestDomainDemo(t *testing.T) {
	fmt.Println(metax.GetDomainSeparator(common.HexToAddress("0x84a0856b038eaAd1cC7E297cF34A7e72685A8693"), big.NewInt(1)))
}
