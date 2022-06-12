package dto

type CreateCustumerRequest struct {
	Nome       string `json:"nome" form:"nome" binding:"required,min=1"`
	Email      string `json:"email" form:"email" binding:"required"`
	Celular    string `json:"celular" form:"celular" binding:"required"`
	DataInicio string `json:"dataInicio" form:"dataInicio" binding:"required"`
}

type UpdateCustumerRequest struct {
	ID      int64  `json:"id" form:"id"`
	Nome    string `json:"nome" form:"nome" binding:"required,min=1"`
	Email   string `json:"email" form:"email" binding:"required"`
	Celular string `json:"celular" form:"celular" binding:"required"`
}
