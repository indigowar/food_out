SERVICE_DIRS := $(wildcard ./services/*)

up:
	docker-compose -f build/docker-compose.yaml up --build

run:
	@echo "coming soon..."

gen:
	@for dir in $(SERVICE_DIRS); do \
		$(MAKE) -C $$dir gen; \
	done

test:
	@for dir in $(SERVICE_DIRS); do \
		$(MAKE) -C $$dir test; \
	done
