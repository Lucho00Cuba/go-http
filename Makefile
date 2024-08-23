# Variables de color (ajusta o define según sea necesario)
YELLOW = \033[33m
GREEN = \033[32m
CYAN = \033[36m
RESET = \033[0m

ifneq (,$(wildcard ./.env))
    include .env
    export
endif

TARGETS := $(basename $(notdir $(wildcard scripts/*)))
TARGET := $(firstword $(MAKECMDGOALS))
ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))

.PHONY: help

## Help
help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo "  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}"
	@echo ''
	@echo "${CYAN}Commands${RESET}"
	@for target in $(TARGETS); do \
	    echo "    ${YELLOW}$$target${RESET}"; \
	done
	@echo ''
	@awk 'BEGIN {FS = ":.*?## "} { \
	        if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
	        else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
	        }' $(MAKEFILE_LIST)

### Ensure all scripts have execution permissions
ensure-permissions:
	@for script in scripts/*; do \
	    if [ ! -x $$script ]; then \
	        echo "Adding execute permissions to $$script"; \
	        chmod +x $$script; \
	    fi \
	done

%::
	@true

$(TARGETS): ensure-permissions
	@./scripts/$@ $(ARGS)