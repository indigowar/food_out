SERVICE_DIRS := $(wildcard ./services/*)

up:
	docker-compose --env-file .env \
		-f build/accounts.yml \
		-f build/auth.yml \
		-f build/kafka.yml \
		-f build/media_manager.yml \
		-f build/menu.yml \
		-f build/orders.yml \
		-f build/order_history.yml \
		-f build/traefik.yml \
	up --build --remove-orphans


down:
	docker-compose --env-file .env \
		-f build/accounts.yml \
		-f build/auth.yml \
		-f build/kafka.yml \
		-f build/media_manager.yml \
		-f build/menu.yml \
		-f build/orders.yml \
		-f build/order_history.yml \
		-f build/traefik.yml \
	down

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
