install:
	go get

run-server-with-sqlite:
	docker build -t exchange-server-image .
	docker run -p 8000:8000 --name exchange-server -it exchange-server-image

run-client:
	go run main.go client
