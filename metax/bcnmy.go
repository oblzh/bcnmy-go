package metax

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"

	"github.com/oblzh/bcnmy-go/abi/forwarder"
)

type Bcnmy struct {
	ctx    context.Context
	logger *logrus.Entry

	ethClient    *ethclient.Client
	sleepTimeSec time.Duration
	httpClient   *http.Client

	// DAPP abi and address
	abi     abi.ABI
	address common.Address

	authToken string
	apiKey    string
	/// method apiID
	apiID map[string]struct {
		ID              string
		ContractAddress string
	}

	batchId *big.Int
	chainId *big.Int

	trustedForwarder struct {
		Address  common.Address
		Contract *forwarder.Forwarder
	}

	// backend config
	email             string
	password          string
	backendHttpClient *http.Client
}

func NewBcnmy(httpRpc string, apiKey string, timeout time.Duration) (*Bcnmy, error) {
	var err error
	bcnmy := &Bcnmy{
		ctx:    context.Background(),
		logger: logrus.WithField("metax", "bcnmy"),
		apiKey: apiKey,
		apiID: make(map[string]struct {
			ID              string
			ContractAddress string
		}),
		batchId:      big.NewInt(0),
		httpClient:   &http.Client{Timeout: timeout},
		sleepTimeSec: time.Duration(5),
	}
	bcnmy.ethClient, err = ethclient.DialContext(bcnmy.ctx, httpRpc)
	if err != nil {
		bcnmy.logger.WithError(err).Error("DialContext ethclient failed")
		return nil, err
	}
	bcnmy.chainId, err = bcnmy.ethClient.ChainID(bcnmy.ctx)
	if err != nil {
		bcnmy.logger.WithError(err).Error("ethClient getchainId failed")
		return nil, err
	}

	forwarderAddress, ok := ForwarderAddressMap[bcnmy.chainId.String()]
	if !ok {
		err = fmt.Errorf("Chain ID not supported: %v", bcnmy.chainId)
		bcnmy.logger.Error(err.Error())
		return nil, err
	}

	forwarderContract, err := forwarder.NewForwarder(forwarderAddress, bcnmy.ethClient)
	if err != nil {
		bcnmy.logger.WithError(err).Error("Load Forwarder Contract failed")
		return nil, err
	}

	bcnmy.trustedForwarder = struct {
		Address  common.Address
		Contract *forwarder.Forwarder
	}{
		Address:  forwarderAddress,
		Contract: forwarderContract,
	}
	resp, err := bcnmy.GetMetaAPI(bcnmy.ctx)
	if err != nil {
		bcnmy.logger.WithError(err).Error(err.Error())
		return nil, err
	}
	for _, info := range resp.ListAPI {
		// filter non contractAddress
		if common.IsHexAddress(info.ContractAddress) {
			bcnmy.apiID[fmt.Sprintf("%s-%s", common.HexToAddress(info.ContractAddress).Hex(), info.Method)] = struct {
				ID              string
				ContractAddress string
			}{
				ID:              info.ID,
				ContractAddress: info.ContractAddress,
			}
		}
	}
	return bcnmy, nil
}

func (b *Bcnmy) WithDapp(jsonABI string, dappAddress common.Address) (*Bcnmy, error) {
	var err error
	b.address = dappAddress
	b.abi, err = abi.JSON(strings.NewReader(jsonABI))
	if err != nil {
		b.logger.WithError(err).Error("jsonABI parse failed")
		return nil, err
	}
	return b, nil
}

func (b *Bcnmy) WithAuthToken(authToken string) *Bcnmy {
	b.authToken = authToken
	return b
}

func (b *Bcnmy) WithFieldTimeout(timeout time.Duration) *Bcnmy {
	b.httpClient = &http.Client{Timeout: timeout}
	return b
}

func (b *Bcnmy) WithSleepTimeSec(sleepTimeSec int64) *Bcnmy {
	b.sleepTimeSec = time.Duration(sleepTimeSec)
	return b
}

func (b *Bcnmy) GetAuthorization() string {
	return fmt.Sprintf("User %s", b.authToken)
}

func (b *Bcnmy) WithBackend(email string, password string, timeout time.Duration) error {
	b.email = email
	b.password = password

	jar, err := cookiejar.New(nil)
	if err != nil {
		return err
	}
	b.backendHttpClient = &http.Client{
		Jar:     jar,
		Timeout: timeout,
	}
	return nil
}
