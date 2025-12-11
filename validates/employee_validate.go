package validates

import (
    "employee/utils"
    "net/http"
    "strconv"
)

func ValidateEmployeeInsertUpdate(w http.ResponseWriter, r *http.Request) error {
    ageStr := r.FormValue("age")
    salaryStr := r.FormValue("salary")
    age, _ := strconv.Atoi(ageStr)
    salary, _ := strconv.ParseFloat(salaryStr, 64)

    if ageStr == "" {
        utils.WriteError(w, "Age is required", 422)
        return nil
    }
    if age <= 0 {
        utils.WriteError(w, "Age must be a positive integer", 422)
        return nil
    }
    if salaryStr == "" {
        utils.WriteError(w, "Salary is required", 422)
        return nil
    }
    if salary <= 0 {
        utils.WriteError(w, "Salary must be a positive number", 422)
        return nil
    }
    if salary <= 0 {
        utils.WriteError(w, "Salary must be a positive number", 422)
        return nil
    }

    return nil
}
