package entity

type District struct {
	Id     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	CityId int    `json:"city_id,omitempty"`
}
