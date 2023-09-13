container_tag = punch:1.2

build: 
	go build

podman:
	podman build --tag $(container_tag) .

docker:
	podman build --tag $(container_tag) .