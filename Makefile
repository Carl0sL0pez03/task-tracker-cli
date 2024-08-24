# Binary name
BINARY_NAME := task-cli

# JSON file path
TASKS_FILE := tasks.json

# Compiling the binary
build:
	go build -o $(BINARY_NAME)

# Execute the binary with parameters
run: build
	./$(BINARY_NAME) $(ARGS)

# Clean the generated files
clean:
	rm -f $(BINARY_NAME)

# Execute the binary with an example command
example: build
	./$(BINARY_NAME) list

# Help to display the available commands
help:
	@echo "Available commands:"
	@echo "  make build      - Compiles project"
	@echo "  make run ARGS='[args]' - Execute the binary with arguments"
	@echo "  make clean      - Removes the compiled binary"
	@echo "  make example    - Run the binary with an example command"
	@echo "  make help       - Show help"