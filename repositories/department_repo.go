package repositories

import (
    "employee/config"
    "employee/models"
    "log"
)

func GetAllDepartments() ([]models.Department, error) {
    rows, err := config.DB.Query("SELECT id, name FROM departments")
    if err != nil {
        log.Printf("Error querying departments: %v", err)
        return nil, err
    }
    defer rows.Close()

    var departments []models.Department
    for rows.Next() {
        var dept models.Department
        if err := rows.Scan(&dept.Id, &dept.Name); err != nil {
            log.Printf("Error scanning department row: %v", err)
            return nil, err
        }
        departments = append(departments, dept)
    }

    if err = rows.Err(); err != nil {
        log.Printf("Row iteration error: %v", err)
        return nil, err
    }

    return departments, nil
}
