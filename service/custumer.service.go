package service

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/mashingan/smapping"
	"github.com/thpacheco/golang_heroku/dto"
	"github.com/thpacheco/golang_heroku/entity"
	"github.com/thpacheco/golang_heroku/repo"
	_custumer "github.com/thpacheco/golang_heroku/service/custumer"
)

type CustumerService interface {
	All(userID string) (*[]_custumer.CustumerResponse, error)
	CreateCustumer(custumerRequest dto.CreateCustumerRequest, userID string) (*_custumer.CustumerResponse, error)
	UpdateCustumer(updateCustumerRequest dto.UpdateCustumerRequest, userID string) (*_custumer.CustumerResponse, error)
	FindOneCustumerByID(custumerID string) (*_custumer.CustumerResponse, error)
	DeleteCustumer(custumerID string, ID string) error
}

type custumerService struct {
	custumerRepo repo.CustumerRepository
}

func NewCustumerService(custumerRepo repo.CustumerRepository) CustumerService {
	return &custumerService{
		custumerRepo: custumerRepo,
	}
}

func (c *custumerService) All(custumerID string) (*[]_custumer.CustumerResponse, error) {
	custumers, err := c.custumerRepo.All(custumerID)
	if err != nil {
		return nil, err
	}

	custms := _custumer.NewCustumerArrayResponse(custumers)
	return &custms, nil
}

func (c *custumerService) CreateCustumer(custumerRequest dto.CreateCustumerRequest, userID string) (*_custumer.CustumerResponse, error) {
	custumer := entity.Custumer{}
	err := smapping.FillStruct(&custumer, smapping.MapFields(&custumerRequest))

	if err != nil {
		log.Fatalf("Failed map %v", err)
		return nil, err
	}

	id, _ := strconv.ParseInt(userID, 0, 64)
	custumer.ID = id
	p, err := c.custumerRepo.InsertCustumer(custumer)
	if err != nil {
		return nil, err
	}

	res := _custumer.NewCustumerResponse(p)
	return &res, nil
}

func (c *custumerService) FindOneCustumerByID(custumerID string) (*_custumer.CustumerResponse, error) {
	custumer, err := c.custumerRepo.FindOneCustumerByID(custumerID)

	if err != nil {
		return nil, err
	}

	res := _custumer.NewCustumerResponse(custumer)
	return &res, nil
}

func (c *custumerService) UpdateCustumer(updatecustumerRequest dto.UpdateCustumerRequest, custumerID string) (*_custumer.CustumerResponse, error) {
	custumer, err := c.custumerRepo.FindOneCustumerByID(fmt.Sprintf("%d", updatecustumerRequest.ID))
	if err != nil {
		return nil, err
	}

	uid, _ := strconv.ParseInt(custumerID, 0, 64)
	if custumer.ID != uid {
		return nil, errors.New("produk ini bukan milik anda")
	}

	custumer = entity.Custumer{}
	err = smapping.FillStruct(&custumer, smapping.MapFields(&updatecustumerRequest))

	if err != nil {
		return nil, err
	}

	custumer.ID = uid
	custumer, err = c.custumerRepo.UpdateCustumer(custumer)

	if err != nil {
		return nil, err
	}

	res := _custumer.NewCustumerResponse(custumer)
	return &res, nil
}

func (c *custumerService) DeleteCustumer(custumerID string, ID string) error {
	custumer, err := c.custumerRepo.FindOneCustumerByID(custumerID)
	if err != nil {
		return err
	}

	if fmt.Sprintf("%d", custumer.ID) != ID {
		return errors.New("produk ini bukan milik anda")
	}

	c.custumerRepo.DeleteCustumer(custumerID)
	return nil

}
