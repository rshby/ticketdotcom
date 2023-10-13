package entity

type City struct {
	Id         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	ProvinceId int    `json:"province_id,omitempty"`
}
