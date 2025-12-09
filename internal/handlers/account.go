package handlers

import (
	"account-management/internal/repo"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountsHandler struct{ db *repo.MgmtDB }

func NewAccountsHandler(db *repo.MgmtDB) *AccountsHandler { return &AccountsHandler{db: db} }

type CreateReq struct {
	Email     string `json:"email" binding:"required,email"`
	FullName  string `json:"full_name" binding:"required"`
	BirthDate string `json:"birth_date" binding:"required"`
	Country   string `json:"country" binding:"required"`
}

func (h *AccountsHandler) Create(c *gin.Context) {
	var req CreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	a := repo.Account{
		Email: req.Email, FullName: req.FullName, BirthDate: req.BirthDate, Country: req.Country,
	}
	out, err := h.db.CreateAccount(c, a)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "create failed"})
		return
	}
	c.JSON(http.StatusCreated, out)
}

func (h *AccountsHandler) Get(c *gin.Context) {
	id := c.Param("account_id")
	a, err := h.db.GetAccount(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, a)
}

func (h *AccountsHandler) List(c *gin.Context) {
	items, err := h.db.ListAccounts(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "list failed"})
		return
	}
	c.JSON(http.StatusOK, items)
}
