run:
	go run main.go

docker-mysql-create:
	docker run -d -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=secretpassword mysql:8.0

docker-mysql-force-delete:
	docker rm -f mysql

docker-mysql-recreate: docker-mysql-force-delete docker-mysql-create