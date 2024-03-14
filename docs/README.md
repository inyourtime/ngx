# Documents

Place your documents, Miro Links, Postman Collection or even swagger files

## Directory structure

```text
docs
├── ...
└── postman_collection
    └── POSTMAN\ API\ Documents.postman_collection.json
```

## If you prefer to use Go Swagger, you can add this useful script to your Makefile

```Makefile
check-swagger:
  which swag || go install github.com/swaggo/swag/cmd/swag@latest

genswag: check-swagger
  swag init
```
