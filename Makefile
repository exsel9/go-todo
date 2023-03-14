run:
	go run main.go

docker-mysql:
	docker run -d -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=secretpassword mysql:8.0