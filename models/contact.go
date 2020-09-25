package models

//NewContact is the model for inserting a new contact row to the database
type NewContact struct {
	CountryID   int64  `json:"country_id" form:"country_id"`
	CityID      int64  `json:"city_id" form:"city_id"`
	Street      string `json:"street" form:"street"`
	Zipcode     string `json:"zipcode" form:"zipcode"`
	Website     string `json:"website" form:"website"`
	PhoneNo     string `json:"phone_no" form:"phone_no"`
	TelephoneNo string `json:"telephone_no" form:"telephone_no"`
	Fax         string `json:"fax" form:"fax"`
	Email       string `json:"email" form:"email"`
	SkypeID     string `json:"skype_id" form:"skype_id"`
}
