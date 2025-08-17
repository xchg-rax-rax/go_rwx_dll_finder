build: ./src/main.go
	mkdir -p ./bin
	GOOS=windows GOARCH=amd64 go build  -o ./bin/rwxfinder.exe ./src/main.go

clean:
	rm -rf ./bin/*
