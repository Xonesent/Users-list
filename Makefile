## Заблаговременно должен быть запущен docker
## Нужно проследить, что порт 5436 не занят прослушкой другого хоста (5436 - твой // 5432 - postgres)

## 1) Сейчас сборка ориентирована на докер композ / чтобы пользоваться этим надо в конфигах заменить: host: "localhost" / port: "5436"
hand-build:
	docker run --name=db -e POSTGRES_PASSWORD='test' -p 5436:5432 -d --rm postgres
	migrate -path ./schema -database 'postgres://postgres:test@localhost:5436/postgres?sslmode=disable' up
	go run cmd/main.go

## 1) Текущая сборка через докер образ = блатная
docker-build:
	docker-compose up --build user-app

## 1) Запуск unit тестирования на покрытие кода / пока в разработке
test:
	go test -v ./.../...

## 1) Смысл для чего нужен понятен, но еще ничего не реализовано
swag:
	swag init -g cmd/main.go