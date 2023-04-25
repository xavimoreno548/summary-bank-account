### Testing
test:
	@echo "==> Running unit tests"
	go test ./... -covermode=atomic -coverpkg=./... -count=1 -race

run-local:
	@echo "==> Building and running by Docker"
	docker build -t summary-bank-account .
	docker run -p 8080:8080 summary-bank-account