package dto

import "backend/model/entity"

type CityDetail[T any] struct {
	Province *entity.Province `json:"province,omitempty"`
	Cities   T                `json:"cities,omitempty"`
}
