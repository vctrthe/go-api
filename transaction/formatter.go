package transaction

import "time"

type CampaignTransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatCampTransaction(transaction Transaction) CampaignTransactionFormatter {
	transactionFormatter := CampaignTransactionFormatter{}
	transactionFormatter.ID = transaction.ID
	transactionFormatter.Name = transaction.User.Name
	transactionFormatter.Amount = transaction.Amount
	transactionFormatter.CreatedAt = transaction.CreatedAt
	return transactionFormatter
}

func FormatCampTransactions(transactions []Transaction) []CampaignTransactionFormatter {
	if len(transactions) == 0 {
		return []CampaignTransactionFormatter{}
	}
	var transFormatter []CampaignTransactionFormatter
	for _, transaction := range transactions {
		formatter := FormatCampTransaction(transaction)
		transFormatter = append(transFormatter, formatter)
	}
	return transFormatter
}
