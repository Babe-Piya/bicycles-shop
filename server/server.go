package server

import (
	"bicycles-shop/endpoint/bicycles"
	"bicycles-shop/model"
	"bicycles-shop/repository"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const defaultPort = "8080"

func Run() {
	handler, err := setUpServer()
	if err != nil {
		log.Printf("Error setUpServer ---> %+v", err.Error())
	}

	srv := &http.Server{Addr: ":" + defaultPort, Handler: handler}

	go func() {
		log.Printf("Listening and serving HTTP on %s\n", defaultPort)
		err := srv.ListenAndServe()
		if err != nil {
			if err == http.ErrServerClosed {
				log.Printf("Error Server shut down. Waiting for connections to drain. --> %+v", err)
			} else {
				log.Printf("Error failed to start server on port: %+v", srv.Addr)
			}
		}
	}()

	// Wait for an interrupt
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)    // interrupt signal sent from terminal
	signal.Notify(sigint, syscall.SIGTERM) // sigterm signal sent from system
	<-sigint
}

func setUpServer() (http.Handler, error) {
	db, err := setUpDatabase()
	if err != nil {
		return nil, fmt.Errorf("cannot setup database %w", err)
	}

	initLogger()

	//repo
	bicyclesRepo := repository.New(db)

	//service
	bicyclesService := bicycles.NewService(bicyclesRepo)

	router := mux.NewRouter()

	// consumer API.
	consumerSubRouter := router.PathPrefix("/").Subrouter()

	// subscription route.
	subscriptionSubRouter := consumerSubRouter.PathPrefix("/bicycle").Subrouter()
	subscriptionSubRouter.Handle("/get-list", bicycles.MakeGetListBicyclesHandler(bicyclesService)).Methods(http.MethodGet)
	subscriptionSubRouter.Handle("/buy", bicycles.MakeBuyBicycleHandler(bicyclesService)).Methods(http.MethodPut)
	subscriptionSubRouter.Handle("/create", bicycles.MakeCreateBicycleHandler(bicyclesService)).Methods(http.MethodPost)

	return router, nil
}

func setUpDatabase() (*gorm.DB, error) {
	dsn := model.DBConnection()
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func initLogger() {
	var (
		logger *zap.Logger
		err    error
	)
	logger, err = zap.NewDevelopment()

	if err != nil {
		log.Fatalln(err)
	}

	zap.ReplaceGlobals(logger)
}
