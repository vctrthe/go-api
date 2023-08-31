package transaction

import "go-api/user"

type GetTransByCampaignInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}
