package server

import (
	authHttp "GOSecretProject/core/auth/delivery/http"
	authInterfaces "GOSecretProject/core/auth/interfaces"
	authRepository "GOSecretProject/core/auth/repository/postgres"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kataras/golog"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

type App struct {
	authRepo authInterfaces.AuthRepository
}

func NewApp() *App {
	log.Print(os.Getenv("HAHA_DB_USER"), " !!! ", os.Getenv("HAHA_DB_PASSWORD"))
	dbinfo := fmt.Sprintf("host=127.0.0.1 user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("HAHA_DB_USER"), os.Getenv("HAHA_DB_PASSWORD"), "ios")
	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		golog.Error(err.Error())
		return nil
	}
	err = db.Ping()
	if err != nil {
		golog.Error("DB: ", err.Error())
		return nil
	}

	authRepo := authRepository.NewAuthRepository(db)

	return &App{
		authRepo: authRepo,
	}
}

func (app *App) StartRouter() {
	router := mux.NewRouter()

	commonRouter := router.PathPrefix("/api").Subrouter()

	authHttp.RegisterHTTPEndpoints(commonRouter, app.authRepo)

	http.Handle("/", router)

	port := 9000
	golog.Infof("Server started at port :%d", port)
	err := http.ListenAndServeTLS(fmt.Sprintf(":%d", port),
		"/etc/letsencrypt/live/ios.hahao.ru/fullchain.pem",
		"/etc/letsencrypt/live/ios.hahao.ru/privkey.pem",
		nil)

	if err != nil {
		golog.Error("Server ios.haha failed: ", err)
	}
}
