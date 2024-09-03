container_tag := "punch:1.3"

build: 
	go build

podman:
	podman build --tag {{container_tag}} .

docker:
	docker build --tag {{container_tag}} .