-- ================================================
-- Create database
-- ================================================
CREATE DATABASE go_employee;

-- ================================================
-- Switch to database
-- ================================================
\c go_employees;

-- ================================================
-- Create departments table
-- ================================================
CREATE TABLE departments (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- ================================================
-- Create employees table
-- ================================================
CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    age INT NOT NULL,
    position VARCHAR(255),
    department_id INT REFERENCES departments(id) ON DELETE SET NULL,
    salary NUMERIC(15,2)
);

-- ================================================
-- Insert data into departments table
-- ================================================
INSERT INTO departments (name) VALUES
('Human Resources'),
('Finance'),
('Engineering'),
('Marketing'),
('Operations');

