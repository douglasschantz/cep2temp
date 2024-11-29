package cep2temp

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"github.com/douglasschantz/cep2temp/internal/cep2temp/cep"
	"github.com/douglasschantz/cep2temp/internal/cep2temp/weather"
	"github.com/douglasschantz/cep2temp/internal/platform/errors"
)

const (
	ErrRequired = "ERR_CEP_REQUIRED"
	ErrInvalid  = "ERR_CEP_INVALID"
)

type Request struct {
	CEP string `json:"cep"`
}

type Response struct {
	TempCelsius    float64 `json:"temp_C"`
	TempFahrenheit float64 `json:"temp_F"`
	TempKelvin     float64 `json:"temp_K"`
}

func (r *Response) ToJSON() []byte {
	jsonBytes, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return jsonBytes
}

func NewRequest(cep string) Request {
	return Request{
		CEP: cep,
	}
}

func (r *Request) Validate() error {
	if r.CEP == "" {
		return errors.NewUnprocessableError(ErrRequired)
	}

	r.CEP = strings.ReplaceAll(r.CEP, "-", "")
	if len(r.CEP) != 8 {
		return errors.NewUnprocessableError(ErrInvalid)
	}

	return nil
}

func (r *Request) BuildCEPRequest() cep.Request {
	return cep.Request{Cep: r.CEP}
}

func NewWeatherRequest(response cep.Response) weather.Request {
	query := fmt.Sprintf("%s,%s", response.City, response.State)
	return weather.Request{Query: query}
}

func NewResponse(resp weather.Response) Response {
	return Response{
		TempCelsius:    resp.Current.TempCelsius,
		TempFahrenheit: resp.Current.TempFahrenheit,
		TempKelvin:     math.Round(resp.Current.TempCelsius+273.15*100) / 100,
	}
}
