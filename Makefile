run:
	echo "running server"
	go run server.go

migrateDB:
	echo "running migrations"
	go run migrate/migrate.go

build:
	echo "building application"
	go build -o bin/main main.go

compile:
	echo "Compiling for every OS and Platform"
	Compiling for every OS and Platform
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go
	GOOS=linux GOARCH=386 go build -o bin/main-freebsd-386 main.go
	GOOS=windows GOARCH=386 go build -o bin/main-freebsd-386 main.go