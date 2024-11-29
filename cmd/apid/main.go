package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/douglasschantz/cep2temp/cmd/apid/config"
	"github.com/douglasschantz/cep2temp/cmd/apid/router"
	"github.com/douglasschantz/cep2temp/internal/cep2temp"
	"github.com/douglasschantz/cep2temp/internal/cep2temp/cep"
	"github.com/douglasschantz/cep2temp/internal/cep2temp/weather"
	platHttp "github.com/douglasschantz/cep2temp/internal/platform/http"
)

func main() {
	cfg, err := config.LoadConfig("./configs")
	if err != nil {
		panic(fmt.Sprintf("error starting configs: %s", err.Error()))
	}

	httpClient := platHttp.NewDefaultClient()
	cepService := cep.NewService(httpClient, cfg.CepAPIConfig())
	weatherService := weather.NewService(httpClient, cfg.WeatherAPIConfig())
	service := cep2temp.NewService(cepService, weatherService)

	handler := cep2temp.NewHandler(service)

	r := router.SetupRouter(handler)

	log.Println("server running on port 8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		panic(fmt.Sprintf("error on listen and serve: %s", err.Error()))
	}
}
