package service

import "github.com/v-smrnv/go-ad-api/internal/models"

type EmployeeService struct{}

func GetEmployeeInfo(id string) (models.Employee, error) {
	employee := models.Employee{
		Name:   "Ivan",
		Branch: "DO MOSKVA",
	}
	return employee, nil
}
