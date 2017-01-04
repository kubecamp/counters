NAME=counters
RELEASE:=$(shell git rev-parse --verify --short HEAD)
USER=ipedrazas
BINARY=main

all: push

clean:
	docker rmi ${NAME} &>/dev/null || true
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi


build: clean
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
	docker build --pull=true --no-cache -t ${USER}/${NAME}:${RELEASE} .
	docker tag  ${USER}/${NAME}:${RELEASE} ${USER}/${NAME}:latest

push: build
	# docker login -u ${USER}
	docker push ${USER}/${NAME}:${RELEASE}
	docker push ${USER}/${NAME}:latest
