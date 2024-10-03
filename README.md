# SimpleAPI
## A RESTful API built using Go, Gin, and PostgreSQL for managing a product catalog. This application allows you to perform CRUD (Create, Read, Update, Delete) operations on products.

## Technologies Used
- Go: The programming language used for building the application.
- Gin: A web framework for Go to handle routing and middleware.
- GORM: An Object-Relational Mapping (ORM) library for Golang, used for interacting with the PostgreSQL database.
- PostgreSQL: A relational database management system used to store product data.
- godotenv: A Go library to load environment variables from a .env file.

## Setup Instructions
1. Clone the Repository
2. Create an env file like the one below
```
DB_HOST=db
DB_USER=postgres
DB_PASSWORD=mypassword
DB_NAME=postgres
DB_PORT=5432
```
3. Run postgresql on your local machine or through docker.
4. Build and run the go application.

*Note: You can import the API curl requests to postman from the simplerAPI json file located inside the main project folder*

*Note2: Dockerization of the app is not set up completely, there were some issues loading the .env file.*


