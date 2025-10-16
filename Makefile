#auth-service:50051 pg 5431
#chat-service:50052    5432
#social-service:50053  5433
#users-service:50054   5434


build:
	docker-compose build

up:
	docker-compose up -d

down:
	docker-compose down

logs:
	docker-compose logs -f

clean:
	docker-compose down -v --rmi local


.PHONY: build up