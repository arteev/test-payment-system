package api

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"net/http"
	"strconv"
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
// @Success 200 {object} service.Response{data=dto.WalletResponse} "Success operation"
// @Router /api/v1/payment/wallet [post]
//
func (a *API) NewWallet(ctx context.Context, in service.DataObject) (response interface{}, responseErr error) {
	request := in.(*dto.NewWalletRequest)
	log := apilog.LogStart(ctx, a.log, "wallet_name", request.Name)
	defer func() { apilog.LogFinish(ctx, log, response, responseErr) }()

	name := request.Name
	if name == "" {
		if uid, err := uuid.NewV1(); err != nil {
			log.Error(err)
			return nil, ErrSomethingWentWrong
		} else {
			name = uid.String()
		}
	}
	wallet, err := a.db.NewWallet(ctx, request.Name)
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("%w: %v", ErrFailedCreatedWallet, err)
	}

	return dto.NewWalletResponse(*wallet), nil
}

// GetWallet get wallet by id
// @Summary Get wallet
// @Description get wallet by id
// @Tags Payment System
// @Accept json
// @Produce json
// @Param id query int true "Wallet ID"
// @Success 200 {object} service.Response{data=dto.WalletResponse} "Success operation"
// @Router /api/v1/payment/wallet [get]
//
func (a *API) GetWallet(ctx context.Context, r *http.Request) (response interface{}, responseErr error) {
	log := apilog.LogStart(ctx, a.log, apilog.ValuesFromRequest(r)...)
	defer func() { apilog.LogFinish(ctx, log, response, responseErr) }()

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		return nil, fmt.Errorf("%w: id", ErrInvalidParameter)
	}

	wallet, err := a.db.GetWallet(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	return dto.NewWalletResponse(*wallet), nil
}

// Deposit deposit to wallet
// @Summary Deposit to wallet
// @Description Transfer funds to the wallet, return the deposit record
// @Tags Payment System
// @Accept json
// @Produce json
// @Param Payload body dto.DepositRequest true "Request Payload"
// @Success 200 {object} service.Response{data=dto.DepositResponse} "Success operation"
// @Router /api/v1/payment/wallet/deposit [post]
//
func (a *API) Deposit(ctx context.Context, in service.DataObject) (response interface{}, responseErr error) {
	request := in.(*dto.DepositRequest)
	log := apilog.LogStart(ctx, a.log, "wallet_id", request.WalletID,
		"amount", request.Amount)
	defer func() { apilog.LogFinish(ctx, log, response, responseErr) }()
	wallet, err := a.db.GetWallet(ctx, request.WalletID)
	if err != nil {
		return nil, err
	}
	deposit, err := a.db.Deposit(ctx, wallet.ID, request.Amount)
	if err != nil {
		return nil, err
	}
	return dto.NewDepositResponse(*deposit), nil
}

// Transfer transfer money
// @Summary transfer money
// @Description Transferring money between wallets
// @Tags Payment System
// @Accept json
// @Produce json
// @Param Payload body dto.TransferRequest true "Request Payload"
// @Success 200 {object} service.Response{data=dto.WalletTransfer} "Success operation"
// @Router /api/v1/payment/wallet/deposit [post]
//
func (a *API) Transfer(ctx context.Context, in service.DataObject) (response interface{}, responseErr error) {
	request := in.(*dto.TransferRequest)
	log := apilog.LogStart(ctx, a.log, "wallet_from", request.WalletFrom,
		"wallet_to", request.WalletTo, "amount", request.Amount)
	defer func() { apilog.LogFinish(ctx, log, response, responseErr) }()
	walletFrom, err := a.db.GetWallet(ctx, request.WalletFrom)
	if err != nil {
		return nil, err
	}
	walletTo, err := a.db.GetWallet(ctx, request.WalletTo)
	if err != nil {
		return nil, err
	}

	transfer, err := a.db.Transfer(ctx, walletFrom.ID, walletTo.ID, request.Amount)
	if err != nil {
		return nil, err
	}

	return dto.NewTransferResponse(*transfer), nil
}
