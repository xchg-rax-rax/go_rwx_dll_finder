build: ./src/main.go
	mkdir -p ./bin
	GOOS=windows GOARCH=amd64 go build ./src/main.go -o ./bin/rwxfinder.exe

clean:
	rm -rf ./bin/*
