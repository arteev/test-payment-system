// +build integration

package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"syscall"
	"test-payment-system/internal/app/payment/database"
	"test-payment-system/internal/app/payment/di"
	"test-payment-system/internal/app/payment/dto"
	"testing"
	"time"

	"test-payment-system/internal/app/payment/config"
	"test-payment-system/internal/pkg/service"
	"test-payment-system/pkg/logger"
)

func TestPayment(t *testing.T) {
	suite.Run(t, &paymentSuite{})
}

type paymentSuite struct {
	suite.Suite
	dbMigrate database.Migrater
	db        database.Database
	port      int
}

func (s *paymentSuite) SetupSuite() {
	rand.Seed(time.Now().UTC().UnixNano())
	configFile := "payment_test.yaml"
	container := di.BuildContainer(configFile)
	err := container.Invoke(func(c *config.Config) error {
		c.Logger.Level = "error"
		s.port = c.API.Port
		return logger.SetupLogger(c.Mode, c.Logger)
	})
	s.Require().NoError(err)

	err = container.Invoke(func(db *database.DB) {
		s.dbMigrate = db
		s.db = db
	})
	s.Require().NoError(err)

	var svc *service.Service
	err = container.Invoke(func(s *service.Service) error {
		svc = s
		return nil
	})
	s.Require().NoError(err)
	go func() {
		err := svc.Start()
		s.Require().NoError(err)
	}()
	time.Sleep(time.Second * 2)
}

func (s *paymentSuite) TearDownSuite() {
	process, err := os.FindProcess(syscall.Getpid())
	s.Require().NoError(err)
	process.Signal(syscall.SIGINT)
}

func (s *paymentSuite) SetupTest() {
	s.dbMigrate.MigrateUp()
}

func (s *paymentSuite) TearDownTest() {
	s.dbMigrate.MigrateDown()
}

func (s *paymentSuite) Test_NewWallet() {
	urlAPI := fmt.Sprintf("http://localhost:%d/api/v1/payment/wallet", s.port)
	walletName := fmt.Sprintf("wallet-%d", rand.Uint64())
	payload := fmt.Sprintf(`{"name":"%s"}`, walletName)
	request, err := http.NewRequest(http.MethodPost, urlAPI, strings.NewReader(payload))
	s.NoError(err)

	request.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(request)
	s.NoError(err)
	s.Equal(http.StatusOK, resp.StatusCode)

	responseData, err := ioutil.ReadAll(resp.Body)
	s.NoError(err)
	defer resp.Body.Close()

	response, err := ParseResponse(responseData)
	s.NoError(err)
	expectSuccess(s.Suite, response)

	// check response
	gotWallet := &dto.WalletResponse{}
	err = json.Unmarshal(response.Data, gotWallet)
	s.Require().NoError(err)
	s.Equal(gotWallet.Name, walletName)
	s.Zero(gotWallet.Balance)
	s.NotZero(gotWallet.ID)

	// check from db
	ctx := context.Background()
	fromDBWallet, err := s.db.GetWallet(ctx, gotWallet.ID)
	s.Require().NoError(err)
	s.Equal(gotWallet.ID, fromDBWallet.ID)
	s.Equal(gotWallet.Name, fromDBWallet.Name)
	s.Equal(gotWallet.CreatedAt, fromDBWallet.CreatedAt.Unix())
	s.Equal(gotWallet.UpdatedAt, fromDBWallet.UpdatedAt.Unix())
}

func (s *paymentSuite) Test_GetWallet() {
	ctx := context.Background()
	walletName := fmt.Sprintf("wallet-%d", rand.Uint64())
	newWallet, err := s.db.NewWallet(ctx, walletName)
	s.Require().NoError(err)

	urlAPI := fmt.Sprintf("http://localhost:%d/api/v1/payment/wallet?id=%d", s.port, newWallet.ID)
	request, err := http.NewRequest(http.MethodGet, urlAPI, nil)
	s.NoError(err)

	request.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(request)
	s.NoError(err)
	s.Equal(http.StatusOK, resp.StatusCode)

	responseData, err := ioutil.ReadAll(resp.Body)
	s.NoError(err)
	defer resp.Body.Close()

	response, err := ParseResponse(responseData)
	s.NoError(err)
	expectSuccess(s.Suite, response)

	// check response
	gotWallet := &dto.WalletResponse{}
	err = json.Unmarshal(response.Data, gotWallet)
	s.Require().NoError(err)
	s.Zero(gotWallet.Balance)
	s.Equal(newWallet.ID, gotWallet.ID)
	s.Equal(newWallet.Name, gotWallet.Name)
	s.Equal(newWallet.CreatedAt.Unix(), gotWallet.CreatedAt)
	s.Equal(newWallet.UpdatedAt.Unix(), gotWallet.UpdatedAt)
	s.Equal(newWallet.Balance, gotWallet.Balance)
}

func (s *paymentSuite) Test_DepositWallet() {
	ctx := context.Background()
	walletName := fmt.Sprintf("wallet-%d", rand.Uint64())
	amount := 6.99

	newWallet, err := s.db.NewWallet(ctx, walletName)
	s.Require().NoError(err)

	payload := fmt.Sprintf(`{"wallet_id": %d, "amount":%f }`, newWallet.ID, amount)
	urlAPI := fmt.Sprintf("http://localhost:%d/api/v1/payment/deposit", s.port)
	request, err := http.NewRequest(http.MethodPost, urlAPI, strings.NewReader(payload))
	s.NoError(err)

	request.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(request)
	s.NoError(err)
	s.Equal(http.StatusOK, resp.StatusCode)

	responseData, err := ioutil.ReadAll(resp.Body)
	s.NoError(err)
	defer resp.Body.Close()

	response, err := ParseResponse(responseData)
	s.NoError(err)
	expectSuccess(s.Suite, response)

	// check response
	deposit := &dto.DepositResponse{}
	err = json.Unmarshal(response.Data, deposit)
	s.Require().NoError(err)
	s.NotZero(deposit.ID)
	s.NotZero(deposit.CreatedAt)
	s.Equal(amount, deposit.Amount)
	s.Equal(newWallet.ID, deposit.WalletID)
 	fromDBWallet, err := s.db.GetWallet(ctx, newWallet.ID)
	s.Require().NoError(err)
	s.Equal(fromDBWallet.Balance, amount)
}

type Response struct {
	Error   int             `json:"error"`
	Success bool            `json:"success"`
	Data    json.RawMessage `json:"data"`
}

func ParseResponse(data []byte) (*Response, error) {
	response := &Response{}
	err := json.Unmarshal(data, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func expectSuccess(s suite.Suite, response *Response) {
	if !response.Success {
		s.Fail("expected success")
	}
}
