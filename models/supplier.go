package models

//Supplier is the model for suppliers table in the database
type Supplier struct {
	ID              int64      `json:"id"`
	Name            string     `json:"name"`
	Code            string     `json:"code"`
	IsPreferred     bool       `json:"is_preferred"`
	CityID          NullInt64  `json:"city_id"`
	CountryID       NullInt64  `json:"country_id"`
	CityName        NullString `json:"city_name"`
	CountryName     NullString `json:"country_name"`
	Email           NullString `json:"email"`
	PhoneNo         NullString `json:"phone_no"`
	ServiceTypeID   NullInt64  `json:"service_type_id"`
	ServiceTypeName NullString `json:"service_type_name"`
}

//NewSupplier is the model for inserting a new supplier
type NewSupplier struct {
	Supplier NewCompany    `json:"supplier" form:"supplier"`
	Contact  NewContact    `json:"contact" form:"contact"`
	Bank     NewBankDetail `json:"bank" form:"bank"`
}
