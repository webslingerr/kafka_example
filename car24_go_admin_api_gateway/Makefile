CURRENT_DIR=$(shell pwd)

APP=$(shell basename ${CURRENT_DIR})

APP_CMD_DIR=${CURRENT_DIR}/cmd

IMG_NAME=${APP}
REGISTRY=gitlab.udevs.io:5050
PROJECT_NAME=car24
UPDATE_TAG=${UPDATE_TAG}
TAG=latest
ENV_TAG=latest
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
# Including

POSTGRESQL_URL='postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=disable'


deploy-to-prod:
	kubectl apply -f .kube/prod
	kubectl set image -n microservices deployment/${DEPLOYMENT} ${DEPLOYMENT}=${REGISTRY}/${PROJECT_NAME}/${APP}:${UPDATE_TAG}
pull-submodule:
	git submodule update --init --recursive

update-submodule:
	git submodule update --remote --merge

submodule-gen:
	rm -rf modules/car24
	mkdir -p modules/car24
	rsync -r --exclude '.git' car24_modules/ modules/car24

swag-init:
	swag init -g api/main.go -o api/docs

vendor-update:
	go get -u ./...
	go mod vendor

run: 
	go run cmd/main.go

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

install:
	swag init -g api/main.go -o api/docs
	go mod download
	go mod vendor
	go run cmd/main.go

mark-as-production-image:
	docker tag ${REGISTRY}/${IMG_NAME}:${TAG} ${REGISTRY}/${IMG_NAME}:production
	docker push ${REGISTRY}/${IMG_NAME}:production

build-image:
	docker build --rm -t ${REGISTRY}/${PROJECT_NAME}/${IMG_NAME}:${TAG} .
	docker tag ${REGISTRY}/${PROJECT_NAME}/${IMG_NAME}:${TAG} ${REGISTRY}/${PROJECT_NAME}/${IMG_NAME}:${ENV_TAG}

push-image:
	docker push ${REGISTRY}/${PROJECT_NAME}/${IMG_NAME}:${TAG}
	docker push ${REGISTRY}/${PROJECT_NAME}/${IMG_NAME}:${ENV_TAG}

migrate-local-up:
	migrate -database ${POSTGRESQL_URL} -path migrations up

migrate-local-down:
	migrate -database ${POSTGRESQL_URL} -path migrations down

create-new-migration:
	migrate create -ext sql -dir migrations -seq $(name)