package handlers

import (
    "employee/services"
    "employee/dto"
    "employee/models"
    "html/template"
    "net/http"
    "log"
)

type EmployeePageData struct {
    Employees []dto.EmployeeResponse
    Keyword string
    Page int
    PageSize int
    TotalPages int
    Total int
    Pages []int
}

type DepartmentPageData struct {
    Departments []models.Department
}

func EmployeeIndexHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.New("index").Funcs(template.FuncMap{
        "add": func(a, b int) int {
            return a + b
        },
        "sub": func(a, b int) int {
            return a - b
        },
    }).ParseFiles("templates/index.html"))

    employees, err, keyword, page, pageSize, total, totalPages := services.FetchEmployees(r)
    if err != nil {
        http.Error(w, "Error fetching employees", http.StatusInternalServerError)
        return
    }

    pages := make([]int, totalPages)
    for i := 0; i < totalPages; i++ {
        pages[i] = i + 1
    }

    data := EmployeePageData{
        Employees: employees,
        Keyword: keyword,
        Page: page,
        PageSize: pageSize,
        TotalPages: totalPages,
        Total: total,
        Pages: pages,
    }

    err = tmpl.Execute(w, data)
    if err != nil {
        log.Printf("Error executing template: %v", err)
        http.Error(w, "Error rendering page", http.StatusInternalServerError)
    }
}

func EmployeeCreateHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.New("create").ParseFiles("templates/create.html"))

    departments, err := services.FetchAllDepartments()
    if err != nil {
        http.Error(w, "Error fetching departments", http.StatusInternalServerError)
        return
    }

    data := DepartmentPageData{
        Departments: departments,
    }

    err = tmpl.Execute(w, data)
    if err != nil {
        log.Printf("Error executing template: %v", err)
        http.Error(w, "Error rendering page", http.StatusInternalServerError)
    }
}

func EmployeeInsertHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    err := services.InsertEmployee(r)
    if err != nil {
        log.Printf("Error inserting employee: %v", err)
        http.Error(w, "Error creating employee", http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/employees", http.StatusSeeOther)
}


func EmployeeDeleteHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    err := services.DeleteEmployee(r)
    if err != nil {
        log.Printf("Error deleting employee: %v", err)
        http.Error(w, "Error deleting employee", http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/employees", http.StatusSeeOther)
}

func EmployeeEditHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.New("edit").ParseFiles("templates/edit.html"))

    employeeId := r.URL.Query().Get("id")
    employee, err := services.FetchEmployeeById(employeeId)
    if err != nil {
        http.Error(w, "Error fetching employee", http.StatusInternalServerError)
        return
    }

    departments, err := services.FetchAllDepartments()
    if err != nil {
        http.Error(w, "Error fetching departments", http.StatusInternalServerError)
        return
    }

    data := struct {
        Employee models.Employee
        Departments []models.Department
    }{
        Employee: employee,
        Departments: departments,
    }

    err = tmpl.Execute(w, data)
    if err != nil {
        log.Printf("Error executing template: %v", err)
        http.Error(w, "Error rendering page", http.StatusInternalServerError)
    }
}

func EmployeeUpdateHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    err := services.UpdateEmployee(r)
    if err != nil {
        log.Printf("Error updating employee: %v", err)
        http.Error(w, "Error updating employee", http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/employees", http.StatusSeeOther)
}
