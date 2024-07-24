SERVICE_DIRS := $(wildcard ./services/*)

up:
	docker-compose --env-file .env \
		-f build/accounts.yml \
		-f build/auth.yml \
		-f build/kafka.yml \
		-f build/media_manager.yml \
		-f build/menu.yml \
		-f build/traefik.yml \
	up --build --remove-orphans

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
