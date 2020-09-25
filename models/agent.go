package models

//Agent in the model for agents table
type Agent struct {
	ID          int64      `json:"id"`
	Code        string     `json:"code"`
	Name        string     `json:"name"`
	CityID      NullInt64  `json:"city_id"`
	CountryID   NullInt64  `json:"country_id"`
	CityName    NullString `json:"city_name"`
	CountryName NullString `json:"country_name"`
}

//NewAgent is the model for inserting a new agent row to the database
type NewAgent struct {
	Agent   NewCompany `json:"agent" form:"agent"`
	Contact NewContact `json:"contact" form:"contact"`
}
