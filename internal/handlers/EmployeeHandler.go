package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/v-smrnv/go-ad-api/internal/models"
	"github.com/v-smrnv/go-ad-api/internal/services"
)

type EmployeeHandler struct {
	service services.EmployeeInterface
}

func NewEmployeeHandler(s services.EmployeeInterface) *EmployeeHandler {
	return &EmployeeHandler{service: s}
}

func (h *EmployeeHandler) GetEmployeeInfo(id string) (models.Employee, error) {
	return h.service.GetEmployeeInfo(id)
}

func (h *EmployeeHandler) GetEmployee(c *gin.Context) {
	id := c.Param("id")
	empl, err := h.service.GetEmployeeInfo(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, empl)

}
