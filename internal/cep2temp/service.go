package cep2temp

import (
	"context"
	"log"

	"github.com/douglasschantz/cep2temp/internal/cep2temp/cep"
	"github.com/douglasschantz/cep2temp/internal/cep2temp/weather"
)

type Service interface {
	GetTemperatureByCEP(ctx context.Context, request Request) (Response, error)
}

type service struct {
	weatherService weather.Service
	cepService     cep.Service
}

func (s service) GetTemperatureByCEP(ctx context.Context, request Request) (Response, error) {
	log.Printf("starting cep2temp for cep %s\n", request.CEP)
	if err := request.Validate(); err != nil {
		log.Printf("error validating request: %s\n", err.Error())
		return Response{}, err
	}

	cepRequest := request.BuildCEPRequest()
	cepResponse, err := s.cepService.GetInfo(ctx, cepRequest)
	if err != nil {
		log.Printf("error getting cep info: %v\n", err)
		return Response{}, err
	}

	weatherRequest := NewWeatherRequest(cepResponse)
	weatherResponse, err := s.weatherService.GetInfo(ctx, weatherRequest)
	if err != nil {
		log.Printf("error getting weather info: %v\n", err)
		return Response{}, err
	}

	resp := NewResponse(weatherResponse)
	log.Printf("finished cep2temp for cep %s | temp_c: %.2f, | temp_f: %.2f, | temp_k: %.2f\n",
		request.CEP, resp.TempCelsius, resp.TempFahrenheit, resp.TempKelvin)
	return resp, nil
}

func NewService(cepService cep.Service, weatherService weather.Service) Service {
	return &service{
		weatherService: weatherService,
		cepService:     cepService,
	}
}
