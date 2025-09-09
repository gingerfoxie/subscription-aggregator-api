package handlers

import (
	"net/http"
	"strconv"
	"time"

	model "subscription-service/internal/models"
	"subscription-service/internal/service"
	"subscription-service/pkg/logger"

	"github.com/gin-gonic/gin"
)

type SubscriptionHandler struct {
	service service.SubscriptionService
}

func NewSubscriptionHandler(s service.SubscriptionService) *SubscriptionHandler {
	return &SubscriptionHandler{service: s}
}

// @Summary Создать подписку
// @Description Создает новую запись о подписке
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param subscription body model.Subscription true "Данные подписки"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /subscriptions [post]
func (h *SubscriptionHandler) Create(c *gin.Context) {
	var sub model.Subscription
	if err := c.ShouldBindJSON(&sub); err != nil {
		logger.Error("Invalid JSON: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Валидация формата даты
	if !isValidMonthYear(sub.StartDate) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "start_date must be in MM-YYYY format"})
		return
	}
	if sub.EndDate != nil && !isValidMonthYear(*sub.EndDate) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "end_date must be in MM-YYYY format"})
		return
	}

	if err := h.service.Create(&sub); err != nil {
		logger.Error("Failed to create subscription: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create subscription"})
		return
	}

	logger.Info("Created subscription: ", sub.ID)
	c.JSON(http.StatusCreated, gin.H{"message": "Subscription created", "id": sub.ID})
}

// @Summary Получить подписку по ID
// @Tags subscriptions
// @Produce json
// @Param id path int true "ID подписки"
// @Success 200 {object} model.Subscription
// @Failure 404 {object} map[string]string
// @Router /subscriptions/{id} [get]
func (h *SubscriptionHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	sub, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subscription not found"})
		return
	}

	c.JSON(http.StatusOK, sub)
}

// @Summary Обновить подписку
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param id path int true "ID подписки"
// @Param subscription body model.Subscription true "Данные для обновления"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /subscriptions/{id} [put]
func (h *SubscriptionHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var sub model.Subscription
	if err := c.ShouldBindJSON(&sub); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	sub.ID = uint(id)

	if err := h.service.Update(&sub); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subscription not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated"})
}

// @Summary Удалить подписку
// @Tags subscriptions
// @Param id path int true "ID подписки"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /subscriptions/{id} [delete]
func (h *SubscriptionHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subscription not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}

// @Summary Список подписок
// @Tags subscriptions
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Param user_id query string false "User ID"
// @Param service_name query string false "Service name"
// @Success 200 {object} map[string]interface{}
// @Router /subscriptions [get]
func (h *SubscriptionHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	filters := make(map[string]interface{})
	if userID := c.Query("user_id"); userID != "" {
		filters["user_id"] = userID
	}
	if serviceName := c.Query("service_name"); serviceName != "" {
		filters["service_name"] = serviceName
	}

	subs, total, err := h.service.List(filters, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch subscriptions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":        subs,
		"total":       total,
		"page":        page,
		"page_size":   limit,
		"total_pages": (total + int64(limit) - 1) / int64(limit),
	})
}

// @Summary Общая стоимость подписок за период
// @Tags analytics
// @Produce json
// @Param user_id query string false "User ID"
// @Param service_name query string false "Service name"
// @Param start query string true "Start period (MM-YYYY)"
// @Param end query string true "End period (MM-YYYY)"
// @Success 200 {object} map[string]int
// @Failure 400 {object} map[string]string
// @Router /total [get]
func (h *SubscriptionHandler) GetTotalCost(c *gin.Context) {
	userID := c.Query("user_id")
	serviceName := c.Query("service_name")
	start := c.Query("start")
	end := c.Query("end")

	if start == "" || end == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "start and end are required"})
		return
	}

	if !isValidMonthYear(start) || !isValidMonthYear(end) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "start and end must be in MM-YYYY format"})
		return
	}

	total, err := h.service.GetTotalCost(userID, serviceName, start, end)
	if err != nil {
		logger.Error("Error calculating total cost: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate cost"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"total_cost": total})
}

func isValidMonthYear(s string) bool {
	_, err := time.Parse("01-2006", s)
	return err == nil
}
