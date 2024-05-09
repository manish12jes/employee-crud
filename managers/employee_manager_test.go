package managers

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/manish12jes/employee-crud/models"
)

func TestGetEmployeeById(t *testing.T) {
	empId := 1
	test1 := struct {
		name string
		args int
		want int
	}{
		name: "should return employe with id 1",
		args: empId,
		want: empId,
	}

	t.Run(test1.name, func(t *testing.T) {
		got, err := func() (*models.Employee, error) {
			employee, err := GetEmployeeById(test1.args)
			if err != nil {
				return nil, err
			}
			return employee, nil
		}()
		if err == nil {
			assert.Equal(t, test1.want, got.ID)
		}
	})

}
