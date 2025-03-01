package main

import (
	"flag"
	"fmt"
	"github.com/rodziievskyi-maksym/tatl-test-app/api"
	authrepository "github.com/rodziievskyi-maksym/tatl-test-app/internal/auth/infrastructure/repository"
	"github.com/rodziievskyi-maksym/tatl-test-app/internal/user/applicaton"
	userrepository "github.com/rodziievskyi-maksym/tatl-test-app/internal/user/infrastructure/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

const MysqlDsnPattern = "%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	fmt.Println("tatl-test-app")

	//I'm explicitly avoid config file due to task assignment.
	address := flag.String("port", ":8080", "Application port")
	dbUser := flag.String("dbuser", "", "Database User")
	dbPass := flag.String("dbpass", "", "Database Password")
	dbHost := flag.String("dbhost", "127.0.0.1:3306", "Database Host")
	dbName := flag.String("dbname", "", "Database Name")

	flag.Parse()

	var dsn string
	if *dbUser != "" && *dbPass != "" && *dbHost != "" && *dbName != "" {
		dsn = fmt.Sprintf(MysqlDsnPattern, *dbUser, *dbPass, *dbHost, *dbName)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to establish database connection %s", err.Error())
	}

	userRepository := userrepository.NewUserRepository(db)
	userController := applicaton.NewUserController(userRepository)

	authRepository := authrepository.NewAuthRepository(db)

	httpServer := api.NewHTTPServer(*address, userController, authRepository)

	if err = httpServer.Start(); err != nil {
		log.Fatalf("Failed to start HTTP server: %s", err.Error())
	}
}
