package dto

type CreateCustumerRequest struct {
	Name      string `json:"name" form:"name" binding:"required,min=1"`
	Email     string `json:"email" form:"email" binding:"required"`
	Telephone string `json:"telephone" form:"telephone" binding:"required"`
}

type UpdateCustumerRequest struct {
	ID        int64  `json:"id" form:"id"`
	Name      string `json:"name" form:"name" binding:"required,min=1"`
	Email     string `json:"email" form:"email" binding:"required"`
	Telephone string `json:"telephone" form:"telephone" binding:"required"`
}
