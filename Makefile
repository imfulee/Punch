build: 
	go build

podman:
	podman build --tag punch:1.0.0 .

docker:
	podman build --tag punch:1.0.0 .