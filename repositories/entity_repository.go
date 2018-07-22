package repositories

import (
	"log"
	. "restapi-sample/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// EntitiesRepository is data access object
type EntitiesRepository struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	// COLLECTION name on mongodb database
	COLLECTION = "entities"
)

// init is called after all the variable declarations in the package have evaluated their initializers,
// and those are evaluated only after all the imported packages have been initialized.
// import --> const --> var --> init()
func (repository *EntitiesRepository) Init() {
	repository.Server = "localhost:27017"
	repository.Database = "darwin"
	repository.Connect()
}

// Connect establishs a connection to database
func (repository *EntitiesRepository) Connect() {
	session, err := mgo.Dial(repository.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(repository.Database)
}

// FindAll returns the list of all entities
func (repository *EntitiesRepository) FindAll() ([]Entity, error) {
	var entities []Entity
	err := db.C(COLLECTION).Find(bson.M{}).All(&entities)
	return entities, err
}

// FindByID returns an entity by its id
func (repository *EntitiesRepository) FindByID(id string) (Entity, error) {
	var entity Entity
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&entity)
	return entity, err
}

// Insert an entity into database
func (repository *EntitiesRepository) Insert(entity Entity) error {
	err := db.C(COLLECTION).Insert(&entity)
	return err
}

// Delete an existing entity
func (repository *EntitiesRepository) Delete(entity Entity) error {
	err := db.C(COLLECTION).Remove(&entity)
	return err
}

// Update an existing entity
func (repository *EntitiesRepository) Update(entity Entity) error {
	err := db.C(COLLECTION).UpdateId(entity.ID, &entity)
	return err
}
