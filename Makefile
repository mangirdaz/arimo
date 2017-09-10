SHELL=/bin/bash 
ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
IMAGES= \
mangirdas/nanny-client \
mangirdas/ocp-example-api:v0.1 \
mangirdas/ocp-example-api:v0.2 \
mangirdas/ocp-example-api:v0.3 \
mangirdas/docker-debug-container

SRC_REGISTRY=docker.io
DST_REGISTRY=localhost:5000


binary: 
	go build -tags="containers_image_openpgp"

docker-distribution: setup list

setup:
	docker stop registry ;\
	docker rm registry ; \
	docker run -d -p 5000:5000 --restart=always -v ${ROOT_DIR}/registry:/var/lib/registry --name registry registry:2 

populate: 
	@echo "The list is :"
	for img in  $(IMAGES); \
    do \
    echo $${img}	 ; \
	skopeo copy docker://${SRC_REGISTRY}/$${img} docker://${DST_REGISTRY}/$${img} --dest-tls-verify=false ;\
    done


