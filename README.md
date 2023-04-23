# to-buy-api
Golang API of shopping list to quickly go through API development in Golang

Tools used:
- hot reload: air
- db: postgres
- orm: gorm (go get gorm.io/driver/postgres) https://github.com/cosmtrek/air
- gitlab-runner / docker
- VPS (Ubuntu 22.04)

Useful commands:
- docker compose build
- docker compose up
- docker compose run --service-ports api bash -> to run 'go get {package-name}' commands
- docker build . -t local-tobuy-api
- docker run -dp 3000:3000 --env-file .env local-tobuy-api -> running with .env file
- docker stop local-tobuy-api || docker rm local-tobuy-api
- DB_USER=stray DB_NAME=tobuy DB_PASSWORD=sezamopen7 DB_PORT=5432 DB_HOST=#dbhost# go test -v ./...

Topics to cover:
- Goroutines
- Testing
- DB access / Gorm