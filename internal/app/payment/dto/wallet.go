package dto

import "test-payment-system/internal/pkg/service"

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

// NewWalletResponse dto to return when new wallet was created
type NewWalletResponse struct {
	ID        int     `json:"id"`
	Name      string  `json:"Name"`
	CreatedAt int64   `json:"created_at"`
	Balance   float64 `json:"balance"`
}
