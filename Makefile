
test:
	@go test -v ./...

bench:
	@go test -bench=BenchmarkLogger_JSON -benchtime=10s -benchmem