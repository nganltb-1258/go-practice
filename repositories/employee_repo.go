package repositories

import (
    "database/sql"
    "employee/config"
    "employee/models"
    "employee/dto"
    "log"
)

func GetAllEmployees(keyword string) ([]dto.EmployeeResponse, error) {
    query := `SELECT e.id, e.name, e.age, e.position, e.salary, d.name
        FROM employees e
        JOIN departments d ON e.department_id = d.id`
    if keyword != "" {
        query += ` WHERE e.name ILIKE '%' || $1 || '%'`
    }

    query += ` ORDER BY e.id;`

    var rows *sql.Rows
    var err error
    if keyword != "" {
        rows, err = config.DB.Query(query, keyword)
    } else {
        rows, err = config.DB.Query(query)
    }

    if err != nil {
        log.Printf("Error querying employees: %v", err)
        return nil, err
    }
    defer rows.Close()

    var employees []dto.EmployeeResponse
    for rows.Next() {
        var emp dto.EmployeeResponse
        if err := rows.Scan(&emp.Id, &emp.Name, &emp.Age, &emp.Position, &emp.Salary, &emp.DepartmentName); err != nil {
            log.Printf("Error scanning employee row: %v", err)
            return nil, err
        }
        employees = append(employees, emp)
    }

    if err = rows.Err(); err != nil {
        log.Printf("Row iteration error: %v", err)
        return nil, err
    }

    return employees, nil
}

func InsertEmployee(emp models.Employee) error {
    _, err := config.DB.Exec(
        "INSERT INTO employees (name, age, position, department_id, salary) VALUES ($1, $2, $3, $4, $5)",
        emp.Name, emp.Age, emp.Position, emp.DepartmentId, emp.Salary,
    )

    if err != nil {
        log.Printf("Error inserting employee: %v", err)
        return err
    }

    return nil
}

func DeleteEmployeeById(id int) error {
    _, err := config.DB.Exec("DELETE FROM employees WHERE id = $1", id)
    if err != nil {
        log.Printf("Error deleting employee with ID %d: %v", id, err)
        return err
    }
    return nil
}

func GetEmployeeById(id int) (models.Employee, error) {
    var emp models.Employee
    err := config.DB.QueryRow(
        "SELECT id, name, age, position, department_id, salary FROM employees WHERE id = $1",
        id,
    ).Scan(&emp.Id, &emp.Name, &emp.Age, &emp.Position, &emp.DepartmentId, &emp.Salary)

    if err != nil {
        log.Printf("Error fetching employee with ID %d: %v", id, err)
        return models.Employee{}, err
    }

    return emp, nil
}

func UpdateEmployee(emp models.Employee) error {
    _, err := config.DB.Exec(
        "UPDATE employees SET name = $1, age = $2, position = $3, department_id = $4, salary = $5 WHERE id = $6",
        emp.Name, emp.Age, emp.Position, emp.DepartmentId, emp.Salary, emp.Id,
    )

    if err != nil {
        log.Printf("Error updating employee with ID %d: %v", emp.Id, err)
        return err
    }

    return nil
}
