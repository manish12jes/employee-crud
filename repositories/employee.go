package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/manish12jes/employee-crud/models"
)

type Employee models.Employee

const db_file = "DB/employee_db.json"

func GetEmployees(offset, limit int) ([]models.Employee, int, error) {
	// _sql := fmt.Sprintf("SELECT id, name, position, salary FROM employees order by id desc limit %d offset %d", limit, offset)
	filteredEmp := []models.Employee{}
	employees, err := dbEmployee()
	total := len(employees)
	if err != nil {
		return nil, 0, errors.New("employee not found")
	}
	for i := offset; i <= offset+limit; i++ {
		if emp, ok := employees[i]; ok {
			filteredEmp = append(filteredEmp, emp)
		}
	}
	return filteredEmp, total, nil
}

func GetEmployeeById(id int) (*models.Employee, error) {
	employees, err := dbEmployee()
	if err != nil {
		return nil, errors.New("employee not found")
	}
	if emp, ok := employees[id]; ok {
		return &emp, nil
	}
	return nil, errors.New("employee not found")
}

func GetAllEmployeesDb() (map[int]models.Employee, error) {
	return dbEmployee()
}

func dbEmployee() (map[int]models.Employee, error) {
	employees := map[int]models.Employee{}
	data, err := os.ReadFile(db_file)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(data, &employees)
	return employees, nil
}

func Save(e models.EmployeeParams) (*models.Employee, error) {
	newEmp := models.Employee{
		Name:     e.Name,
		Position: e.Position,
		Salary:   e.Salary,
	}
	return Employee(newEmp).save()
}

func Update(id int, e models.EmployeeParams) (*models.Employee, error) {
	employee, err := GetEmployeeById(id)
	if err != nil {
		return employee, err
	}
	employee.Name = e.Name
	employee.Salary = e.Salary
	employee.Position = e.Position

	return Employee(*employee).save()
}

func Delete(id int) error {
	employees, err := GetAllEmployeesDb()
	if len(employees) == 0 || err != nil {
		return errors.New("employee not found")
	}

	if _, ok := employees[id]; !ok {
		return errors.New("employee not found")
	}
	delete(employees, id)
	data, _ := json.MarshalIndent(employees, "", " ")
	_ = os.WriteFile(db_file, data, 0644)
	return nil
}

func (employee Employee) save() (*models.Employee, error) {
	employees, _ := dbEmployee()

	//New employee
	if employee.ID == 0 {
		employee.ID = len(employees) + 1
	}

	emp := models.Employee(employee)
	employees[employee.ID] = emp
	data, _ := json.MarshalIndent(employees, "", " ")
	_ = os.WriteFile(db_file, data, 0644)

	return &emp, nil
}
