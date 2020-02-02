docker-run: docker-down
	@./build/build.sh
	@docker-compose -f build/docker/docker-compose.yaml up --build

docker-down:
	@docker-compose -f build/docker/docker-compose.yaml down

app-shell:
	@docker exec -it purwalenta_binary sh
