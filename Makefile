TEST ?= $$(go list ./...)
ifndef POSTGRES_DATA_SOURCE
	export POSTGRES_DATA_SOURCE=postgres://postgres@/terraform_provider_sql?sslmode=disable
endif
ifndef MYSQL_DATA_SOURCE
	export MYSQL_DATA_SOURCE=root@/terraform_provider_sql
endif

default: build
.PHONY: default

help:
	@echo "Main commands:"
	@echo "  help            - show this message"
	@echo "  build (default) - build the terraform provider"
	@echo "  test            - runs unit tests"
	@echo "  testacc         - runs acceptance tests"
.PHONY: help

build:
	go build
.PHONY: build

test:
	go test $(TEST) -v $(TESTARGS)
.PHONY: test

testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS)
.PHONY: testacc

release-patch: guard
	@make release PATCH=$$(( $(PATCH) + 1 ))

release-minor: guard
	@make release MINOR=$$(( $(MINOR) + 1 )) PATCH=0

release-major: guard
	@make release MAJOR=$$(( $(MAJOR) + 1 )) MINOR=0 PATCH=0

release: guard
	@sed -i'.bak' 's/^VERSION=.*$$/VERSION=$(MAJOR).$(MINOR).$(PATCH)/g' Makefile
	@rm Makefile.bak
	@git add Makefile
	@git commit -m 'bump version to $(MAJOR).$(MINOR).$(PATCH)'
	@git tag -a v$(MAJOR).$(MINOR).$(PATCH) -m 'v$(MAJOR).$(MINOR).$(PATCH)'
	@git push --follow-tags

guard:
# @git diff-index --quiet HEAD || (echo "There are changes in the repo, won't release. Commit everything and run this from a clean repo"; exit 1)
ifneq ($(shell echo `git branch --show-current`),master)
	@echo "Releases can only be done from master" && exit 1
endif
