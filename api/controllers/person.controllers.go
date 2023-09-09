package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/IsaacDSC/rinha-backend-dsc/internal/models"
	"github.com/IsaacDSC/rinha-backend-dsc/shared/di"
	"github.com/IsaacDSC/rinha-backend-dsc/shared/dto"

	"github.com/go-chi/chi/v5"
)

type PersonController struct{}

func (p *PersonController) Start(router *chi.Mux) {
	router.Post("/pessoas", p.createPerson)
	router.Get("/pessoas/{id}", p.retrievePerson)
	router.Get("/pessoas", p.searchPerson)
	router.Get("/contagem-pessoas", p.totalPerson)
}

func (*PersonController) createPerson(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, _ := io.ReadAll(r.Body)
	var person dto.Person
	err := json.Unmarshal(body, &person)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	if len(person.Name) == 0 || len(person.LastName) == 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	service := di.GetInstanceService(models.Person(person))
	output, err := service.CreatePerson(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Add("Location", "/pessoas/"+output.ID)
	w.WriteHeader(http.StatusCreated)
}

func (*PersonController) retrievePerson(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "id")
	service := di.GetInstanceService(models.Person{
		ID: ID,
	})
	person, err := service.RetrievePerson(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	output, err := json.Marshal(person)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}

func (*PersonController) searchPerson(w http.ResponseWriter, r *http.Request) {
	query_params := r.URL.Query()
	service := di.GetInstanceService(models.Person{
		Name:     query_params.Get("nome"),
		LastName: query_params.Get("apelido"),
		Birthday: query_params.Get("nascimento"),
		Stack:    []string{query_params.Get("stack")},
	})
	persons, err := service.SearchPerson(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	output, _ := json.Marshal(persons)
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}

func (*PersonController) totalPerson(w http.ResponseWriter, r *http.Request) {
	service := di.GetInstanceService(models.Person{})
	totalPerson := service.TotalPerson(r.Context())
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("{\"total\": \"%d\"}", totalPerson)))
}
