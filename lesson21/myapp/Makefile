all: clean run
build:
	@echo "executed build!"
	@go build server.go
	@#mv main lesson

clean:
	@echo "executed clean!"
	@if [ -e "./server" ]; then rm -f "../server"; else echo "File does not exist"; fi

run: build
	@echo "executed run!"
	@./server

