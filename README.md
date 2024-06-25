# Splitly User Management Service

This is a User Management Service built with Go and SQL. It uses the Gin framework for handling HTTP requests and bcrypt for password hashing.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Contributing](#contributing)

## Installation

To install this project, you need to have Go installed on your machine. Then, clone this repository and run the following command in the project directory:

```bash
go mod download
```

## Usage

To start the server, run the following command in the project directory:

```bash
go run main.go
```

## API Endpoints

- `GET /users`: Fetch all users
- `GET /users/:id`: Fetch a single user by ID
- `POST /users`: Create a new user
- `PUT /users/:id`: Update a user by ID
- `DELETE /users/:id`: Delete a user by ID

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
