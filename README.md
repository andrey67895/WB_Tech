Для проверки кода:
`go vet && docker run --rm -v $(pwd):/app -w /app golang:1.24 sh -c "go install honnef.co/go/tools/cmd/staticcheck@latest && staticcheck -f stylish ./..."`
