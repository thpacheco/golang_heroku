package _custumer

import (
	"github.com/thpacheco/golang_heroku/entity"
)

type CustumerResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
	DateStart string `json:"dateStart"`
}

func NewCustumerResponse(custumer entity.Custumer) CustumerResponse {
	return CustumerResponse{
		ID:        custumer.ID,
		Name:      custumer.Name,
		Email:     custumer.Email,
		Telephone: custumer.Telephone,
		DateStart: custumer.DataStart,
	}
}

func NewCustumerArrayResponse(custumers []entity.Custumer) []CustumerResponse {
	custumerRes := []CustumerResponse{}
	for _, v := range custumers {
		p := CustumerResponse{
			ID:        v.ID,
			Name:      v.Name,
			Email:     v.Email,
			Telephone: v.Telephone,
		}
		custumerRes = append(custumerRes, p)
	}
	return custumerRes
}
