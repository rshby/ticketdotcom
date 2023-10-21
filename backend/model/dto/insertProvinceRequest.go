package dto

type InsertProvinceRequest struct {
	Name string `json:"name" binding:"required"`
}
