package api

import (
	"fmt"
	"io"
	"strconv"
	"test-payment-system/internal/app/payment/database/model"
	"time"
)

func writeOperationsCSV(w io.Writer, operations []*model.Operation) error {
	if err := writeHeaderOperationCSV(w); err != nil {
		return err
	}
	for _, operation := range operations {
		_, err := w.Write([]byte(makeRowCSVFromOperation(operation)))
		if err != nil {
			return err
		}
	}
	return nil
}

func writeHeaderOperationCSV(w io.Writer) error {
	_, err := w.Write([]byte(`"Date","ID","Wallet ID","Wallet Name","Unit","Operation Type","Amount",` +
		`"Wallet Destination ID","Wallet Destination"` + "\r\n"))
	return err
}

func makeRowCSVFromOperation(operation *model.Operation) string {
	const layoutTimeExcel = "2006-01-02 15:04:05"
	var walletToID string
	if operation.WalletToID != nil {
		walletToID = strconv.FormatUint(uint64(*operation.WalletToID), 10)
	}
	return fmt.Sprintf("%s,%d,%d,%s,%s,%s,%f,%s,%s\r\n",
		operation.CreatedAt.Format(layoutTimeExcel),
		operation.ID, operation.WalletID, operation.WalletName,
		operation.Unit, operation.OperationType, operation.Amount,
		walletToID, operation.WalletToName.String,
	)
}

func makeFileNameCSV(unit model.Unit, timeFrom, timeTo time.Time) string {
	filename := "operation."
	if unit != "" {
		filename += string(unit) + "."
	}
	if !timeFrom.IsZero() {
		filename += timeFrom.Format("2006-01-02-150405") + "."
	}
	if !timeTo.IsZero() {
		filename += timeTo.Format("2006-01-02-150405") + "."
	}
	return filename + "csv"
}
