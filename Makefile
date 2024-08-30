serve:
	go run .

serve-mqtt:
	HOSTING_MODE=MQTT go run .

hotserve:
	templ generate --watch --proxy="http://localhost:8080" --cmd="go run ."

templ:
	templ generate

docker:
	docker build -t routehub-client-hub:latest .

podman:
	podman build -t routehub-client-hub:latest .

podman-serve-mqtt:
	make podman
	podman rm -f routehub-client-hub-mqtt || true
	podman run -d --name routehub-client-hub-mqtt --env-file ./.env -e HOSTING_MODE=MQTT --net=host  routehub-client-hub:latest

keydb:
	docker run -d --name keydb -p 6379:6379 eqalpha/keydb

keydbPodman:
	podman run -d --name keydb -p 6379:6379 eqalpha/keydb

timescaledb:
	docker run -d --name timescaledb -p 5432:5432 -e POSTGRES_PASSWORD=password timescale/timescaledb-ha:pg16

podmanTimescaledb:
	podman run -d --name timescaledb -p 5432:5432 -e POSTGRES_PASSWORD=password timescale/timescaledb-ha:pg16