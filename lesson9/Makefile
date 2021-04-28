build:
	go build main.go

clean:
	if [ -e "./main" ]; then rm "./main"; else echo "File does not exist"; fi

run:
	if [ -e "./main" ]; then ./main; else go build main.go; ./main; fi

