package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/v-smrnv/go-ad-api/internal/handlers"
	"github.com/v-smrnv/go-ad-api/internal/models"
	"github.com/v-smrnv/go-ad-api/internal/services"
)

func EmployeeRouteRegister(r *gin.Engine) {
	employee := models.Employee{
		Name:   "ssds",
		Branch: "ss",
	}

	router := gin.Default()
	api := router.Group("api")
	s := services.NewEmployeeService(employee)

	handler := handlers.NewEmployeeHandler(s)

	api.GET("/:id/employee", handler.GetEmployee)

}
