PROJECT_NAME = "tatl-test-app"

BINARY_NAME = ${PROJECT_NAME}
BINARIES = "./bin"
MAIN_DIR = "cmd/${BINARY_NAME}"
GIT_LOCAL_NAME = "rodziievskyi-maksym"
GIT_LOCAL_EMAIL = "rodziyevskydev@gmail.com"
GITHUB = "github.com/${GIT_LOCAL_NAME}/${PROJECT_NAME}"
GIT_SSH_KEY_NAME = "newgit_id_rsa"

POSTGRES_URL = "postgresql://postgres:postgres@localhost:5434/daily-dose?sslmode=disable"


init:
	@echo "::> Creating a module root..."
	@go mod init ${GITHUB}
	@mkdir "cmd" && mkdir "cmd/"${BINARY_NAME}
	@touch ${MAIN_DIR}/main.go
	@echo "package main\n\nimport \"fmt\"\n\nfunc main(){\n\tfmt.Println(\"${BINARY_NAME}\")\n}" > ${MAIN_DIR}/main.go
	@touch VERSION && echo 0.0.1 > VERSION
	@echo "::> Finished!"

build:
	@echo "::> Building..."
	@go build -o ${BINARIES}/${BINARY_NAME} ${MAIN_DIR}/main.go
	@echo "::> Finished!"

run:
	@go build -o ${BINARIES}/${BINARY_NAME} ${MAIN_DIR}/main.go
	@${BINARIES}/${BINARY_NAME}

clean:
	@echo "::> Cleaning..."
	@go clean
	@rm -rf ${BINARIES}
	@go mod tidy
	@echo "::> Finished"

git-con:
	@eval "$(ssh-agent -s)"
	@ssh-add ~/.ssh/${GIT_SSH_KEY_NAME}

local-git:
	@git config --local user.name ${GIT_LOCAL_NAME}
	@git config --local user.email ${GIT_LOCAL_EMAIL}
	@git config --local --list

## Database operations
postgres:
	docker run --name go-simple-bank-db -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.1-alpine

create-db:
	docker exec -it go-simple-bank-db createdb --username=root --owner=root go-simple-bank

drop-db:
	docker exec -it go-simple-bank-db dropdb go-simple-bank

test:
	go test -v -cover ./...

.PNONY: init build run clean git-con local-git git-init test