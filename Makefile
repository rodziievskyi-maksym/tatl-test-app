PROJECT_NAME = "tatl-test-app"

BINARY_NAME = ${PROJECT_NAME}
BINARIES = "./bin"
MAIN_DIR = "cmd/${BINARY_NAME}"
GIT_LOCAL_NAME = "rodziievskyi-maksym"
GIT_LOCAL_EMAIL = "rodziyevskydev@gmail.com"
GITHUB = "github.com/${GIT_LOCAL_NAME}/${PROJECT_NAME}"
GIT_SSH_KEY_NAME = "newgit_id_rsa"

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
	@go run ${MAIN_DIR}/main.go --port=8080 --dbuser=root --dbpass=secret --dbhost=localhost --dbname=mydb

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


test:
	go test -v -cover ./...

.PNONY: init build run clean git-con local-git git-init test