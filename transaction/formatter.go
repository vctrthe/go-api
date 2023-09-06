package transaction

import "time"

type CampaignTransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type UserTransactionFormatter struct {
	ID        int                    `json:"id"`
	Amount    int                    `json:"amount"`
	Status    string                 `json:"status"`
	CreatedAt time.Time              `json:"created_at"`
	Campaign  UserCampTransFormatter `json:"campaign"`
}

type UserCampTransFormatter struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
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

func FormatUserTransaction(transaction Transaction) UserTransactionFormatter {
	formatter := UserTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.CreatedAt = transaction.CreatedAt
	campaignFormat := UserCampTransFormatter{}
	campaignFormat.Name = transaction.Campaign.Name
	campaignFormat.ImageUrl = ""
	if len(transaction.Campaign.CampaignImages) > 0 {
		campaignFormat.ImageUrl = transaction.Campaign.CampaignImages[0].FileName
	}
	formatter.Campaign = campaignFormat
	return formatter
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter {
	if len(transactions) == 0 {
		return []UserTransactionFormatter{}
	}
	var transFormatter []UserTransactionFormatter
	for _, transaction := range transactions {
		formatter := FormatUserTransaction(transaction)
		transFormatter = append(transFormatter, formatter)
	}
	return transFormatter
}
