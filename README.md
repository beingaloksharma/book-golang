# Operations on Book

## Data Stored

```go
var BooksData []BookDTO
```

## Book DTO

```go
type BookDTO struct {
	ID          string  `json:"id"`           // Unique identifier for the book
	Title       string  `json:"title"`        // Title of the book
	Author      string  `json:"author"`       // Author's name
	Publisher   string  `json:"publisher"`    // Publisher's name
	PublishedAt string  `json:"published_at"` // Publication date (could be string or time.Time)
	ISBN        string  `json:"isbn"`         // ISBN number
	Pages       int     `json:"pages"`        // Number of pages
	Language    string  `json:"language"`     // Language of the book
	Price       float64 `json:"price"`        // Price of the book
}
```
## Sucess DTO
```go
type SuccessDTO struct {
	SuccessCode    string `json:"status_code"`
	SuccessMessage string `json:"status_message,omitempty"`
	Total          int    `json:"total,omitempty"`
	CustomMessage  any    `json:"books,omitempty"`
}
```

## Patch DTO
```go
type PatchBookDTO struct {
	ID       string  `json:"id"`       // Unique identifier for the book
	Pages    int     `json:"pages"`    // Number of pages
	Language string  `json:"language"` // Language of the book
	Price    float64 `json:"price"`    // Price of the book
}
```

## Error DTO
```go
type ErrorDTO struct {
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}
```

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
