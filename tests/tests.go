package tests

import (
	. "restapi-sample/controllers"
	//. "restapi-sample/models"
	. "restapi-sample/repositories"
	"testing"
)

func TestAllEntities(t *testing.T) {
	var EntitiesController = &Controller{Repository: EntitiesRepository{}}
	EntitiesController.Init()
}
