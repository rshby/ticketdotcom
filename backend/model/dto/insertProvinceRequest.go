package dto

type InsertProvinceRequest struct {
	Name string `json:"name,omitempty" binding:"required"`
}
