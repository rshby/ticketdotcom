package dto

type GenderRequest struct {
	Code string `json:"code,omitempty" binding:"max=1,required"`
	Name string `json:"name,omitempty" binding:"required"`
}
