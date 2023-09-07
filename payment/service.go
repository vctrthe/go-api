package payment

import (
	"go-api/config"
	"go-api/transaction"
	"go-api/user"
	"strconv"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type Service interface {
	GetToken(transaction transaction.Transaction, user user.User) (string, error)
}

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) GetToken(transaction transaction.Transaction, user user.User) (string, error) {
	midtrans.ServerKey = config.MidtransServerKey
	midtrans.ClientKey = config.MidtransClientKey
	midtrans.Environment = midtrans.Sandbox

	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			Email: user.Email,
			FName: user.Name,
		},
	}

	snapTrans, err := snap.CreateTransaction(snapReq)
	if err != nil {
		return "", err
	}

	redirectURL := snapTrans.RedirectURL

	return redirectURL, nil
}
