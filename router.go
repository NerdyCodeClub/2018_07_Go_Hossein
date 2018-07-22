package main

import (
	"log"
	"net/http"
	. "restapi-sample/controllers"
	. "restapi-sample/repositories"

	"github.com/gorilla/mux"
)

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the list of routes of our API
type Routes []Route

var EntitiesController = &Controller{Repository: EntitiesRepository{}}

var routes = Routes{
	Route{
		"GetEntities",
		"GET",
		"/entities",
		EntitiesController.AllEntities,
	},
	Route{
		"GetEntity",
		"GET",
		"/entities/{id}",
		EntitiesController.FindEntity,
	},
	Route{
		"CreateNewEntity",
		"POST",
		"/entities",
		EntitiesController.CreateNewEntity,
	},
	Route{
		"UpdateEntity",
		"PUT",
		"/entities",
		EntitiesController.UpdateEntity,
	},
	Route{
		"DeleteEntity",
		"DELETE",
		"/entities",
		EntitiesController.DeleteEntity,
	},
}

// NewRouter configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		log.Println(route.Name)
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
