deps:
	@go get -v ./...

build:
	@go build -o ./out/refund-latam-mock
	@echo "Refund Latam Mock built in ./out/refund-latam-mock"
	
run: build
	./out/refund-latam-mock ${PORT}