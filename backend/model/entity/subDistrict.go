package entity

type SubDistrict struct {
	Id         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	PostalCode string `json:"postal_code,omitempty"`
	DistrictId int    `json:"district_id,omitempty"`
}
