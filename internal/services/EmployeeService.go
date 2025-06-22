package services

import "github.com/v-smrnv/go-ad-api/internal/models"

type EmployeeService struct {
	employee models.Employee
}

type EmployeeInterface interface {
	GetEmployeeInfo(id string) (models.Employee, error)
}

func NewEmployeeService(e models.Employee) *EmployeeService {
	return &EmployeeService{employee: e}
}

// Dummy for learning
func (s *EmployeeService) GetEmployeeInfo(id string) (models.Employee, error) {

	return s.employee, nil
}

type EmployeeADService struct {
	host string
}

// TODO:	add implementation for Active Directory
func (s *EmployeeADService) GetEmployeeInfo(id string) (models.Employee, error) {

	adHost := s.host
	employee := models.Employee{
		Name:   adHost,
		Branch: "ss",
	}
	return employee, nil
}
