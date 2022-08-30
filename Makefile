COMPOSE = docker compose

start:
	$(COMPOSE) up

recreate:
	$(COMPOSE) up --force-recreate

connect:
	$(COMPOSE) exec -it api /bin/bash