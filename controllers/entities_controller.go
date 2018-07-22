package controllers

import (
	"encoding/json"
	"net/http"
	. "restapi-sample/models"
	. "restapi-sample/repositories"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

//Controller ...
type Controller struct {
	Repository EntitiesRepository
}

func (c *Controller) Init() {
	c.Repository = EntitiesRepository{}
	c.Repository.Init()
}

// AllEntities endpoint
func (c *Controller) AllEntities(w http.ResponseWriter, r *http.Request) {
	entities, err := c.Repository.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, entities)
}

// FindEntity endpoint
func (c *Controller) FindEntity(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	entity, err := c.Repository.FindByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Entity ID")
		return
	}
	respondWithJSON(w, http.StatusOK, entity)
}

// CreateNewEntity endpoint
func (c *Controller) CreateNewEntity(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var entity Entity
	if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	entity.ID = bson.NewObjectId()
	if err := c.Repository.Insert(entity); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, entity)
}

// UpdateEntity endpoint
func (c *Controller) UpdateEntity(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var entity Entity
	if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := c.Repository.Update(entity); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

// DeleteEntity endpoint
func (c *Controller) DeleteEntity(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var entity Entity
	if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := c.Repository.Delete(entity); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
