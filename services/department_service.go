package services

import (
    "employee/models"
    "employee/repositories"
    "log"
)

func FetchAllDepartments() ([]models.Department, error) {
    departments, err := repositories.GetAllDepartments()
    if err != nil {
        log.Printf("Error fetching departments: %v", err)
        return nil, err
    }
    return departments, nil
}
