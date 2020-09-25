package controllers

import (
	"log"
	"time"

	dbMod "touristapp.com/db"
	Models "touristapp.com/models"
)

//AddCompany inserts a new company into the database
func AddCompany(newCompany Models.NewCompany) (int64, error) {
	queryNewCompany := `
		INSERT INTO companies(name, company_type_id, contact_id, bank_detail_id, currency, 
			date_of_establishment, parent_id, is_sister_company, local_name, code, 
			is_preferred, service_type_id, created_at, updated_at)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $13)
		RETURNING id
	`

	var companyID int64

	if err := dbMod.DB.QueryRow(
		queryNewCompany,
		newCompany.Name,
		newCompany.CompanyTypeID,
		newCompany.ContactID,
		dbMod.NewNullID(newCompany.BankDetailID),
		dbMod.NewNullString(newCompany.Currency),
		dbMod.NewNullDate(newCompany.DateOfEstablishment),
		dbMod.NewNullID(newCompany.ParentID),
		newCompany.IsSisterCompany,
		dbMod.NewNullString(newCompany.LocalName),
		newCompany.Code,
		newCompany.IsPreferred,
		dbMod.NewNullID(newCompany.ServiceTypeID),
		time.Now(),
	).Scan(&companyID); err != nil {
		dbMod.Rollback()
		log.Printf("Error while inserting or scanning the Company: %s", err)
		return 0, err
	}

	return companyID, nil
}
