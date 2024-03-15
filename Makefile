test:
	mkdir -p test_result
	go test ./... -cover -coverprofile=test_result/coverage.out
	go tool cover -html=test_result/coverage.out -o test_result/cover.html
build:
	go build -o server .
run:
	./server
clean:
	rm -f server
check-swagger:
	export PATH=$$(go env GOPATH)/bin:$$PATH && which swag || go install github.com/swaggo/swag/cmd/swag@latest
genswag:
	export PATH=$$(go env GOPATH)/bin:$$PATH && swag init

# on MacOS please add export PATH=$(go env GOPATH)/bin:$PATH in your .zshrc file
tools:
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/cosmtrek/air@latest

# build-local:
# 	docker build -t rw-fiber . 
# build-dev:
# 	docker buildx build --push --tag inyourtime/ecommerce-be:dev --platform=linux/amd64 .	