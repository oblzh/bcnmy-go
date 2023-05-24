package test

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	metax "github.com/oblzh/bcnmy-go/metax"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestConvertToJson(t *testing.T) {

	metaTxMessage := &metax.MetaTxMessage{
		From:          common.HexToAddress("0xD1cc56810a3947d1D8b05448afB9889c6cFCF0F1"),
		To:            common.HexToAddress("0x56b71565f6e7f9de4c3217a6e5d4133bc7fc67eb"),
		Token:         common.HexToAddress("0x0000000000000000000000000000000000000000"),
		TxGas:         150000,
		TokenGasPrice: "0",
		BatchId:       big.NewInt(0),
		BatchNonce:    big.NewInt(19),
		Deadline:      big.NewInt(1684815127),
		Data:          "0x71234eb00000000000000000000000006a22dda833c14ca6189f32e0dbcdf41ac2a3c951000000000000000000000000c015fb756fd4d49c6280eca2d47df30e8f6d083100000000000000000000000000000000000000000000000000000000000186a000000000000000000000000000000000000000000000000000000000000186a0000000000000000000000000000000000000000000000000000000000000000f00000000000000000000000000000000000000000000000000000000646c3d170000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001c3aee4900e4d2df6cd60bab10699f3a4336ea527db9f20c3a13905167cce340d3018f6c8bb754c437bfb7e9ad291ff468e34e94e0169d8d9509eb0ef3fca1415c",
	}

	typedData := apitypes.TypedData{
		Types:       metax.SignedTypes,
		PrimaryType: metax.ForwardRequestType,
		Domain: apitypes.TypedDataDomain{
			Name:              metax.ForwardRequestName,
			Version:           metax.Version,
			VerifyingContract: common.HexToAddress("0x69015912AA33720b842dCD6aC059Ed623F28d9f7").Hex(),
			Salt:              hexutil.Encode(common.LeftPadBytes(big.NewInt(80001).Bytes(), 32)),
		},
		Message: metaTxMessage.TypedData(),
	}
	typedDataHash, _ := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	signature := hexutil.MustDecode("0xf38bbb6b1af600828c95711b6e5ca4eb8119739ceae84e3ad032aa3fc886c1777e9dbf4b91a7fa508aa5041f044820afa92fa0db31b9be6ac81f5b3d6090f24d1c")
	fmt.Println(signature)
	fmt.Println(typedDataHash)

	s := metax.ConvertToJsonStr(metaTxMessage)
	assert.NotNil(t, s)
	fmt.Println(s)
}
