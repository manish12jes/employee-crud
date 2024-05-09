package managers

import (
	"github.com/manish12jes/employee-crud/models"
	"github.com/manish12jes/employee-crud/repositories"
)

// the managers will be used to implement all business logic

func GetEmployees(offset, limit int) ([]models.Employee, int, error) {
	employess, total, err := repositories.GetEmployees(offset, limit)
	if err != nil {
		return employess, 0, err
	}
	return employess, total, nil
}

func GetEmployeeById(id int) (*models.Employee, error) {
	employee, err := repositories.GetEmployeeById(id)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

func CreateEmployee(e models.EmployeeParams) (*models.Employee, error) {
	return repositories.Save(e)
}

func UpdateEmployee(e models.EmployeeParams, id int) (*models.Employee, error) {
	return repositories.Update(id, e)
}

func DeleteEmployeeById(id int) error {
	return repositories.Delete(id)
}
