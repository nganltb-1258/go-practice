package services

import (
    "employee/dto"
    "employee/models"
    "employee/repositories"
    "net/http"
    "log"
    "strconv"
)

func FetchEmployees(r *http.Request) ([]dto.EmployeeResponse, error, string, int, int, int, int) {
    // Get condition from param
    keyword := r.URL.Query().Get("keyword")
    page := r.URL.Query().Get("page")
    pageSize := r.URL.Query().Get("page_size")

    if pageSize == "" {
        pageSize = "10"
    }
    if page == "" {
        page = "1"
    }

    pageInt, _ := strconv.Atoi(page)
    pageSizeInt, _ := strconv.Atoi(pageSize)

    // Get total count to calculate total pages
    total, err := repositories.CountEmployees(keyword)
    if err != nil {
        log.Printf("Error counting employees: %v", err)
        return nil, err, keyword, pageInt, pageSizeInt, 0, 0
    }
    totalPages := (total + pageSizeInt - 1) / pageSizeInt
    if pageInt > totalPages && totalPages != 0 {
        pageInt = totalPages
    }

    var employees []dto.EmployeeResponse
    if total == 0 {
        return employees, nil, keyword, pageInt, pageSizeInt, total, totalPages
    }

    var offset int
    offset = (pageInt - 1) * pageSizeInt
    employees, err = repositories.GetEmployeesByCondition(keyword, offset, pageSizeInt)
    if err != nil {
        log.Printf("Error fetching employees: %v", err)
        return nil, err, keyword, pageInt, pageSizeInt, total, totalPages
    }

    return employees, nil, keyword, pageInt, pageSizeInt, total, totalPages
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
