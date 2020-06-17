ifneq (,$(wildcard ./common.env))
    include common.env
    export
endif

.DEFAULT_GOAL := help

.PHONY: help
help: ## Self-documented Makefile
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' Makefile | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

cmd-exists-%:
	@hash $(*) > /dev/null 2>&1 || \
		(echo "ERROR: '$(*)' must be installed and available on your PATH."; exit 1)

guard-%:
	@if [ -z '${${*}}' ]; then echo 'ERROR: environment variable $* not set' && exit 1; fi


gen-app-env: guard-APP_ENV guard-APP_NAME  ## Generate dotenv for application
	@make -j gen-app-env-api gen-app-env-foo gen-app-env-bar gen-app-env-baz

gen-app-env-%: cmd-exists-envsubst
	@envsubst < ./$(*)/app.template.env > ./$(*)/app.env
	@echo "[$(*)] app.env generated."

gen-app-yaml: guard-APP_ENV guard-APP_NAME  ## Generate yaml for GAE
	@make -j gen-app-yaml-api gen-app-yaml-foo gen-app-yaml-bar gen-app-yaml-baz

gen-app-yaml-%: cmd-exists-envsubst gen-app-env-%
	@sh -ac '. ./$(*)/app.env && envsubst < ./$(*)/gae/app.template.yaml > ./$(*)/gae/app.yaml'
	@echo "[$(*)] app.yaml generated."


deploy-gae: cmd-exists-go cmd-exists-gcloud guard-GAE_VERSION  ## Deploy to GAE
	@GO111MODULE=on go mod vendor
	@GO111MODULE=off gcloud app deploy \
		api/gae/app.yaml \
		foo/gae/app.yaml \
		bar/gae/app.yaml \
		baz/gae/app.yaml \
		--quiet --no-promote --version=$(GAE_VERSION)
	@rm -rf ./vendor


deploy-run:  ## Deploy to Cloud Run
	@make -j deploy-run-api deploy-run-foo deploy-run-bar deploy-run-baz

deploy-run-%: cmd-exists-docker cmd-exists-gcloud guard-GCP_PROJECT guard-APP_NAME guard-RUN_REGION
	@docker build . --file ./$(*)/run/Dockerfile --tag gcr.io/$(GCP_PROJECT)/$(APP_NAME)/$(*)
	@docker push gcr.io/$(GCP_PROJECT)/$(APP_NAME)/$(*)
	@gcloud run deploy $(APP_NAME)-$(*) --platform managed --image gcr.io/$(GCP_PROJECT)/$(APP_NAME)/$(*) \
		--allow-unauthenticated --region $(RUN_REGION) --update-env-vars=`awk '{ORS=","} {print}' ./$(*)/app.env`