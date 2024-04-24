.PHONY: run
run: run-go run-py

run-go:
	@echo "Starting Go API..."
	@go run cmd/api/main.go &

run-py:
	@echo "Starting Python API..."
	@cd remote/ && python transcribe_server.py &
