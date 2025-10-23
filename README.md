## Operations on User

# Signup
<b>URL </b> - http://localhost:8080/api/v1/signup   
<b>CURL </b>
```cmd
curl -X 'POST' \
  'http://localhost:8080/api/v1/signup' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "Alok Kumar Sharma",
  "password": "admin",
  "username": "admin"
}'
```
<b> Request Body </b>
```json
{
  "name": "Alok Kumar Sharma",
  "password": "admin",
  "username": "admin"
}
```

<b>Response</b>
```json
{
  "status_code": "200",
  "status_message": "New User Created Successfully",
  "response": {
    "password": "*****",
    "username": "admin"
  }
}
```

# Signin
<b>URL </b> - http://localhost:8080/api/v1/signin   
<b>CURL </b>
```cmd
curl -X 'POST' \
  'http://localhost:8080/api/v1/signin' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "password": "admin",
  "username": "admin"
}'
```
<b> Request Body </b>
```json
{
  "password": "admin",
  "username": "admin"
}
```
<b>Response</b>
```json
{
  "status_code": "200",
  "status_message": "Login SuccessFully for Username - admin",
  "response": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NjEyMzA1MTcsInVzZXJuYW1lIjoiYWRtaW4ifQ.78EfuvZULiPWj6QGLpT3y0HMDArcQvsD2q6Da2CSL_s"
}
```
# Add a new Address
<b>URL </b> - http://localhost:8080/api/v1/address   
<b>CURL </b>
```cmd
curl -X 'POST' \
  'http://localhost:8080/api/v1/user/address' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NjEyMzA1MTcsInVzZXJuYW1lIjoiYWRtaW4ifQ.78EfuvZULiPWj6QGLpT3y0HMDArcQvsD2q6Da2CSL_s' \
  -H 'Content-Type: application/json' \
  -d '{
  "address": "Address"
}'
```
<b> Request Body </b>
```json
{
  "address": "Address"
}
```
<b>Response</b>
```json
{
  "status_code": "200",
  "status_message": "User Address added Successfully",
  "response": [
    {
      "address": "Address"
    }
  ]
}
```

# Get User Details
<b>URL </b> - http://localhost:8080/api/v1/user/profile/admin
<b>CURL </b>
```cmd
curl -X 'GET' \
  'http://localhost:8080/api/v1/user/profile/admin' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NjEyMzA1MTcsInVzZXJuYW1lIjoiYWRtaW4ifQ.78EfuvZULiPWj6QGLpT3y0HMDArcQvsD2q6Da2CSL_s'
```
<b>Response</b>
```json
{
  "status_code": "200",
  "status_message": "Profile of Username admin",
  "response": {
    "name": "Alok Kumar Sharma",
    "username": "admin",
    "address": [
      {
        "address": "Address"
      }
    ]
  }
}
```

## Operations on Book

# POST - /book/create

Request
```json
{
    "id": "1",
    "title": "Go Lang",
    "author": "Google",
    "publisher": "Google",
    "published_at": "2015-10-26",
    "isbn": "978-0134190440",
    "pages": 400,
    "language": "English",
    "price": 50
}
```

Response
```json
{
    "status_code": "200",
    "status_message": "Book Record Created",
    "books": [
        {
            "id": "1",
            "title": "Go Lang",
            "author": "Google",
            "publisher": "Google",
            "published_at": "2015-10-26",
            "isbn": "978-0134190440",
            "pages": 400,
            "language": "English",
            "price": 50
        }
    ]
}
```

# GET - /book/book/:id

Request :: /book/book/1

Response
```json
{
    "status_code": "200",
    "books": {
        "id": "1",
        "title": "Go Lang",
        "author": "Google",
        "publisher": "Google",
        "published_at": "2015-10-26",
        "isbn": "978-0134190440",
        "pages": 500,
        "language": "Hindi",
        "price": 500
    }
}
```

# GET - /book/books

Request :: /book/books

Response
```json
{
    "status_code": "200",
    "total": 2,
    "books": [
        {
            "id": "1",
            "title": "Go Lang",
            "author": "Google",
            "publisher": "Google",
            "published_at": "2015-10-26",
            "isbn": "978-0134190440",
            "pages": 500,
            "language": "Hindi",
            "price": 500
        },
        {
            "id": "2",
            "title": "Java Programming",
            "author": "",
            "publisher": "",
            "published_at": "2015-10-26",
            "isbn": "978-0134190440",
            "pages": 400,
            "language": "English",
            "price": 50
        }
    ]
}
```

#  DELETE - /book/book/:id

Request :: /book/book/2


Response
```json
{
    "status_code": "200",
    "status_message": "Book having book ID - 2 and Title - Java Programming is deleted"
}
```

# PATCH - /book/book

Request 
```json
{
    "id": "1",
    "pages": 500,
    "language": "Hindi",
    "price": 500
}
```

Response
```json
{
    "status_code": "200",
    "status_message": "Book Id 1 record has been updated successfully",
    "books": {
        "id": "1",
        "title": "Go Lang",
        "author": "Google",
        "publisher": "Google",
        "published_at": "2015-10-26",
        "isbn": "978-0134190440",
        "pages": 500,
        "language": "Hindi",
        "price": 500
    }
}
```
