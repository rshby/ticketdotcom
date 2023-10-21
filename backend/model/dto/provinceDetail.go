package dto

import "backend/model/entity"

type ProvinceDetail struct {
	Id     int           `json:"id,omitempty"`
	Name   string        `json:"name,omitempty"`
	Cities []entity.City `json:"cities"`
}
