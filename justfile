# Run test
[group('development')]
test:
    go test ./...

# Run the application
[group('development')]
run:
    go run .
