package models

//NewCompany is the model for the companies table for insteting a new row
type NewCompany struct {
	Name                string `json:"name" form:"name"`
	CompanyTypeID       int64  `json:"company_type_id" form:"company_type_id"`
	ContactID           int64  `json:"contact_id" form:"contact_id"`
	BankDetailID        int64  `json:"bank_detail_id" form:"bank_detail_id"`
	Currency            string `json:"currency" form:"currency"`
	DateOfEstablishment string `json:"date_of_establishment" form:"date_of_establishment"`
	ParentID            int64  `json:"parent_id" form:"parent_id"`
	IsSisterCompany     bool   `json:"is_sister_company" form:"is_sister_company"`
	LocalName           string `json:"local_name" form:"local_name"`
	Code                string `json:"code" form:"code"`
	IsPreferred         bool   `json:"is_preferred" form:"is_preferred"`
	ServiceTypeID       int64  `json:"service_type_id" form:"service_type_id"`
}
