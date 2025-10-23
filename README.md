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

# Add a new book

<b>URL</b> - http://localhost:8080/api/v1/book   
<b>CURL</b>    
```cmd  
curl -X 'POST' \
  'http://localhost:8080/api/v1/book' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NjEyMzA1MTcsInVzZXJuYW1lIjoiYWRtaW4ifQ.78EfuvZULiPWj6QGLpT3y0HMDArcQvsD2q6Da2CSL_s' \
  -H 'Content-Type: application/json' \
  -d '{
  "author": "Go Programmers",
  "id": "GO1011121314",
  "isbn": "1234509876",
  "language": "English",
  "pages": 500,
  "price": 250,
  "published_at": "October 23, 2025",
  "publisher": "Basics of Golang",
  "title": "Unknown Publisher"
}'
```

<b> Request Body </b>
```json
{
  "status_code": "200",
  "status_message": "Book Record Created",
  "response": [
    {
      "id": "GO1011121314",
      "title": "Unknown Publisher",
      "author": "Go Programmers",
      "publisher": "Basics of Golang",
      "published_at": "October 23, 2025",
      "isbn": "1234509876",
      "pages": 500,
      "language": "English",
      "price": 250
    }
  ]
}
```

<b> Response </b>
```json
{
  "status_code": "200",
  "status_message": "Book Record Created",
  "response": [
    {
      "id": "GO1011121314",
      "title": "Unknown Publisher",
      "author": "Go Programmers",
      "publisher": "Basics of Golang",
      "published_at": "October 23, 2025",
      "isbn": "1234509876",
      "pages": 500,
      "language": "English",
      "price": 250
    }
  ]
}
```


