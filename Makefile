docker-run: docker-down
	@./build/build.sh
	@docker-compose -f build/docker/docker-compose.yaml up --build --force-recreate

docker-down:
	@docker-compose -f build/docker/docker-compose.yaml down

docker-reset: docker-down
	@docker image rm purwalenta_binary
	@docker image rm purwalenta_postgre
	@docker image rm purwalenta_consul_server
	@docker volume rm docker_purwalenta_binary
	@docker volume rm docker_purwalenta_postgre
	@docker volume rm docker_purwalenta_consul_server

docker-hard-reset: docker-down
	@docker container prune
	@docker image prune
	@docker volume prune
	@docker system prune -a

app-shell:
	@docker exec -it purwalenta_binary sh
