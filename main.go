package main

import (
    "employee/config"
    "employee/handlers"
    "employee/utils"
    "net/http"
)

func main() {
    // Connect DB
    utils.LoadEnv()
    config.ConnectDB()

    // Serve static files
    fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // Handle employee page
    http.HandleFunc("/employees", handlers.EmployeeIndexHandler)
    http.HandleFunc("/employees/new", handlers.EmployeeCreateHandler)
    http.HandleFunc("/employees/insert", handlers.EmployeeInsertHandler)
    http.HandleFunc("/employees/edit", handlers.EmployeeEditHandler)
    http.HandleFunc("/employees/update", handlers.EmployeeUpdateHandler)
    http.HandleFunc("/employees/delete", handlers.EmployeeDeleteHandler)

    // Start server
    port := utils.GetEnv("APP_PORT", "8080")
    http.ListenAndServe(":" + port, nil)
}
