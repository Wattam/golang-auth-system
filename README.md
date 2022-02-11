# Shoe Store Authentication System
![Shoe Store Logo](img/../assets/ShoeStore.png)

### REST API for a shoe store with Golang, GORM, Gin Web Framework, jwt-go and Swagger.

This API implements CRUD functions for shoes and users in addition to an authentication system through a JWT token. A local PostgreSQL database was used, setted through the `/database/databaseConfig.go` file.


### To run the project
- Once you are in the `/main` folder run the command:
```
go run main.go
```
- Or run directly the command:
```
go run main/main.go
```
---
### To get all the dependencies
- Gin Web Framework:
```
go get -u github.com/gin-gonic/gin
```
- GORM (ORM library):
```
go get -u gorm.io/gorm
```
- Go Postgres driver:
```
go get github.com/lib/pq
```
- jwt-go:
```
go get github.com/golang-jwt/jwt
```
- Swagger 2.0:
	- Go to [Swagger](https://github.com/go-swagger/go-swagger/releases), download the latest version of the executable, rename it to `swagger` then put it on your `%GOPATH%/bin`.

---
### To syncronize all the dependencies
```
export GO111MODULE=on
```
```
go mod tidy
```
---
### To generate Swagger documentation
- Run one of the following commands:
```
swagger serve ./swagger.json
```
```
swagger serve -F swagger ./swagger.json
```

---
### Packages
- `main`: router creation and calls for GET, POST, PUT and DELETE requests.
- `model`: data model for Shoe, User and Login.
- `database`: database configuration and initialization.
- `controller`: Shoe, User and Login controllers with functions to perform HTTP requests.
- `service`: JWT and sha256 services.
- `main/middleware`: middleware for authentication.

---
### Models
- #### `Login`

|  Attribute |  Type  |
|:----------:|:------:|
| Credential | string |
|  Password  | string |

- #### `Shoe`

| Attribute |   Type  |
|:---------:|:-------:|
|     ID    |   uint  |
|    Name   |  string |
|   Color   |  string |
|   Price   | float64 |

- #### `User`

| Attribute |  Type  |
|:---------:|:------:|
|     ID    |  uint  |
|  Username | string |
|   Email   | string |
|  Password | string |



---
### API Requests
#### `Login Request`

- `POST` `localhost:8080/api/v1/login`: attempts to authenticate a user on the system, using his credentials, generating a JWT token in case of success.
- Body example:
```json
{
	"credential": "username1",
	"password": "password1"
}
```
- Response example:
```json
{
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOjEsImV4cCI6MTY0NDU0MzY4MiwiaWF0IjoxNjQ0NTQwMDgyLCJpc3MiOiJ1c2VyLWFwaSJ9.CfkSEiBLr1x7Fct0XhZHpz468kD-a-nQ1feAFqhOPqQ"
}
```
- In order for any other request to be allowed they need to have a valid JWT token in their `Bearer` field.
---

#### `GetAllShoes Request`

- `GET` `localhost:8080/api/v1/shoes/get`: gets all shoes from the database.
- Response example:
```json
[
	{
		"id": "1",
		"name": "Name 1",
		"color": "Color 1",
		"price": "1"
	},
	{
		"id": "2",
		"name": "Name 2",
		"color": "Color 2",
		"price": "2"
	}
]
```

---
#### `GetShoe Request`

- `GET` `localhost:8080/api/v1/shoes/{id}`: gets a specific shoe from the database using the ID.
- Response example:
```json
{
	"id": "1",
	"name": "Name 1",
	"color": "Color 1",
	"price": "1"
}
```

---
#### `PostShoe Request`

- `POST` `localhost:8080/api/v1/shoes/post`: creates a shoe in the database passing the parameters through the body.
- Body example:
```json
{
	"name":"Shoe 1",
	"color":"Color 1",
	"price":"1"
}
```
- Response example:
```json
{
	"id": "1",
	"name": "Name 1",
	"color": "Color 1",
	"price": "1"
}
```

---
#### `PutShoe Request`

- `PUT` `localhost:8080/api/v1/shoes/put`: edits a shoe in the database with the ID equal to the one passed on the body.
- Body example:
```json
{
	"id":"1",
	"name":"Shoe 1 updated",
	"color":"Color 1 updated",
	"price":"1.1"
}
```
- Response example:
```json
{
	"id": "1",
	"name": "Shoe 1 updated",
	"color": "Color 1 updated",
	"price": "1.1"
}
```

---
#### `DeleteShoe Request`

- `DELETE` `localhost:8080/api/v1/shoes/{id}`: deletes a shoe from the database using the ID.

---
#### `GetAllUsers Request`

- `GET` `localhost:8080/api/v1/users/get`: gets all users from the database.
- Response example:
```json
[
	{
		"id": "1",
		"username": "username1",
		"email": "email1@email1",
		"password": "8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918"
	},
	{
		"id": "2",
		"username": "username2",
		"email": "email2@email2",
		"password": "46070d4bf934fb0d4b06d9e2c46e346944e322444900a435d7d9a95e6d7435f5"
	}
]
```

---
#### `GetUser Request`

- `GET` `localhost:8080/api/v1/users/{id}`: gets a specific user from the database using the ID.
- Response example:
```json
{
	"id": "1",
	"username": "username1",
	"email": "email1@email1",
	"password": "8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918"
}
```

---
#### `PostUser Request`

- `POST` `localhost:8080/api/v1/users/post`: creates a user in the database passing the parameters through the body.
- Body example:
```json
{
	"username": "username1",
	"email": "email1@email1",
	"password": "password1"
}
```
- Response example:
```json
{
	"id": "1",
	"username": "username1",
	"email": "email1@email1",
	"password": "8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918"
}
```

---
#### `PutUser Request`

- `PUT` `localhost:8080/api/v1/users/put`: edits a user in the database with the ID equal to the one passed on the body.
- Body example:
```json
{
	"id": "1",
	"username": "usernameEdited",
	"email": "emailEdited@emailEdited",
	"password": "passwordEdited"
}
```
- Response example:
```json
{
	"id": "1",
	"username": "usernameEdited",
	"email": "emailEdited@emailEdited",
	"password": "7e8b9345207e74ab0138e11f344afcf10f28a70afbca98e52e9bd370e16e0408"
}
```

---
#### `DeleteUser Request`

- `DELETE` `localhost:8080/api/v1/users/{id}`: deletes a user from the database using the ID.

---