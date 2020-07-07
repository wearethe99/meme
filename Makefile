include .env.make

REGISTRY_DOMAIN=registry.gitlab.com
REGISTRY_NAME=registry.gitlab.com/wearethe99/meme

help:
	@echo '  make docker:push'
	@echo '  make run:bot'

run\:server:
	go run *.go server

run\:bot:
	go run *.go bot --token=${BOT_TOKEN} --prefix=${BOT_PREFIX} --text=${BOT_TEXT}

docker\:push:
	docker login ${REGISTRY_DOMAIN} --username ${REGISTRY_USERNAME} --password ${REGISTRY_PASSWORD}
	docker build --no-cache -f ./Dockerfile -t ${REGISTRY_NAME}:latest .
	docker push ${REGISTRY_NAME}:latest
