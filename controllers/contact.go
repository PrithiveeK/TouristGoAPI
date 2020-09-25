package controllers

import (
	"log"
	"time"

	dbMod "touristapp.com/db"
	Models "touristapp.com/models"
)

//AddContact inserts new contact information to the database
//for a company or user
func AddContact(newContact Models.NewContact) (int64, error) {
	queryNewContact := `
		INSERT INTO contacts(country_id, city_id, zipcode, fax, phone_no,
			telephone_no, email, street, website, skype_id, created_at, updated_at)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $11)
		RETURNING id
	`

	var contactID int64

	if err := dbMod.DB.QueryRow(
		queryNewContact,
		newContact.CountryID,
		newContact.CityID,
		dbMod.NewNullString(newContact.Zipcode),
		dbMod.NewNullString(newContact.Fax),
		dbMod.NewNullString(newContact.PhoneNo),
		dbMod.NewNullString(newContact.TelephoneNo),
		newContact.Email,
		dbMod.NewNullString(newContact.Street),
		dbMod.NewNullString(newContact.Website),
		dbMod.NewNullString(newContact.SkypeID),
		time.Now(),
	).Scan(&contactID); err != nil {
		dbMod.Rollback()
		log.Printf("Error while inserting or scanning the contact: %s", err)
		return 0, err
	}

	return contactID, nil
}
