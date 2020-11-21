package dto

import (
	"fmt"
	"test-payment-system/internal/app/payment/database/model"
	"test-payment-system/internal/pkg/service"
)

// NewWalletRequest struct of request for new wallet
type NewWalletRequest struct {
	Name string `json:"name"`
}

// Validate check if necessary fields are not empty
func (r *NewWalletRequest) Validate() error {
	return nil
}

// NewWalletRequestMeta meta type for struct of new wallet request
type NewWalletRequestMeta struct{}

// New return ptr for struct of new wallet request
func (m *NewWalletRequestMeta) New() service.DataObject {
	return &NewWalletRequest{}
}

// WalletResponse dto to return when new wallet was created or getting
type WalletResponse struct {
	ID        uint    `json:"id"`
	Name      string  `json:"Name"`
	CreatedAt int64   `json:"created_at"`
	UpdatedAt int64   `json:"updated_at"`
	Balance   float64 `json:"balance"`
}

// NewWalletResponse create WalletResponse from wallet model
func NewWalletResponse(wallet model.Wallet) *WalletResponse {
	return &WalletResponse{
		ID:        wallet.ID,
		Name:      wallet.Name,
		CreatedAt: wallet.CreatedAt.Unix(),
		UpdatedAt: wallet.UpdatedAt.Unix(),
		Balance:   wallet.Balance,
	}
}

// DepositRequest struct of request for deposit
type DepositRequest struct {
	WalletID uint    `json:"wallet_id"`
	Amount   float64 `json:"amount"`
}

// Validate check if necessary fields are not empty
func (r *DepositRequest) Validate() error {
	if r.Amount <= 0 {
		return fmt.Errorf("invalid parameter: amount of the deposit must be positive")
	}
	if r.WalletID <= 0 {
		return fmt.Errorf("invalid parameter: wallet id required")
	}
	return nil
}

// DepositRequestMeta meta type for struct of deposit wallet request
type DepositRequestMeta struct{}

// New return ptr for struct of new wallet request
func (m *DepositRequestMeta) New() service.DataObject {
	return &DepositRequest{}
}

// DepositResponse dto return with a new deposit to the wallet
type DepositResponse struct {
	ID        uint    `json:"id"`
	CreatedAt int64   `json:"created_at"`
	WalletID  uint    `json:"wallet_id"`
	Amount    float64 `json:"amount"`
}

// NewDepositResponse create and returns DepositResponse from deposit model
func NewDepositResponse(deposit model.WalletDeposit) *DepositResponse {
	return &DepositResponse{
		ID:        deposit.ID,
		CreatedAt: deposit.CreatedAt.Unix(),
		WalletID:  deposit.WalletID,
		Amount:    deposit.Amount,
	}
}

// TransferRequestMeta meta type for struct of transfer money request
type TransferRequestMeta struct{}

// New return ptr for struct of new wallet request
func (m *TransferRequestMeta) New() service.DataObject {
	return &TransferRequest{}
}

// TransferRequest dto return with a transfer money
type TransferRequest struct {
	WalletFrom uint    `json:"wallet_from"`
	WalletTo   uint    `json:"wallet_to"`
	Amount     float64 `json:"amount"`
}

// Validate check if necessary fields are not empty
func (r *TransferRequest) Validate() error {
	if r.Amount <= 0 {
		return fmt.Errorf("invalid parameter: amount of the deposit must be positive")
	}
	if r.WalletFrom <= 0 {
		return fmt.Errorf("invalid parameter: wallet_from required")
	}
	if r.WalletTo <= 0 {
		return fmt.Errorf("invalid parameter: wallet_to required")
	}
	if r.WalletFrom == r.WalletTo {
		return fmt.Errorf("source wallet must not match destination wallet")
	}
	return nil
}

// TransferResponse dto return with a new transfer money
type TransferResponse struct {
	ID         uint    `json:"id"`
	WalletFrom uint    `json:"wallet_from"`
	WalletTo   uint    `json:"wallet_from"`
	Amount     float64 `json:"amount"`
	CreatedAt  int64   `json:"created_at"`
}

// TransferResponse create and returns TransferResponse from transfer money model
func NewTransferResponse(transfer model.WalletTransfer) *TransferResponse {
	return &TransferResponse{
		ID:         transfer.ID,
		WalletFrom: transfer.WalletFrom,
		WalletTo:   transfer.WalletTo,
		Amount:     transfer.Amount,
		CreatedAt:  transfer.CreatedAt.Unix(),
	}
}
