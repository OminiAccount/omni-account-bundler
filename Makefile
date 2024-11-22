DOCKERCOMPOSE-OA := docker-compose -f docker-compose.yml
OABUNDLER := oa-bundler
RUN-OABUNDLER := $(DOCKERCOMPOSE-OA) up -d $(OABUNDLER)
STOP-OABUNDLER := $(DOCKERCOMPOSE-OA) stop $(OABUNDLER) && $(DOCKERCOMPOSE-OA) rm -f $(OABUNDLER)

.PHONY: run
run:
	$(RUN-OABUNDLER)

.PHONY: stop
stop:
	$(STOP-OABUNDLER)


## Help display.
## Pulls comments from beside commands and prints a nicely formatted
## display with the commands and their usage information.
.DEFAULT_GOAL := help

.PHONY: help
help: ## Prints this help
		@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| sort \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'