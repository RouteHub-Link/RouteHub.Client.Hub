serve:
	go run .

hotserve:
	templ generate --watch --proxy="http://localhost:8080" --cmd="go run ."

templ:
	templ generate -f TEMPL_EXPERIMENT=rawgo

docker:
	docker build -t routehub-client-hub:latest .

podman:
	podman build -t routehub-client-hub:latest .

keydb:
	docker run -d --name keydb -p 6379:6379 eqalpha/keydb

keydbPodman:
	podman run -d --name keydb -p 6379:6379 eqalpha/keydb

clickhouse:
	docker run -d -p 8123:8123 -p 9000:9000  --name clickhouse-server --ulimit nofile=262144:262144 clickhouse/clickhouse-server

podmanClickhouse:
	podman run -d -p 8123:8123 -p 9000:9000  --name clickhouse-server --ulimit nofile=262144:262144 clickhouse/clickhouse-server

timescaledb:
	docker run -d --name timescaledb -p 5432:5432 -e POSTGRES_PASSWORD=password timescale/timescaledb-ha:pg16

podmanTimescaledb:
	podman run -d --name timescaledb -p 5432:5432 -e POSTGRES_PASSWORD=password timescale/timescaledb-ha:pg16