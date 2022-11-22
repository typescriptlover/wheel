cmd := docker-compose -p wheel
devcmd := $(cmd) -f docker-compose.dev.yml
logs := docker-compose logs -f --tail=10

# build docker compose dev
dev:
	$(devcmd) up -d --build
	
# build docker compose prod, 
# and view logs
prod:
	$(cmd) up -d --build && $(logs)