package services

import (
    "employee/dto"
    "employee/models"
    "employee/repositories"
    "net/http"
    "log"
    "strconv"
)

func FetchAllEmployees(keyword string) ([]dto.EmployeeResponse, error) {
    employees, err := repositories.GetAllEmployees(keyword)

    if err != nil {
        log.Printf("Error fetching employees: %v", err)
        return nil, err
    }
    return employees, nil
}

func InsertEmployee(r *http.Request) error {
    ageStr := r.FormValue("age")
    salaryStr := r.FormValue("salary")
    departmentStr := r.FormValue("department")
    age, _ := strconv.Atoi(ageStr)
    salary, _ := strconv.ParseFloat(salaryStr, 64)
    department, _ := strconv.Atoi(departmentStr)

    employee := models.Employee{
        Name: r.FormValue("name"),
        Age: age,
        Position: r.FormValue("position"),
        DepartmentId: department,
        Salary: salary,
    }

    err := repositories.InsertEmployee(employee)
    if err != nil {
        log.Printf("Error creating employee: %v", err)
        return err
    }
    return nil
}

func DeleteEmployee(r *http.Request) error {
    idStr := r.FormValue("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        log.Printf("Invalid employee ID: %v", err)
        return err
    }

    err = repositories.DeleteEmployeeById(id)
    if err != nil {
        log.Printf("Error deleting employee: %v", err)
        return err
    }
    return nil
}

func FetchEmployeeById(idStr string) (models.Employee, error) {
    id, err := strconv.Atoi(idStr)
    if err != nil {
        log.Printf("Invalid employee ID: %v", err)
        return models.Employee{}, err
    }

    employee, err := repositories.GetEmployeeById(id)
    if err != nil {
        log.Printf("Error fetching employee by ID: %v", err)
        return models.Employee{}, err
    }
    return employee, nil
}

func UpdateEmployee(r *http.Request) error {
    idStr := r.FormValue("id")
    ageStr := r.FormValue("age")
    salaryStr := r.FormValue("salary")
    departmentStr := r.FormValue("department")

    id, _ := strconv.Atoi(idStr)
    age, _ := strconv.Atoi(ageStr)
    salary, _ := strconv.ParseFloat(salaryStr, 64)
    department, _ := strconv.Atoi(departmentStr)

    employee := models.Employee{
        Id: id,
        Name: r.FormValue("name"),
        Age: age,
        Position: r.FormValue("position"),
        DepartmentId: department,
        Salary: salary,
    }

    err := repositories.UpdateEmployee(employee)
    if err != nil {
        log.Printf("Error updating employee: %v", err)
        return err
    }
    return nil
}
