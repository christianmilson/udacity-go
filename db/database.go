package db

import (
	"github.com/christianmilson/udacity-go/models"
)

type Database struct {
	records map[int]models.Customer
	lastId  int
}

func NewDb() *Database {
	return &Database{
		lastId:  0,
		records: make(map[int]models.Customer),
	}
}

func (db Database) FetchAll() map[int]models.Customer {
	return db.records
}

func (db Database) FindById(id int) models.Customer {
	return db.records[id]
}

func (db *Database) Seed(numRecords int, model *models.Customer) {
	// For each item, generate a faker, and then insert into db
	if numRecords > 0 {
		for i := 0; i < numRecords; i++ {
			db.Insert(model.Faker())
		}
	}
}

func (db *Database) Insert(record models.Customer) {
	// Increment the last id
	db.lastId++

	// Update record id, to the new id
	record.Id = db.lastId

	// Insert the record
	db.records[db.lastId] = record
}

func (db *Database) Update(record models.Customer) {
	// Update the record, using its key => value
	db.records[record.Id] = record
}

func (db *Database) Delete(id int) {
	// Remove the record from records list
	delete(db.records, id)
}
