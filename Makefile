# Main Compose
COMMON_COMPOSE=docker-compose -f docker-compose.yml

#####################################################
# Cluster
#####################################################
.PHONY: provision build up upd down downall
PROVISION_APP=go install github.com/cosmtrek/air@v1.29.0
provision: build gotidy
	${COMMON_COMPOSE} run --rm app sh -c '${PROVISION_APP}'

build:
	${COMMON_COMPOSE} build

up:
	${COMMON_COMPOSE} up

upd:
	${COMMON_COMPOSE} up -d

down:
	${COMMON_COMPOSE} down --volumes

# Remove image, Volumes, Networks and Orphans
downall:
	${COMMON_COMPOSE} down --rmi all --volumes --remove-orphans

#####################################################
# Go
#####################################################
.PHONY: gosh gotidy
gosh:
	${COMMON_COMPOSE} run --rm app bash

gotidy:
	${COMMON_COMPOSE} run --rm app sh -c 'go mod tidy'
