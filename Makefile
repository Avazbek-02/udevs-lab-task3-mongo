CURRENT_DIR=$(shell pwd)

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}

SWAGGER := ~/go/bin/swag
SWAGGER_DOCS := docs
SWAGGER_INIT := $(SWAGGER) init -g ./api/router.go -o $(SWAGGER_DOCS)

swag-gen:
	$(SWAGGER_INIT)

run:
	$(SWAGGER_INIT)
	@docker network create global || true
	@docker-compose up --build -d
	@echo "Containers are up and running"

stop:
	@docker-compose down
	@echo "Containers stopped"

rebuild:
	@docker-compose down
	@docker-compose up --build -d
	@echo "Containers rebuilt and running"
