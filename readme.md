# http-haven

This project is a simple Go HTTP server built as a set of exercises. It is designed to help understand core HTTP concepts in Go using only the standard library. It covers routing, query parameters, headers, request methods, status codes, redirects, form parsing, ServeMux sub-routing, and HTML templates.

## Repo submission

Repo to submit when you’re done: http-haven

All exercises are implemented inside a single Go module under the http-haven directory.

## How to run

Run the server:

go run .

Server runs on:
```
http://localhost:8080
```
## Project structure

handlers/ → all request handlers (business logic)

routes/ → route registration

main.go → application entry point

http-resurgence/ → exercise drafts / workspace files

test_endpoints.sh → automated test script

go.mod → Go module definition

## Exercises

1. /ping

GET /ping

+ Response:

```
pong
```

2. /hello

+ GET /hello?name=Alice

+ Response:
```
Hello, Alice!
```
+ If no name is provided:
```
Hello, Guest!
```

3. /count

+ GET /count → returns instruction message

+ POST /count → returns number of characters in request body


4. /calculate

+ GET /calculate?op=add&a=10&b=5

+ Supported operations:

+ add

+ subtract

+ multiply

+ Response:
```
Result: 15
```
+ Invalid input returns:
```
400 Bad Request
```

5. /agent

GET /agent

+ Response:
```
You are visiting us using: <User-Agent>
```

6. /dashboard

+ Header required:

X-API-Key: secret123

+ Valid:

Welcome to the secure dashboard

+ Invalid or missing:

401 Unauthorized


7. /legacy → /v2

+ GET /legacy

Redirects permanently (301) to:
/v2

Response:
/v2 → Welcome to version 2


8. HTTP Resurgence (Advanced)

method-inspector → returns HTTP method used

echo → echoes POST body

headers → reads custom headers

form → parses form data

status → dynamic HTTP status codes

api/v1/ping → returns pong (ServeMux subtree)

api/v1/greet → greeting endpoint

render → HTML template rendering

## Key Concepts Practiced

HTTP routing using net/http

Query parameters (r.URL.Query)

Headers (r.Header.Get)

HTTP methods (r.Method)

Status codes (http.Status*)

Request body handling (io.ReadAll)

Form parsing (r.ParseForm)

Redirects (http.Redirect)

ServeMux subtree routing

HTML templates (html/template)

## Testing

chmod +x test_endpoints.sh
```
./test_endpoints.sh
```
+ The script validates:

All 7 core endpoints

All HTTP Resurgence endpoints

Correct status codes, headers, and responses

## Notes

Built using only Go standard library

No frameworks used

Single server architecture

Focus is HTTP fundamentals

Designed for backend learning progression

## Final Goal

By completing this project you should understand:

How HTTP servers work internally in Go

How routing works under the hood

How headers, methods, and status codes behave

How real backend APIs are built from scratch