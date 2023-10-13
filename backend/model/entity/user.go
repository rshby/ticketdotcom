package entity

type User struct {
	Id           int    `json:"id,omitempty"`
	FullName     string `json:"full_name,omitempty"`
	Phone        string `json:"phone,omitempty"`
	GenderId     int    `json:"gender_id,omitempty"`
	AddressId    int    `json:"address_id,omitempty"`
	EmailAccount string `json:"email_account,omitempty"`
}
