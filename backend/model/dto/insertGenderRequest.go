package dto

type GenderRequest struct {
	Code string `json:"code" binding:"max=1,required"`
	Name string `json:"name" binding:"required"`
}
