# MYPATH-BE

Backend API for MyPath career growth application.

## Features

- User registration with secure password hashing
- PostgreSQL connection with retry logic
- Modular Go project structure

## Requirements

- Go 1.20+
- PostgreSQL
- `.env` file with the following:

```env
PORT=8080
DATABASE_URL=postgres://postgres:yourpassword@localhost:5432/mydb
