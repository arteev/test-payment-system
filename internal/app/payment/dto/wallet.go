package dto

import (
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
