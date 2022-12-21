# GolangAPI

First you need to Install Go

then run this command to run a postgres container which maps port 5432 to localhost:5432

```shell
go run .
```

this will make available the following endpoints:

### GET=> http://localhost:8080/home
returns a JSON object with:
a greeting message

### GET => http://localhost:8080/healthcheck
returns a JSON object with:
a healthcheck status message and timestamp

### POST => http://localhost:8080/users
returns a JSON object with:
a confirmation message and status code

### GET => http://localhost:8080/users
returns a JSON array with:
user objects 

### GET => http://localhost:8080/users/:id
returns a JSON object with:
user properties

### PUT => http://localhost:8080/users/:id
returns a JSON object with:
a confirmation message and status code

### DELETE => http://localhost:8080/users/:id
returns a JSON object with:
a confirmation message and status code
