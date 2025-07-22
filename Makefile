.Phony: run stop migrate

migrate:
	docker exec -it backend_exercise_db psql -U backend_exercise_user -W backend_exercise

run:
	docker compose up --build -d

stop:
	docker compose down