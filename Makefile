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