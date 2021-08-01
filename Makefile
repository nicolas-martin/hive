# .PHONY: all
# all: package

.DEFAULT_GOAL := help
.PHONY: help
help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
	| sed -n 's/^\(.*\): \(.*\)##\(.*\)/\1\3/p' \
	| column -t  -s ':'

.PHONY: api-test 
api-test: ## :Runs the api tests.
	cd ./api && go test ./...

.PHONY: api-test api-tag
api-tag: ## :Runs tests and tag new image.
	docker build -t hive-api ./api

.PHONY: api-run
api-run: ## :Runs the latest hive-api tag.
	docker run -it --rm --env-file ./api/.env -t hive-api

.PHONY: scheduler-test 
scheduler-test: ## :Runs the scheduler tests.
	cd ./scheduler && go test ./...

.PHONY: scheduler-test scheduler-tag
scheduler-tag: ## :Runs tests and tag new image.
	docker build -t hive-scheduler ./scheduler/

.PHONY: scheduler-run
scheduler-run: ## :Runs the latest hive-scheduler tag.
	docker run -it --rm --env-file ./scheduler/.env -t hive-scheduler
