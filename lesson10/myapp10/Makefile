all: clean run
build:
	@echo "executed build!"
	@go build main.go
	@#mv main lesson

clean:
	@echo "executed clean!"
	@if [ -e "./main" ]; then rm -f "../main"; else echo "File does not exist"; fi

run: build
	@echo "executed run!"
	@./main

