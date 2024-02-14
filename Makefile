
.DEFAULT_GOAL := build

APP = promptpay

build:
	@echo "Building..."
	go build -o ${APP}

run: build
	@echo "Running..."
	./${APP}

test:
	@echo "Testing..."
	go test ./...

clean:
	@echo "Cleaning up..."
	go clean

dockerbuild:
	@echo "Building docker image..."
	docker build -t ${APP} .

dockerrun: dockerbuild
	@echo "Running docker image..."
	docker stop promptpay || true
	docker run --rm -d --name ${APP} -p 8080:8080 ${APP}

dockerrunpure:
	@echo "Running docker image..."
	@docker stop promptpay >/dev/null 2>&1 || true
	docker run --rm -d --name ${APP} -p 8080:8080 ${APP}

dockerstop:
	@echo "Stopping docker image..."
	@docker stop promptpay >/dev/null 2>&1 || true
