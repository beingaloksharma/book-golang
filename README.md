
# GoBookStore API

## Overview

BookStore is a lightweight bookstore web API built with **Golang (Gin-Gonic)**.  
It allows users to register, log in, manage books, add them to a cart, place orders, and retrieve order summaries.  
   
The app features **modular architecture**, **JWT authentication**, and **file-based order summaries** on checkout.

---

## Table of Contents
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Models](#models)
- [DTOs](#dtos)
- [API Endpoints](#api-endpoints)
- [Request Examples](#request-examples)
- [Response Examples](#response-examples)
- [Order File Generation](#order-file-generation)
---

## Tech Stack

- **Language:** Go 1.20+
- **Framework:** Gin-Gonic
- **Auth:** JWT
- **Validation:** go-playground/validator
- **Logging:** ZeroLog
- **Storage:** In-memory data maps (User, Book, Cart, Orders)

---

## Project Structure

go-bookstore-api     
├── main.go # Entry point and route registration  
├── users.go # User creation, login, profile, address  
├── books.go # CRUD operations for books  
├── cart.go # Add/view cart and checkout    
├── orders.go # Order placement, retrieval, file creation    
├── jwt.go # JWT token generation and validation   
├── util.go # Helpers for validation & JSON binding   
└── orders_output/ # Folder auto-created for order summary text files   

---

## Models

| Model | Description | Fields |
|-------|--------------|---------|
| **ModelUser** | Stores user login info | `Name`, `Username`, `Password` |
| **ModelBook** | Represents a book entity | `ID`, `Title`, `Author`, `Publisher`, `PublishedAt`, `ISBN`, `Pages`, `Language`, `Price` |
| **CartItem** | Items persisted in user cart | `BookID`, `Quantity`, `Price` |
| **CheckOut** | User’s cart summary | `Username`, `Cart`, `Total` |
| **Orders** | Represents placed orders | `OrderID`, `Name`, `Cart`, `Address`, `OrderDate`, `Total` |

---

## DTOs

| DTO | Purpose | Fields |
|-----|----------|---------|
| **RequestLogin** | Login payload | `Username`, `Password` |
| **RequestUserAddressToAdd** | Add user address | `Add` |
| **PatchBookDTO** | Partial update of book data | `ID`, `Pages`, `Language`, `Price` |
| **SuccessDTO** | Standard success response | `SuccessCode`, `SuccessMessage`, `Total`, `CustomMessage` |
| **ErrorDTO** | Standard error response | `ErrorCode`, `ErrorMessage` |

---

## API Endpoints

### Authentication
| Route | Method | Description |
|--------|---------|-------------|
| `/signup` | POST | Register new user |
| `/signin` | POST | Login and receive token |

### User
| Route | Method | Description |
|--------|---------|-------------|
| `/user/address` | POST | Add address |
| `/user/profile/:username` | GET | Retrieve user profile summary |

### Books
| Route | Method | Description |
|--------|---------|-------------|
| `/book` | POST | Create new book |
| `/book/books` | GET | Retrieve all books |
| `/book/:id` | GET | Fetch book by ID |
| `/book/:id` | PATCH | Update partial fields |
| `/book/:id` | PUT | Replace existing book |
| `/book/:id` | DELETE | Remove book by ID |

### Cart
| Route | Method | Description |
|--------|---------|-------------|
| `/cart` | POST | Add a book to cart |
| `/cart` | GET | View current user cart |

### Orders
| Route | Method | Description |
|--------|---------|-------------|
| `/order` | POST | Place order — auto-generates text summary file |
| `/order` | GET | View all orders |
| `/order/:id` | GET | View single order details |

---

## Request Examples

### Signup 

POST /signup

```json
{
"name": "ADMIN",
"username": "admin",
"password": "admin123"
}
```

### Login 

POST /signin

```json
{
"username": "admin",
"password": "admin123"
}
```

### Add Address  

POST /user/address

```json
{
"address": "address"
}
```

### Create Book  

POST /book

```json
{
"id": "B1001",
"title": "Learning Go",
"author": "Go",
"publisher": "GO",
"published_at": "2024-09-10",
"isbn": "123456789",
"pages": 200,
"language": "English",
"price": 500
}
```

### Add to Cart
POST /cart

```json
{
"book_id": "B1001"
}
```

## Response Examples

```json
{
"status_code": "200",
"status_message": "Order placed successfully",
"response": {
"order_id": "BOOKOD1234567890",
"name": "ADMIN",
"cart": [
{ "book_id": "B1001", "quantity": 1, "price": 500 }
],
"address": "address",
"order_date": "2025-10-22 16:00:00",
"total": 500
}
}
```

```json
{
"error_code": "400",
"error_message": "Cart is empty, please add items before placing an order"
}
```

---

## Order File Generation

On successful order placement:
- A `.txt` file is created in `/orders_output/`
- Filename format: `<username>_<orderid>.txt`
- Example content:

```text
Order Summary
Username: admin
Name: ADMIN
Order ID: BOOKOD1234567890
Order Date: 2025-10-22 16:00:00
Delivery Address: Address

Items Ordered:

BookID: B1001 | Quantity: 1 | Price: 500

Total Amount: ₹500
Generated On: 2025-10-22 16:00:01

```
