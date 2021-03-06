docker-run: docker-down
	@dep ensure -v
	@docker-compose -f build/docker/docker-compose.yaml up --build --force-recreate

docker-runbg: docker-down
	@dep ensure -v
	@docker-compose -f build/docker/docker-compose.yaml up -d --build --force-recreate

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

consul-backup:
	@docker exec -it purwalenta_consul_server consul snapshot save backup.snap
	@docker cp -a purwalenta_consul_server:backup.snap ~/backup.snap
