package api

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"net/http"
	"strconv"
	"test-payment-system/internal/app/payment/database/model"
	"test-payment-system/internal/app/payment/dto"
	"test-payment-system/internal/pkg/apilog"
	"test-payment-system/internal/pkg/service"
	"time"
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
// @Router /api/v1/payment/deposit [post]
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
// @Success 200 {object} service.Response{data=dto.TransferResponse} "Success operation"
// @Router /api/v1/payment/deposit [post]
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

// Operations wallet operations report
// @Summary Wallet operations report
// @Description report on transactions with a wallet filtered by date and type of operation
// @Tags Payment System
// @Accept json
// @Produce json
// @Param oper query string false "Operation" Enums(deposit,withdraw)
// @Param date_from int64 false "Operation date filter, end of period"
// @Param date_to int64 false "Operation date filter, period start"
// @Success 200 {object} service.Response{data=dto.WalletResponse} "Success operation"
// @Router /api/v1/payment/operations [get]
//

func (a *API) Operations(w http.ResponseWriter, r *http.Request) {
	err := a.operations(w, r)
	if err != nil {
		service.RespondJSON(w, r, nil, err)
	}
}

func (a *API) operations(w http.ResponseWriter, r *http.Request) (responseErr error) {
	ctx := r.Context()
	log := apilog.LogStart(ctx, a.log, apilog.ValuesFromRequest(r)...)
	defer func() { apilog.LogFinish(ctx, log, nil, responseErr) }()

	walletID, err := strconv.Atoi(r.FormValue("wallet_id"))
	if err != nil {
		return fmt.Errorf("%w: wallet_id", ErrInvalidParameter)
	}

	operationName := model.Unit(r.FormValue("oper"))
	var operationSign *model.OperationSign
	switch operationName {
	default:
		return fmt.Errorf("%w: oper", ErrInvalidParameter)
	case "":
		// nothing = nil
	case "deposit":
		current := model.OperationSignIncome
		operationSign = &current
	case "withdraw":
		current := model.OperationSignTransfer
		operationSign = &current
	}

	dateFrom, err := parseTimeFromUnixString(r.FormValue("date_from"))
	if err != nil {
		return fmt.Errorf("%w: date_from", ErrInvalidParameter)
	}
	dateTo, err := parseTimeFromUnixString(r.FormValue("date_to"))
	if err != nil {
		return fmt.Errorf("%w: date_to", ErrInvalidParameter)
	}

	w.Header().Add("Content-Type", "text/csv")
	w.Header().Add("Content-Disposition", "attachment;filename="+makeFileNameCSV(operationName, dateFrom, dateTo))
	w.WriteHeader(http.StatusOK)
	operations, err := a.db.OperationWallet(ctx, uint(walletID), operationSign, dateFrom, dateTo)
	if err != nil {
		return err
	}
	if err = writeOperationsCSV(w, operations); err != nil {
		log.Error(err)
		return ErrSomethingWentWrong
	}

	return nil
}

func parseTimeFromUnixString(unixTimeString string) (time.Time, error) {
	if unixTimeString == "" {
		return time.Time{}, nil
	}
	unixTime, err := strconv.ParseInt(unixTimeString, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(unixTime, 0), nil
}
