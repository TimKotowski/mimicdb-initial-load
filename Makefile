RUN_DB_PERM :=  $(PWD)/volumes/docker-entrypoint-init/create_target.sh
WAIT_PERM := $(PWD)/volumes/docker-entrypoint-init/docker_runing.sh

start-dev:
	docker compose -f compose-dev.yaml down
	docker compose -f compose-dev.yaml up -d --build
	@echo "waiting containers to finishing booting up..."
	chmod +x $(RUN_DB_PERM)
	chmod +x $(WAIT_PERM)

	$(WAIT_PERM)
	$(RUN_DB_PERM)
	@echo "containers up and running"

run-db-script:
	chmod +x $(RUN_DB_PERM)
	$(RUN_DB_PERM)
