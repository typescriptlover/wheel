cmd := docker-compose -p wheel
devcmd := $(cmd) -f docker-compose.dev.yml
logs := docker-compose logs -f --tail=10

# run with nodemon and watch for changes
# npm i -g nodemon (if doesn't run)
dev:
	nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run main.go

# build docker compose dev
docker-dev:
	$(devcmd) up -d --build
	
# build docker compose prod, 
# and view logs
docker-prod:
	$(cmd) up -d --build && $(logs)