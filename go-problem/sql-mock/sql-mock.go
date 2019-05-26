package main

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Repository interface {
	Get(id uuid.UUID) (*Person, error)
	Create(id uuid.UUID, name string) error
}

type repo struct {
	DB *gorm.DB
}

func (p *repo) Create(id uuid.UUID, name string) error {
	person := &Person{
		ID:   id,
		Name: name,
	}

	return p.DB.Create(person).Error
}

func (p *repo) Get(id uuid.UUID) (*Person, error) {
	person := new(Person)

	err := p.DB.Where("id = ?", id).Find(person).Error

	return person, err
}

func CreateRepository(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}

func main() {
	println("Hello world")
}
