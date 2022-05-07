#!make

start:
	go run ./app/main.go
docker_db:
	docker run --name sPostgres -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -d postgres
docker_stop:
	docker rm -f sPostgres