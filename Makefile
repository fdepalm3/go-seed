build:
	go build

test:
	go test -json > reports/test-report.json -cover -coverprofile=reports/test-coverage.out -race ./... && \
	go tool cover -html=reports/test-coverage.out                                                     \

run:
	go run .

lint:
	golangci-lint run --out-format colored-line-number --issues-exit-code 0 > reports/golangci-report.out

generate:
	go generate ./...