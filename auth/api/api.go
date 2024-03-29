package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alekssro/banking/auth/application"
	"github.com/alekssro/banking/auth/domain/entity"
	"github.com/alekssro/banking/auth/domain/repository"
	"github.com/alekssro/banking/lib/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func Start() {

	sanityCheck()
	router := mux.NewRouter()
	authRepository := repository.NewAuthRepository(getDbClient())
	ah := AuthHandler{application.NewLoginService(authRepository, entity.GetRolePermissions())}

	router.HandleFunc("/auth/login", ah.Login).Methods(http.MethodPost)
	router.HandleFunc("/auth/register", ah.NotImplementedHandler).Methods(http.MethodPost)
	router.HandleFunc("/auth/refresh", ah.Refresh).Methods(http.MethodPost)
	router.HandleFunc("/auth/verify", ah.Verify).Methods(http.MethodGet)

	address := os.Getenv("AUTH_HOST")
	port := os.Getenv("AUTH_PORT")
	logger.Info(fmt.Sprintf("Starting OAuth server on %s:%s ...", address, port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

func getDbClient() *sqlx.DB {
	dbAddr := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}

func sanityCheck() {
	envProps := []string{
		"AUTH_HOST",
		"AUTH_PORT",
		"DB_HOST",
		"DB_PORT",
		"DB_USER",
		"DB_PASS",
		"DB_NAME",
	}
	for _, k := range envProps {
		if os.Getenv(k) == "" {
			logger.Error(fmt.Sprintf("Environment variable %s not defined. Terminating application...", k))
		}
	}
}
