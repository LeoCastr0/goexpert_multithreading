package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
	Errors      error  `json:"errors"`
}

type BrasilAPI struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Errors       error  `json:"errors"`
}

func main() {
	http.HandleFunc("/", FindZipHandler)
	http.ListenAndServe(":8080", nil)
}

func FindZipHandler(w http.ResponseWriter, r *http.Request) {
	chanVC := make(chan ViaCEP)
	chanBA := make(chan BrasilAPI)
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	zipParam := r.URL.Query().Get("zip")
	if zipParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	go SearchZipViaCEP(zipParam, chanVC)
	go SearchZipBrasilAPI(zipParam, chanBA)

	select {
	case zipInfo := <-chanVC:
		log.Printf("Brasil API responsed First.")
		SetHttpHeaders(w)
		json.NewEncoder(w).Encode(zipInfo)

	case zipInfo := <-chanBA:
		log.Printf("Brasil API responsed First.")
		SetHttpHeaders(w)
		json.NewEncoder(w).Encode(zipInfo)

	case <-time.After(time.Second):
		log.Printf("The APIs didn't response on time.")
		w.WriteHeader(http.StatusRequestTimeout)
	}
}

func SetHttpHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func SearchZipViaCEP(zip string, ch chan ViaCEP) {
	res, err := http.Get("https://viacep.com.br/ws/" + zip + "/json")
	if err != nil {
		ch <- ViaCEP{Errors: err}
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		ch <- ViaCEP{Errors: err}
	}
	var c ViaCEP
	err = json.Unmarshal(body, &c)
	if err != nil {
		ch <- ViaCEP{Errors: err}
	}
	ch <- c
}

func SearchZipBrasilAPI(zip string, ch chan BrasilAPI) {
	res, err := http.Get("https://brasilapi.com.br/api/cep/v1/" + zip)
	if err != nil {
		ch <- BrasilAPI{Errors: err}
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		ch <- BrasilAPI{Errors: err}
	}
	var c BrasilAPI
	err = json.Unmarshal(body, &c)
	if err != nil {
		ch <- BrasilAPI{Errors: err}
	}
	ch <- c
}
