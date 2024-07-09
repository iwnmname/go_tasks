package accounts

import (
	"awesomeProject/accounts/dto"
	"awesomeProject/accounts/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"sync"
)

func New() *Handler {
	return &Handler{
		accounts: make(map[string]*models.Account),
		guard:    &sync.RWMutex{},
	}
}

type Handler struct {
	accounts map[string]*models.Account
	guard    *sync.RWMutex
}

func (h *Handler) CreateAccount(c echo.Context) error {
	var request dto.ChangeAccountRequest // {"name": "alice", "amount": 50}
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}

	h.guard.Lock()

	if _, ok := h.accounts[request.Name]; ok {
		h.guard.Unlock()

		return c.String(http.StatusForbidden, "account already exists")
	}

	h.accounts[request.Name] = &models.Account{
		Name:   request.Name,
		Amount: request.Amount,
	}

	h.guard.Unlock()

	return c.NoContent(http.StatusCreated)
}

func (h *Handler) GetAccount(c echo.Context) error {
	name := c.QueryParam("name")

	h.guard.RLock()
	defer h.guard.RUnlock()

	if account, ok := h.accounts[name]; ok {
		fmt.Printf("GET: Account %s: amount %d\n", name, account.Amount)
		return c.JSON(http.StatusOK, dto.GetAccountResponse{
			Name:   name,
			Amount: account.Amount,
		})
	}

	return c.String(http.StatusNotFound, "account not found")
}

func (h *Handler) DeleteAccount(c echo.Context) error {
	name := c.QueryParams().Get("name")
	h.guard.Lock()
	if _, ok := h.accounts[name]; !ok {
		h.guard.Unlock()
		return c.String(http.StatusNotFound, "account not found")
	}
	delete(h.accounts, name)
	h.guard.Unlock()
	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) PathAccount(c echo.Context) error {
	name := c.QueryParam("name")
	amountStr := c.QueryParam("amount")
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid amount")
	}

	h.guard.Lock()
	defer h.guard.Unlock()

	account, ok := h.accounts[name]
	if !ok {
		return c.String(http.StatusNotFound, "account not found")
	}

	oldAmount := account.Amount
	account.Amount = amount
	h.accounts[name] = account

	fmt.Printf("PATCH: Account %s: old amount %d, adding %d, new amount %d\n",
		name, oldAmount, amount, account.Amount)

	response := map[string]interface{}{
		"name":   name,
		"amount": account.Amount,
	}
	fmt.Printf("Sending response: %+v\n", response)
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) ChangeAccount(c echo.Context) error {
	name := c.QueryParam("name")
	newName := c.QueryParam("new_name")

	h.guard.Lock()
	defer h.guard.Unlock()

	if account, ok := h.accounts[name]; ok {
		h.accounts[newName] = account
		delete(h.accounts, name)
		return c.NoContent(http.StatusNoContent)
	}

	return c.String(http.StatusNotFound, "account not found")
}
