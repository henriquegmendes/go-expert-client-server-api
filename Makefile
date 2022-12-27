install:
	go get

run-server-with-sqlite:
	docker build -t exchange-server-image .
	docker run -p 8080:8080 --name exchange-server --rm -it exchange-server-image

run-client:
	go run main.go client
