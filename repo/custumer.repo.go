package repo

import (
	"github.com/thpacheco/golang_heroku/entity"
	"gorm.io/gorm"
)

type CustumerRepository interface {
	All() ([]entity.Custumer, error)
	InsertCustumer(custumer entity.Custumer) (entity.Custumer, error)
	UpdateCustumer(custumer entity.Custumer) (entity.Custumer, error)
	DeleteCustumer(custumerID string) error
	FindOneCustumerByID(ID string) (entity.Custumer, error)
	FindAllCustumer(custumerID string) ([]entity.Custumer, error)
}

type custumerRepo struct {
	connection *gorm.DB
}

func NewCustumerRepo(connection *gorm.DB) CustumerRepository {
	return &custumerRepo{
		connection: connection,
	}
}

func (c *custumerRepo) All() ([]entity.Custumer, error) {
	custumers := []entity.Custumer{}
	c.connection.Preload("Custumers").Find(&custumers).Find(&custumers)
	return custumers, nil
}

func (c *custumerRepo) InsertCustumer(custumer entity.Custumer) (entity.Custumer, error) {
	c.connection.Save(&custumer)
	c.connection.Preload("Custumers").Find(&custumer)
	return custumer, nil
}

func (c *custumerRepo) UpdateCustumer(custumer entity.Custumer) (entity.Custumer, error) {
	c.connection.Save(&custumer)
	c.connection.Preload("Custumers").Find(&custumer)
	return custumer, nil
}

func (c *custumerRepo) FindOneCustumerByID(custumerID string) (entity.Custumer, error) {
	var custumer entity.Custumer
	res := c.connection.Preload("Custumers").Where("id = ?", custumerID).Take(&custumer)
	if res.Error != nil {
		return custumer, res.Error
	}
	return custumer, nil
}

func (c *custumerRepo) FindAllCustumer(custumerID string) ([]entity.Custumer, error) {
	custumers := []entity.Custumer{}
	c.connection.Where("user_id = ?", custumerID).Find(&custumers)
	return custumers, nil
}

func (c *custumerRepo) DeleteCustumer(custumerID string) error {
	var custumer entity.Custumer
	res := c.connection.Preload("Custumers").Where("id = ?", custumerID).Take(&custumer)
	if res.Error != nil {
		return res.Error
	}
	c.connection.Delete(&custumer)
	return nil
}
