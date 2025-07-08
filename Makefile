.Phony: run stop

run:
	docker compose up --build -d

stop:
	docker compose down