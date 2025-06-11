PROGY = MyProgy

EXEC = progy

PACKAGE = github.com/gin-gonic/gin\
					github.com/jackc/pgx/v5\
					github.com/jackc/pgx/v5/pgxpool\
					github.com/ilyakaznacheev/cleanenv\
					github.com/joho/godotenv\
					github.com/golang-jwt/jwt/v5

.PHONY: all build run clean init get

all: build run

build:
	mkdir -p build
	go build -o ./build/$(EXEC) ./cmd/main.go

run:
	./build/$(EXEC)

clean:
	rm -rf ./build/$(EXEC)

init:
	go mod init $(PROGY)

get:
	go get -u $(PACKAGE)

docker_up:
	sudo docker-compose up

docker_down:
	sudo docker-compose down -v

docker_clean:
	sudo docker rmi simple_crud-myapp
