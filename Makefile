
ARTEFACTS_BUCKET_NAME=artefacts-lambda-$(PROJECT_NAME)
GOBUILD=env GOOS=linux go build -v -ldflags="-d -s -w" -a -tags netgo -installsuffix netgo -o
SHELL := /bin/bash

.PHONY: help 

help: ## Print this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-25s\033[0m %s\n", $$1, $$2}'

clean: ## Remove binaries
	@rm -rf ./bin || true

test: ## Run all tests
	@cd src; go test ./...; cd ..

init: ## Initialize the go module. Usage: make init modname=<module_name>
	go mod init $(modname)

build: clean ## Build new binaries
	go get -u
	go mod vendor
	$(GOBUILD) bin/sec-jwt-verifier src/jwt-authorizer/main.go
	$(GOBUILD) bin/hello-world src/hello-world/main.go

install-deps: ## Install serverless framework
	npm install -g serverless

package: build ## Package the application 
	serverless package
	@echo "Application packaged successfully"

deploy: build ## Deploy the application into given stage. Usage: `make deploy stage=<stage>`
	serverless deploy --stage $(stage) --verbose

remove: ## Undeploy the application from given stage. Usage: `make remove stage=<stage>`
	serverless remove --stage $(stage)
