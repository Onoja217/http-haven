# http-haven

This project is a simple Go HTTP server built as a set of exercises. It’s mainly for practicing how to build routes, handle requests, work with query parameters, headers, and HTTP status codes using only the Go standard library.

---

## Repo submission

Repo to submit when you’re done: **http-haven**

Each exercise is implemented as part of the project and kept inside the same `http-haven` directory.

---

## How to run

Start the server:

```bash
go run .
```
+ Server runs on:

```
http://localhost:8080
```
### Project structure

+ handlers/ → route logic

+ routes/ → route definitions

+ main.go → entry point

+ test_endpoints.sh → automated test script

+ go.mod → Go module file

### Exercises

1. /ping

+ Simple health check endpoint.

```
GET /ping
```
### Response:

```
pong
```

2. /hello

+ Returns a greeting using query parameters.

```
GET /hello?name=Alice
```
+ If no name is provided:
```
Hello, Guest!
```

3. /count

### Counts characters from request body.

```
GET /count
```
### Returns instruction message.
```
POST /count
```
### Returns number of characters in the request body.

4. /calculate

### Simple math API.
```
GET /calculate?op=add&a=10&b=5
```
### Supported operations:
```
add
subtract
multiply
```
### Response:
```
Result: 15
```
+ Invalid input returns 400 Bad Request.

5. /agent

### Returns the client User-Agent.
```
GET /agent
```
### Response:
```
You are visiting us using: <User-Agent>
```
6. /dashboard

### Protected route using headers.

+ Required header:
```
X-API-Key: secret123
```
### Responses:
```
Valid key → Welcome to the secure dashboard
Invalid/missing → 401 Unauthorized
```
7. /legacy → /v2

### Redirect example.
```
GET /legacy
```
+ Redirects permanently to /v2.
```
/v2
Welcome to version 2
Testing
```
### Run the automated test script:
```
./test_endpoints.sh
```
+ It validates all endpoints from Exercise 1 to 7.

### Notes

+ Built using only Go standard library

+ No frameworks used

+ One server, multiple endpoints

+ Focus is understanding HTTP fundamentals in Go
