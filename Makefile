# Main Compose
COMMON_COMPOSE=docker-compose -f docker-compose.yml

#####################################################
# Cluster
#####################################################
.PHONY: provision build up upd down downall
PROVISION_APP=go install github.com/cosmtrek/air@v1.29.0
provision: build gotidy yarninstall
	${COMMON_COMPOSE} run --rm app sh -c '${PROVISION_APP}'

build:
	${COMMON_COMPOSE} build

up:
	${COMMON_COMPOSE} up

upd:
	${COMMON_COMPOSE} up -d

down:
	${COMMON_COMPOSE} down

# Remove image, Volumes, Networks and Orphans
downall:
	${COMMON_COMPOSE} down --rmi all --volumes --remove-orphans

#####################################################
# Go
#####################################################
.PHONY: gosh gotidy
gosh:
	${COMMON_COMPOSE} run --rm app ash

gotidy:
	${COMMON_COMPOSE} run --rm app sh -c 'go mod tidy'

#####################################################
# Nodejs
#####################################################
.PHONY: yarninstall nodesh
yarninstall:
	${COMMON_COMPOSE} run --rm nodejs yarn install

nodesh:
	${COMMON_COMPOSE} run --rm nodejs bash

#####################################################
# DB
#####################################################
.PHONY: dbc
dbc:
	${COMMON_COMPOSE} run --rm db mysql -h online_shop_db -u root -p

#####################################################
# DB Migration
#####################################################
.PHONY: migrate_g migrate migrate_up migrate_down
migrate_g:
	${COMMON_COMPOSE} run --rm db_migrate g ${NAME} 

migrate:
	${COMMON_COMPOSE} run --rm db_migrate migrate ${ARG}

migrate_up:
	make migrate ARG=up

migrate_down:
	make migrate ARG=down

migrate_force:
	make migrate ARG=force ${VER}

#####################################################
# Misc
#####################################################
# Fix permission
fixperm:
	sudo chown -R $(shell whoami):$(shell whoami) .
