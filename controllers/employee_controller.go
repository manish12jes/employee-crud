package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	validator "github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/manish12jes/employee-crud/error_handler"
	"github.com/manish12jes/employee-crud/helper"
	"github.com/manish12jes/employee-crud/managers"
	"github.com/manish12jes/employee-crud/models"
)

func Employees(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	offset := limit * (page - 1)

	res, total, err := managers.GetEmployees(offset, limit)
	if err != nil {
		error_handler.NotFoundError(w, "")
		return
	}
	pagination := helper.Pagination(limit, page, total)
	StatusOkResponse(w, res, pagination)
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	query := mux.Vars(r)
	id := query["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		error_handler.InvalidIdError(w, "Please retry with valid id")
	}
	res, err := managers.GetEmployeeById(idInt)
	if err != nil {
		error_handler.NotFoundError(w, "Employee not found")
		return
	}
	StatusOkResponse(w, res, nil)
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	query := mux.Vars(r)
	id := query["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		error_handler.BadRequestError(w, "Please retry with valid id")
	}

	body := json.NewDecoder(r.Body)
	var e models.EmployeeParams
	err = body.Decode(&e)
	if err != nil {
		error_handler.BadRequestError(w, "Bad request")
	}

	validate := validator.New()
	err = validate.Struct(e)
	if err != nil {
		error_handler.BadRequestError(w, "Bad request")
	}

	res, err := managers.UpdateEmployee(e, idInt)
	if err != nil {
		error_handler.BadRequestError(w, "Error in updating employee")
	}
	StatusOkResponse(w, res, nil)
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	body := json.NewDecoder(r.Body)
	var e models.EmployeeParams
	err := body.Decode(&e)
	if err != nil {
		error_handler.BadRequestError(w, "Bad request")
		return
	}
	validate := validator.New()
	err = validate.Struct(e)
	if err != nil {
		error_handler.BadRequestError(w, "Bad request: Invalid parameters")
		return
	}
	res, err := managers.CreateEmployee(e)
	if err != nil {
		error_handler.BadRequestError(w, "Error in creating employee")
		return
	}
	StatusCreatedResponse(w, res, nil)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	query := mux.Vars(r)
	id := query["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		error_handler.InvalidIdError(w, "Please retry with valid id")
	}
	err = managers.DeleteEmployeeById(idInt)
	if err != nil {
		error_handler.NotFoundError(w, err.Error())
		return
	}
	res := "employee deleted"
	StatusOkResponse(w, res, nil)
}
