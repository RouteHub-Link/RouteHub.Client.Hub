serve-rest:
	HOSTING_MODE=REST go run .

serve-mqtt:
	HOSTING_MODE=MQTT go run .

hotserve:
	templ generate --watch --proxy="http://localhost:8080" --cmd="go run ."

templ:
	templ generate

docker:
	docker build -t routehub-client-hub:latest .

docker-serve-mqtt:
	make docker
	docker rm -f routehub-client-hub-mqtt || true
	docker run -d --name routehub-client-hub-mqtt --env-file ./.env -e HOSTING_MODE=MQTT --net=host  routehub-client-hub:latest

podman:
	podman build -t routehub-client-hub:latest .

podman-serve-mqtt:
	make podman
	podman rm -f routehub-client-hub-mqtt || true
	podman run -d --name routehub-client-hub-mqtt --env-file ./.env -e HOSTING_MODE=MQTT --net=host  routehub-client-hub:latest


podman-keydb:
	podman run -d --name keydb -p 6379:6379 eqalpha/keydb


podman-Timescaledb:
	podman run -d --name timescaledb -p 5432:5432 -e POSTGRES_PASSWORD=password timescale/timescaledb-ha:pg16

# Different Port Setup

keydb:
	docker run -d --name keydb -p 6380:6379 eqalpha/keydb

timescaledb:
	docker run -d --name timescaledb -p 5532:5432 -e POSTGRES_PASSWORD=password timescale/timescaledb-ha:pg16

serve-mqtt-bash:
	bash -c "export REDIS_PORT=6380 && export HOSTING_MODE=MQTT && go run ."

serve-rest-bash:
	bash -c "export REDIS_PORT=6380 && \
	 export HOSTING_MODE=REST && \
	 export TIMESCALE_DB=postgres://postgres:password@localhost:5532/postgres && \
	 export PORT=8088 && \
	 go run ."


