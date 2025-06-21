package services

import "github.com/v-smrnv/go-ad-api/internal/models"

type EmployeeService struct{
	employee models.Employee
}

func NewEmployeeService(e models.Employee) *EmployeeService{
	return &EmployeeService{employee: e}
}

func (s *EmployeeService) GetEmployeeInfo(id string) (models.Employee, error) {

	return s.employee, nil
}
