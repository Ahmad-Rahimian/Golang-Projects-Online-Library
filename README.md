Online Library

version 2 :

project-root/
│
├── cmd/  
│ └── app/
│ └── main.go  
│
├── internal/  
│ ├── freebook/  
│ │ ├── model.go  
│ │ ├── repository.go  
│ │ ├── service.go  
│ │ └── handler.go  
│ │
│ ├── paidbook/  
│ │ ├── model.go  
│ │ ├── repository.go  
│ │ ├── service.go  
│ │ └── handler.go  
│ │
│ ├── article/  
│ │ ├── model.go  
│ │ ├── repository.go  
│ │ ├── service.go  
│ │ └── handler.go  
│ │
│ └── router/  
│ └── router.go  
│
├── api/  
│ └── api.go  
│
├── pkg/  
│ ├── db.go  
│ └── config.go  
│
├── uploads/  
│ ├── images  
│ └── pdfs  
│
├── docs/  
├── go.mod
└── go.sum

<!--

version 1 :

project-root/
│
├── cmd/
│ └── app/
│ └── main.go
│
├── internal/
│ ├── domain/
│ │ ├── freebook.go
│ │ ├── paidbook.go
│ │ └── article.go
│ │
│ ├── repository/
│ │ ├── freebook_repository.go
│ │ ├── paidbook_repository.go
│ │ └── article_repository.go
│ │
│ ├── service/
│ │ ├── freebook_service.go
│ │ ├── paidbook_service.go
│ │ └── article_service.go
│ │
│ ├── handler/
│ │ ├── freebook_handler.go
│ │ ├── paidbook_handler.go
│ │ └── article_handler.go
│ │
│ └── router/
│ └── router.go
│
├── api/
│ └── api.go
│
├── pkg/
│ ├── db.go
│ └── config.go
|
├── uploads/
│ ├── images
│ └── pdfs
│
├── docs/
├── go.mod
└── go.sum -->
