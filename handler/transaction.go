package handler

import (
	"go-api/helper"
	"go-api/transaction"
	"go-api/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) GetCampaignTransactions(c *gin.Context) {
	var input transaction.GetTransByCampaignInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("Error getting campaign's transactions!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser
	transactions, err := h.service.GetTransByCampID(input)
	if err != nil {
		response := helper.ApiResponse("Failed to get campaign's transactions!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Campaign's Transactions detail", http.StatusOK, "success", transaction.FormatCampTransactions(transactions))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetUserTransaction(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	transactions, err := h.service.GetTransByUserID(userID)
	if err != nil {
		response := helper.ApiResponse("Failed to get user's transactions!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("User's Transactions detail", http.StatusOK, "success", transaction.FormatUserTransactions(transactions))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) CreateTransaction(c *gin.Context) {
	var input transaction.CreateTransactionInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ValidationErrResponse(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Failed to create transaction!", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newTransaction, err := h.service.CreateTransaction(input)
	if err != nil {
		response := helper.ApiResponse("failed to create transaction!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Create transaction success!", http.StatusOK, "success", newTransaction)
	c.JSON(http.StatusOK, response)
}
