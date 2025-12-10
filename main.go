package main

import (
    "employee/config"
    "employee/handlers"
    "employee/middleware"
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
    http.Handle("/employees", middleware.LoggingMiddleware(http.HandlerFunc(handlers.EmployeeIndexHandler)))
    http.Handle("/employees/new", middleware.LoggingMiddleware(http.HandlerFunc(handlers.EmployeeCreateHandler)))
    http.Handle("/employees/insert", middleware.LoggingMiddleware(http.HandlerFunc(handlers.EmployeeInsertHandler)))
    http.Handle("/employees/edit", middleware.LoggingMiddleware(http.HandlerFunc(handlers.EmployeeEditHandler)))
    http.Handle("/employees/update", middleware.LoggingMiddleware(http.HandlerFunc(handlers.EmployeeUpdateHandler)))
    http.Handle("/employees/delete", middleware.LoggingMiddleware(http.HandlerFunc(handlers.EmployeeDeleteHandler)))

    // Start server
    port := utils.GetEnv("APP_PORT", "8080")
    http.ListenAndServe(":" + port, nil)
}
