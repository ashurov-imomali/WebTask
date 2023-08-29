package main

import (
	"encoding/json"
	"errors"
	"log"
	"main/internal/db"
	"main/internal/handlers"
	"main/internal/router"
	"main/internal/service"
	"main/pkg/models"
	"net/http"
	"os"
)

func main() {
	err := start()
	if err != nil {
		log.Println(err)
		return
	}
}

func start() error {
	adr, err := GetConfigs("./configs/configs.json")
	if err != nil {
		return errors.New("Internal error ")
	}
	conn, err := db.ConnectionToDb("./configs/db_settings.json")
	if err != nil {
		return errors.New("Internal error ")
	}
	srv := service.NewService(conn)
	handler := handlers.NewHandler(srv)
	NewRouter := router.NewRouter(handler)
	server := http.Server{
		Addr:    adr.Host + adr.Port,
		Handler: NewRouter,
	}
	err = server.ListenAndServe()
	if err != nil {
		return errors.New("Internal error ")
	}
	return nil
}

func GetConfigs(path string) (*models.Address, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var address models.Address
	err = json.Unmarshal(bytes, &address)
	if err != nil {
		return nil, err
	}
	return &address, nil
}
