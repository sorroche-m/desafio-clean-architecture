package http

import (
	"net/http"

	"github.com/sorroche-m/desafio-clean-architecture/internal/domain"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderUseCase domain.OrderUseCase
}

func NewOrderHandler(orderUseCase domain.OrderUseCase) *OrderHandler {
	return &OrderHandler{
		orderUseCase: orderUseCase,
	}
}

type CreateOrderRequest struct {
	CustomerID string  `json:"customer_id" binding:"required"`
	Amount     float64 `json:"amount" binding:"required"`
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.orderUseCase.CreateOrder(req.CustomerID, req.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}

func (h *OrderHandler) ListOrders(c *gin.Context) {
	orders, err := h.orderUseCase.ListOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) GetOrder(c *gin.Context) {
	id := c.Param("id")

	order, err := h.orderUseCase.GetOrder(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}
