package api

import (
	_ "encoding/json"
	"net/http"

	"github.com/lopz/cs-api-test/internal/database"
	"github.com/lopz/cs-api-test/internal/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/google/uuid"
)

func createPerson(w http.ResponseWriter, r *http.Request) {
	person := models.Person{}
	person.UUID = uuid.New().String()

	if err := render.DecodeJSON(r.Body, &person); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err)
		return
	}

	database.DbAddPerson(person)

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, person)

}

func getAllPerson(w http.ResponseWriter, r *http.Request) {
	//person := []models.Person{}
	result, _ := database.DbGetAllPerson()

	render.Status(r, http.StatusOK)
	render.JSON(w, r, result)
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")
	result, _ := database.DbGetPerson(uuid)
	render.Status(r, http.StatusOK)
	render.JSON(w, r, result)
}

func updatePerson(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")
	person := models.Person{}

	if err := render.DecodeJSON(r.Body, &person); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err)
		return
	}
	person, _ = database.DbUpdatePerson(uuid, person)
	render.Status(r, http.StatusOK)
	render.JSON(w, r, person)
}

func deletePerson(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")
	result, _ := database.DbDeletePerson(uuid)
	//render.Status(r, http.StatusNoContent)
	render.JSON(w, r, result)
}
