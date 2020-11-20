package api

import (
	"context"
	"test-payment-system/internal/app/payment/dto"
	"test-payment-system/internal/pkg/apilog"
	"test-payment-system/internal/pkg/service"
)

// NewWallet create new wallet
// @Summary New wallet
// @Description create new wallet, return last created id
// @Tags Payment System
// @Accept json
// @Produce json
// @Param Payload body dto.NewWalletRequest true "Request Payload"
// @Success 200 {object} service.Response{data=dto.NewWalletResponse} "Success operation"
// @Router /api/v1/payment/wallet [post]
//
func (a *API) NewWallet(ctx context.Context, in service.DataObject) (response interface{}, responseErr error) {
	request := in.(*dto.NewWalletRequest)
	log := apilog.LogStart(ctx, a.log, "wallet_name", request.Name)
	defer func() { apilog.LogFinish(ctx, log, response, responseErr) }()

	//TODO: implement

	return dto.NewWalletResponse{
		ID:        0,
		Name:      request.Name,
		CreatedAt: 0,
		Balance:   0,
	}, nil
}
