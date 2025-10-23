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

# Get Books

<b>URL</b> - http://localhost:8080/api/v1/book/books
<b>CURL</b>    
```cmd  
curl -X 'GET' \
  'http://localhost:8080/api/v1/book/books' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NjEyMzA1MTcsInVzZXJuYW1lIjoiYWRtaW4ifQ.78EfuvZULiPWj6QGLpT3y0HMDArcQvsD2q6Da2CSL_s'
```

<b> Response </b>
```json
{
  "status_code": "200",
  "total": 1,
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

# Get Book By Id

<b>URL</b> - http://localhost:8080/api/v1/book/books
<b>CURL</b>    
```cmd  
curl -X 'GET' \
  'http://localhost:8080/api/v1/book/GO1011121314' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NjEyMzA1MTcsInVzZXJuYW1lIjoiYWRtaW4ifQ.78EfuvZULiPWj6QGLpT3y0HMDArcQvsD2q6Da2CSL_s'
```

<b> Response </b>
```json
{
  "status_code": "200",
  "response": {
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
}
```

# Update Book By Id [PUT]

<b>URL</b> - http://localhost:8080/api/v1/book/GO1011121314
<b>CURL</b>    
```cmd  
curl -X 'PUT' \
  'http://localhost:8080/api/v1/book/GO1011121314' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NjEyMzA1MTcsInVzZXJuYW1lIjoiYWRtaW4ifQ.78EfuvZULiPWj6QGLpT3y0HMDArcQvsD2q6Da2CSL_s' \
  -H 'Content-Type: application/json' \
  -d '{
   "id": "GO1011121314",
    "title": "Known Publisher",
    "author": "Go Programmers",
    "publisher": "Basics of Golang && Concuurency Control",
    "published_at": "October 23, 2025",
    "isbn": "1234509876",
    "pages": 400,
    "language": "English",
    "price": 550
}'
```
<b> Request Body </b>
```json
{
   "id": "GO1011121314",
    "title": "Known Publisher",
    "author": "Go Programmers",
    "publisher": "Basics of Golang && Concuurency Control",
    "published_at": "October 23, 2025",
    "isbn": "1234509876",
    "pages": 400,
    "language": "English",
    "price": 550
}
```

<b> Response </b>
```json
{
  "status_code": "200",
  "response": {
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
}
```

# Update Book By Id [PATCH]

<b>URL</b> - http://localhost:8080/api/v1/book/GO1011121314
<b>CURL</b>    
```cmd  
curl -X 'PATCH' \
  'http://localhost:8080/api/v1/book/GO1011121314' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NjEyMzA1MTcsInVzZXJuYW1lIjoiYWRtaW4ifQ.78EfuvZULiPWj6QGLpT3y0HMDArcQvsD2q6Da2CSL_s' \
  -H 'Content-Type: application/json' \
  -d '{
  "id": "GO1011121314",
  "language": "Hindi",
  "pages": 1000,
  "price": 600
}'
```
<b> Request Body </b>
```json
{
  "id": "GO1011121314",
  "language": "Hindi",
  "pages": 1000,
  "price": 600
}
```

<b> Response </b>
```json
{
  "status_code": "200",
  "status_message": "Book ID GO1011121314 record has been updated successfully",
  "response": {
    "id": "GO1011121314",
    "title": "Known Publisher",
    "author": "Go Programmers",
    "publisher": "Basics of Golang && Concuurency Control",
    "published_at": "October 23, 2025",
    "isbn": "1234509876",
    "pages": 1000,
    "language": "Hindi",
    "price": 600
  }
}
```
# Delete Book By Id [DELETE]

<b>URL</b> - http://localhost:8080/api/v1/book/GO1011121314
<b>CURL</b>    
```cmd  
curl -X 'DELETE' \
  'http://localhost:8080/api/v1/book/GO1011121314' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NjEyMzA1MTcsInVzZXJuYW1lIjoiYWRtaW4ifQ.78EfuvZULiPWj6QGLpT3y0HMDArcQvsD2q6Da2CSL_s'
```
<b> Response </b>
```json
{
  "status_code": "200",
  "status_message": "Book with ID GO1011121314 and Title Known Publisher has been deleted"
}
```

## Operations on Cart

# Add items to cart [POST]
<b>URL </b> - http://localhost:8080/api/v1/cart
<b>CURL </b>
```cmd
curl -X 'POST' \
  'http://localhost:8080/api/v1/cart' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NjEyMzA1MTcsInVzZXJuYW1lIjoiYWRtaW4ifQ.78EfuvZULiPWj6QGLpT3y0HMDArcQvsD2q6Da2CSL_s' \
  -H 'Content-Type: application/json' \
  -d '{
  "book_id": "GO1011121314"
}'
```
<b> Request Body </b>
```json
{
  "book_id": "GO1011121314"
}
```

<b>Response</b>
```json
{
  "cart": [
    {
      "book_id": "GO1011121314",
      "quantity": 1,
      "price": 250
    }
  ],
  "success_code": "200",
  "success_message": "Book added to your cart"
}
```

# View cart 
<b>URL </b> - http://localhost:8080/api/v1/cart [GET]
<b>CURL </b>
```cmd
curl -X 'GET' \
  'http://localhost:8080/api/v1/cart' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NjEyMzA1MTcsInVzZXJuYW1lIjoiYWRtaW4ifQ.78EfuvZULiPWj6QGLpT3y0HMDArcQvsD2q6Da2CSL_s'
```

<b>Response</b>
```json
{
  "cart": [
    {
      "book_id": "GO1011121314",
      "quantity": 1,
      "price": 250
    }
  ],
  "success_code": "200",
  "total_items": 1,
  "total_price": 250
}
```

## Operations on Orders

# Order [POST]
<b>URL </b> - http://localhost:8080/api/v1/order
<b>CURL </b>
```cmd
curl -X 'POST' \
  'http://localhost:8080/api/v1/order' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NjEyMzQyOTgsInVzZXJuYW1lIjoiYWRtaW4ifQ.sCgjMRx36z0ayPjjZL4lMrB7WJe346mjx_tynE0TyYE' \
  -d ''
```

<b>Response</b>
```json
{
  "status_code": "200",
  "status_message": "Order placed successfully",
  "response": {
    "order_id": "BOOKOD3049905315",
    "name": "Alok Kumar Sharma",
    "cart": [
      {
        "book_id": "GO1011121314",
        "quantity": 1,
        "price": 250
      }
    ],
    "address": "Address",
    "order_date": "2025-10-23 20:15:16",
    "total": 250,
    "status": {
      "type": "pending",
      "reason": "Awaiting payment"
    },
    "payment": {
      "paid": false
    }
  }
}
```
# Get Order By Id [GET]
<b>URL </b> - http://localhost:8080/api/v1/order/BOOKOD3049905315
<b>CURL </b>
```cmd
curl -X 'GET' \
  'http://localhost:8080/api/v1/order/BOOKOD3049905315' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NjEyMzQyOTgsInVzZXJuYW1lIjoiYWRtaW4ifQ.sCgjMRx36z0ayPjjZL4lMrB7WJe346mjx_tynE0TyYE'
```

<b>Response</b>
```json
{
  "status_code": "200",
  "status_message": "Fetched order successfully",
  "response": {
    "order_id": "BOOKOD3049905315",
    "name": "Alok Kumar Sharma",
    "cart": [
      {
        "book_id": "GO1011121314",
        "quantity": 1,
        "price": 250
      }
    ],
    "address": "Address",
    "order_date": "2025-10-23 20:15:16",
    "total": 250,
    "status": {
      "type": "pending",
      "reason": "Awaiting payment"
    },
    "payment": {
      "paid": false
    }
  }
}
```

# Update Order Status [GET]
<b>URL </b> - http://localhost:8080/api/v1/order/BOOKOD3049905315/status
<b>CURL </b>
```cmd
curl -X 'PUT' \
  'http://localhost:8080/api/v1/order/BOOKOD3049905315/status' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NjEyMzQyOTgsInVzZXJuYW1lIjoiYWRtaW4ifQ.sCgjMRx36z0ayPjjZL4lMrB7WJe346mjx_tynE0TyYE' \
  -H 'Content-Type: application/json' \
  -d '{
  "reason": "Payment Awaiting",
  "shipping_carrier": "known",
  "tracking_number": "BOOK0987654321",
  "type": "Transit"
}'
```

<b>Response</b>
```json
{
  "status_code": "200",
  "status_message": "Order status updated successfully",
  "response": {
    "order_id": "BOOKOD3049905315",
    "name": "Alok Kumar Sharma",
    "cart": [
      {
        "book_id": "GO1011121314",
        "quantity": 1,
        "price": 250
      }
    ],
    "address": "Address",
    "order_date": "2025-10-23 20:15:16",
    "total": 250,
    "status": {
      "type": "Transit",
      "reason": "Payment Awaiting",
      "tracking_number": "BOOK0987654321",
      "shipping_carrier": "known"
    },
    "payment": {
      "paid": false
    }
  }
}
```

# Update Payment Status [GET]
<b>URL </b> - http://localhost:8080/api/v1/order/BOOKOD3049905315/payment
<b>CURL </b>
```cmd
curl -X 'PUT' \
  'http://localhost:8080/api/v1/order/BOOKOD3049905315/payment' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NjEyMzQyOTgsInVzZXJuYW1lIjoiYWRtaW4ifQ.sCgjMRx36z0ayPjjZL4lMrB7WJe346mjx_tynE0TyYE' \
  -H 'Content-Type: application/json' \
  -d '{
  "method": "UPI",
  "paid": true,
  "paid_on": "October 31, 2025",
  "reference": "UPIBOOKOD3049905315"
}'
```

<b>Response</b>
```json
{
  "status_code": "200",
  "status_message": "Order payment updated successfully",
  "response": {
    "order_id": "BOOKOD3049905315",
    "name": "Alok Kumar Sharma",
    "cart": [
      {
        "book_id": "GO1011121314",
        "quantity": 1,
        "price": 250
      }
    ],
    "address": "Address",
    "order_date": "2025-10-23 20:15:16",
    "total": 250,
    "status": {
      "type": "Transit",
      "reason": "Payment Awaiting",
      "tracking_number": "BOOK0987654321",
      "shipping_carrier": "known"
    },
    "payment": {
      "paid": true,
      "method": "UPI",
      "paid_on": "October 31, 2025",
      "reference": "UPIBOOKOD3049905315"
    }
  }
}
```

