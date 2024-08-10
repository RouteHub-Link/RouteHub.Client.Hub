serve:
	go run .

hotserve:
	templ generate --watch --proxy="http://localhost:8080" --cmd="go run ."

templ:
	templ generate -f TEMPL_EXPERIMENT=rawgo