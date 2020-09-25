package supplier

import (
	"log"

	Controller "touristapp.com/controllers"
	db "touristapp.com/db"
	Models "touristapp.com/models"
)

//AllSuppliers is the structure for dividing the
//preferred and not preferred suppliers
type AllSuppliers struct {
	Preferred []Models.Supplier `json:"preferred"`
	Other     []Models.Supplier `json:"other"`
}

//GetAll fetches all the suppliers available in the datbase
func GetAll() (AllSuppliers, error) {
	var suppliers AllSuppliers

	querySuppliers := `
		SELECT supplier.id AS id, supplier.name AS name, supplier.code AS code,
		supplier.is_preferred AS is_preferred, 
		contact.city_id AS city_id, contact.country_id AS country_id, 
		city.name AS city_name, country.name AS country_name,
		contact.email AS email, contact.phone_no AS phone_no,
		supplier.service_type_id AS service_type_id, st.name AS service_type_name
		FROM companies AS supplier
		LEFT JOIN contacts AS contact ON contact.id = supplier.contact_id
		LEFT JOIN cities AS city ON city.id = contact.city_id
		LEFT JOIN countries AS country ON country.id = contact.country_id
		LEFT JOIN service_types AS st ON st.id = supplier.service_type_id
		WHERE supplier.status = 'ACTIVE' and supplier.company_type_id = 4
	`

	rows, err := db.DB.Query(querySuppliers)
	if err != nil {
		log.Printf("Supplier query error: %s\n", err)
		return AllSuppliers{}, err
	}

	for rows.Next() {
		var supplier Models.Supplier

		if err = rows.Scan(
			&supplier.ID, &supplier.Name, &supplier.Code,
			&supplier.IsPreferred, &supplier.CityID, &supplier.CountryID,
			&supplier.CityName, &supplier.CountryName, &supplier.Email,
			&supplier.PhoneNo, &supplier.ServiceTypeID, &supplier.ServiceTypeName,
		); err != nil {
			log.Printf("Supplier rows scan error: %s\n", err)
			return AllSuppliers{}, err
		}
		if supplier.IsPreferred {
			suppliers.Preferred = append(suppliers.Preferred, supplier)
		} else {
			suppliers.Other = append(suppliers.Other, supplier)
		}
	}

	return suppliers, nil
}

//Get fetches the specified supplier from the database
func Get(id int64) (Models.Supplier, error) {
	var supplier Models.Supplier

	querySuppliers := `
		SELECT supplier.id AS id, supplier.name AS name, supplier.code AS code,
		supplier.is_preferred AS is_preferred, 
		contact.city_id AS city_id, contact.country_id AS country_id, 
		city.name AS city_name, country.name AS country_name,
		contact.email AS email, contact.phone_no AS phone_no,
		supplier.service_type_id AS service_type_id, st.name AS service_type_name
		FROM companies AS supplier
		LEFT JOIN contacts AS contact ON contact.id = supplier.contact_id
		LEFT JOIN cities AS city ON city.id = contact.city_id
		LEFT JOIN countries AS country ON country.id = contact.country_id
		LEFT JOIN service_types AS st ON st.id = supplier.service_type_id
		WHERE supplier.status = 'ACTIVE' and supplier.id = $1 and supplier.company_type_id = 4
	`

	if err := db.DB.QueryRow(querySuppliers, id).Scan(
		&supplier.ID, &supplier.Name, &supplier.Code,
		&supplier.IsPreferred, &supplier.CityID, &supplier.CountryID,
		&supplier.CityName, &supplier.CountryName, &supplier.Email,
		&supplier.PhoneNo, &supplier.ServiceTypeID, &supplier.ServiceTypeName,
	); err != nil {
		log.Printf("Supplier query or scan error: %s\n", err)
		return Models.Supplier{}, err
	}
	return supplier, nil
}

//Add inserts a new supplier into the database
func Add(newSupplier Models.NewSupplier) (Models.Supplier, error) {
	contactID, err := Controller.AddContact(newSupplier.Contact)
	if err != nil {
		log.Printf("Cannot create contact: %s", err)
		return Models.Supplier{}, err
	}

	bankDetailID, err := Controller.AddBankDetail(newSupplier.Bank)
	if err != nil {
		log.Printf("Cannot create bank detail: %s", err)
		return Models.Supplier{}, err
	}

	newSupplier.Supplier.ContactID = contactID
	newSupplier.Supplier.BankDetailID = bankDetailID
	newSupplier.Supplier.CompanyTypeID = 4

	supplierID, err := Controller.AddCompany(newSupplier.Supplier)
	if err != nil {
		log.Printf("Cannot create Supplier: %s", err)
		return Models.Supplier{}, err
	}

	return Get(supplierID)
}
