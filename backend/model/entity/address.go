package entity

type Address struct {
	Id            int    `json:"id,omitempty"`
	Street        string `json:"street,omitempty"`
	SubDistrictId string `json:"sub_district_id,omitempty"`
}
