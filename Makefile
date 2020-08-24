all:build

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o dev .
	docker build -t gogateway/dev .