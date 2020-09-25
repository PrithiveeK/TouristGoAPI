package models

//Service is the model for the services table in the database
type Service struct {
	ID              int64      `json:"id"`
	Code            string     `json:"code"`
	Name            string     `json:"name"`
	CountryID       NullInt64  `json:"country_id"`
	CityID          NullInt64  `json:"city_id"`
	CountryName     NullString `json:"country_name"`
	CityName        NullString `json:"city_name"`
	ServiceTypeID   NullInt64  `json:"service_type_id"`
	ServiceTypeName NullString `json:"service_type_name"`
}

//NewService is the model for inserting a new row of service into
//services table in the database
type NewService struct {
	ServiceTypeID   int64  `json:"service_type_id" form:"service_type_id"`
	Name            string `json:"name" form:"name"`
	CountryID       int64  `json:"country_id" form:"country_id"`
	CityID          int64  `json:"city_id" form:"city_id"`
	Zipcode         string `json:"zipcode" from:"zipcode"`
	Fax             string `json:"fax" from:"fax"`
	PhoneNO         string `json:"phone_no" form:"phone_no"`
	TelephoneNo     string `json:"telephone_no" from:"telephone_no"`
	Email           string `json:"email" form:"email"`
	Website         string `json:"website" from:"website"`
	Street          string `json:"street" from:"street"`
	Description     string `json:"description" from:"description"`
	ChainID         int64  `json:"chain_id" form:"chain_id"`
	StarRating      string `json:"rating" form:"rating"`
	LocationTypeID  int64  `json:"location_type_id" form:"location_type_id"`
	Location2TypeID int64  `json:"location2_type_id" form:"location2_type_id"`
	NoOfRooms       string `json:"no_of_rooms" form:"no_of_rooms"`
	Seats           string `json:"seats" form:"seats"`
	StyleID         int64  `json:"style_id" form:"style_id"`
	IsUNESCO        bool   `json:"is_unesco" form:"is_unesco"`
	HasAC           bool   `json:"has_ac" form:"has_ac"`
	IsPreferred     bool   `json:"is_preferred" form:"is_preferred"`
	IsLicensed      bool   `json:"is_licensed" form:"is_licensed"`
	Code            string `json:"code" form:"code"`
	Placeholder     bool   `json:"placeholder" form:"placeholder"`
	IsHotel         bool   `json:"is_hotel" form:"is_hotel"`
	Is99AService    bool   `json:"is_99A_service" form:"is_99A_service"`
}

//NewPriceDetails is the model for inserting a new row into
//services_price_details table in the database
type NewPriceDetails struct {
	ServiceID     int64  `json:"service_id" form:"service_id"`
	Currency      string `json:"currency" form:"currency"`
	RoomType      int64  `json:"room_type_id" form:"room_type_id"`
	PersonType    int64  `json:"person_type_id" form:"person_type_id"`
	Pricing       string `json:"pricing_type" form:"pricing_type"`
	Charge        int64  `json:"charge_type_id" form:"charge_type_id"`
	MealType      int64  `json:"menu_type_id" form:"menu_type_id"`
	BreakfastType int64  `json:"breakfast_type_id" form:"breakfast_type_id"`
	Description   string `json:"description" form:"description"`
	Type          string `json:"type" form:"type"`
	Category      int64  `json:"category_type_id" form:"category_type_id"`
	Price         int64  `json:"price" form:"price"`
}
