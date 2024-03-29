package test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	token "github.com/oblzh/bcnmy-go/abi/token"
	metax "github.com/oblzh/bcnmy-go/metax"
)

// Finished
func TestDeleteContract(t *testing.T) {
	b := buildBcnmy()

	data := &metax.DeleteContractRequest{
		ContractAddress: "0x0a364431476a8d1dd475590b0a028b40686ce542",
		ContractType:    "SC",
	}

	resp, err := b.DeleteContract(data)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, resp.Code, 143)
}

// Finished
func TestDeleteMethod(t *testing.T) {
	b := buildBcnmy()

	data := &metax.DeleteMethodRequest{
		ContractAddress: "0xa6b71e26c5e0845f74c812102ca7114b6a896ab2",
		Method:          "createProxyWithNonce",
	}

	resp, err := b.DeleteMethod(data)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, resp.Code, 143)
}

// Finished
func TestAddContract(t *testing.T) {
	b := buildBcnmy()
	assert.NotNil(t, b)

	data := &metax.AddContractRequest{
		ContractName:        "TestToken",
		ContractAddress:     "0xeaC94633FFf8C65aD9EFdCF237741D931fa995Cd",
		ContractType:        "SC",
		WalletType:          "",
		MetaTransactionType: "DEFAULT",
		ABI:                 token.TestTokenABI,
	}

	resp, err := b.AddContract(data)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, resp.Code, 200)
}

// Finished
func TestAddMethod(t *testing.T) {
	b := buildBcnmy()
	assert.NotNil(t, b)

	data := &metax.AddMethodRequest{
		ContractAddress: "0xeaC94633FFf8C65aD9EFdCF237741D931fa995Cd",
		ApiType:         "custom",
		Name:            "mintTo",
		MethodType:      "write",
		Method:          "mintTo",
	}

	resp, err := b.AddMethod(data)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, resp.Code, 200)
}

// Finished
func TestCreateDapp(t *testing.T) {
	b := buildBcnmy()
	assert.NotNil(t, b)

	data := &metax.CreateDappRequest{
		DappName:             "test-create",
		NetworkId:            "5",
		EnableBiconomyWallet: false,
	}

	resp, err := b.CreateDapp(data)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, resp.Code, 200)
}
