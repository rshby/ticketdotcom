package dto

type InsertCityRequest struct {
	Name       string `json:"name" binding:"required"`
	ProvinceId int    `json:"province_id" binding:"required,gt=0"`
}
