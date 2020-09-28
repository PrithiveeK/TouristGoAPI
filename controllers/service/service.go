package service

import (
	"log"
	"strconv"
	"time"

	dbMod "touristapp.com/db"
	Models "touristapp.com/models"
)

//GetAll fetches all the services available in the database
func GetAll(query map[string]string) ([]Models.Service, error) {
	var services []Models.Service

	queryServices := `
		SELECT service.id AS id, service.name AS name, service.code AS code,
		service.country_id AS country_id, service.city_id AS city_id,
		country.name AS country_name, city.name AS city_name,
		service.service_type_id AS service_type_id, st.name AS service_type_name
		FROM services AS service
		LEFT JOIN service_types AS st ON st.id = service.service_type_id
		LEFT JOIN countries AS country ON country.id = service.country_id
		LEFT JOIN cities AS city ON city.id = service.city_id
		WHERE service.status = 'ACTIVE'
	`

	if query["type"] == "placeholder" {
		queryServices += ` and service.placeholder `
	} else if query["type"] == "99A" {
		queryServices += ` and service."is_99A_service" `
	}

	rows, err := dbMod.DB.Query(queryServices)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		var service Models.Service

		if err = rows.Scan(
			&service.ID, &service.Name, &service.Code, &service.CountryID,
			&service.CityID, &service.CountryName, &service.CityName,
			&service.ServiceTypeID, &service.ServiceTypeName,
		); err != nil {
			log.Println(err)
			return nil, err
		}

		services = append(services, service)
	}

	return services, nil
}

//Get fetches a specified service from the database
func Get(id int64, query map[string]string) (Models.Service, error) {
	var service Models.Service

	queryServices := `
		SELECT service.id AS id, service.name AS name, service.code AS code,
		service.country_id AS country_id, service.city_id AS city_id,
		country.name AS country_name, city.name AS city_name,
		service.service_type_id AS service_type_id, st.name AS service_type_name
		FROM services AS service
		LEFT JOIN service_types AS st ON st.id = service.service_type_id
		LEFT JOIN countries AS country ON country.id = service.country_id
		LEFT JOIN cities AS city ON city.id = service.city_id
		WHERE service.status = 'ACTIVE'
	`

	if query["type"] == "placeholder" {
		queryServices += ` and service.placeholder `
	} else if query["type"] == "99A" {
		queryServices += ` and service."is_99A_service" `
	}

	if err := dbMod.DB.QueryRow(queryServices).Scan(
		&service.ID, &service.Name, &service.Code, &service.CountryID,
		&service.CityID, &service.CountryName, &service.CityName,
		&service.ServiceTypeID, &service.ServiceTypeName,
	); err != nil {
		log.Println(err)
		return Models.Service{}, err
	}

	return service, nil
}

//Add inserts a new row of service data into the database
func Add(newService Models.NewService, query map[string]string) (Models.Service, error) {
	queryNewService := `
		INSERT INTO services (service_type_id, name, country_id, city_id, zipcode, fax, phone_no, 
			telephone_no, email, website, street, description, chain_id, rating, location_type_id,
			location2_type_id, no_of_rooms, seats, style_id, is_unesco, has_ac, is_preferred, 
			is_licensed, code, placeholder, is_hotel, "is_99A_service", created_at, updated_at)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18,
			$19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $28)
		RETURNING id
	`

	newService.IsHotel = newService.ServiceTypeID == 1
	newService.Placeholder = query["type"] == "placeholder"
	newService.Code = getCode()

	var serviceID int64

	if err := dbMod.DB.QueryRow(
		queryNewService,
		newService.ServiceTypeID,
		newService.Name,
		newService.CountryID,
		newService.CityID,
		dbMod.NewNullString(newService.Zipcode),
		dbMod.NewNullString(newService.Fax),
		dbMod.NewNullString(newService.PhoneNO),
		dbMod.NewNullString(newService.TelephoneNo),
		dbMod.NewNullString(newService.Email),
		dbMod.NewNullString(newService.Website),
		dbMod.NewNullString(newService.Street),
		dbMod.NewNullString(newService.Description),
		dbMod.NewNullID(newService.ChainID),
		dbMod.NewNullString(newService.StarRating),
		dbMod.NewNullID(newService.LocationTypeID),
		dbMod.NewNullID(newService.Location2TypeID),
		dbMod.NewNullString(newService.NoOfRooms),
		dbMod.NewNullString(newService.Seats),
		dbMod.NewNullID(newService.StyleID),
		newService.IsUNESCO,
		newService.HasAC,
		newService.IsPreferred,
		newService.IsLicensed,
		newService.Code,
		newService.Placeholder,
		newService.IsHotel,
		newService.Is99AService,
		time.Now(),
	).Scan(&serviceID); err != nil {
		dbMod.Rollback()
		log.Printf("Error inserting service data: %s", err)
		return Models.Service{}, nil
	}

	return Get(serviceID, query)
}

//getCode fetches the next value to be inserted
func getCode() string {
	var id int
	if err := dbMod.DB.QueryRow(`SELECT nextval('services_id_seq')`).Scan(&id); err != nil {
		log.Printf("Error finging services nextvalue")
		return ""
	}
	return "P" + strconv.Itoa(id)
}

//MapCategories maps the categories to the services
func MapCategories(id int64, newCategories []int64) error {
	if newCategories == nil {
		return nil
	}
	queryNewMC := `
		INSERT INTO service_category_mappings(service_id, category_type_id, created_at, updated_at)
		VALUES($1, $2, $3, $3)
	`
	for _, cID := range newCategories {
		if _, err := dbMod.DB.Query(queryNewMC, id, cID, time.Now()); err != nil {
			dbMod.Rollback()
			log.Printf("Error inserting category: %d in db: %s", cID, err)
		}
	}
	return nil
}

//MapAmenities maps the amenities to the service
func MapAmenities(id int64, newAmenities []int64) error {
	if newAmenities == nil {
		return nil
	}
	queryNewMA := `
		INSERT INTO service_amenity_mappings(service_id, amenity_id, created_at, updated_at)
		VALUES($1, $2, $3, $3)
	`
	for _, aID := range newAmenities {
		if _, err := dbMod.DB.Query(queryNewMA, id, aID, time.Now()); err != nil {
			dbMod.Rollback()
			log.Printf("Error inserting amenity: %d in db: %s", aID, err)
		}
	}
	return nil
}

//MapLinkedServices maps the services linked to the inserted service
func MapLinkedServices(id int64, newLinkedServices []int64) error {
	if newLinkedServices == nil {
		return nil
	}
	queryNewMLS := `
		INSERT INTO linked_services_mappings(service_id, linked_service_id, created_at, updated_at)
		VALUES($1, $2, $3, $3)
	`
	for _, lsID := range newLinkedServices {
		if _, err := dbMod.DB.Query(queryNewMLS, id, lsID, time.Now()); err != nil {
			dbMod.Rollback()
			log.Printf("Error inserting linked service: %d in db: %s", lsID, err)
		}
	}
	return nil
}

//MapSupplier maps the supplier to the service
func MapSupplier(id, sID int64) error {
	queryNewS := `
		INSERT INTO services_supplier_mappings(service_id, supplier_id, created_at, updated_at)
		VALUES($1,$2, $3, $3)
	`

	if _, err := dbMod.DB.Query(queryNewS, id, sID, time.Now()); err != nil {
		dbMod.Rollback()
		log.Printf("Error mapping supplier with services: %s", err)
		return err
	}
	return nil
}

//MapTC maps the terms and conditions to the service
func MapTC(id, tcID int64) error {
	queryNewTC := `
		INSERT INTO services_tc_mappings(service_id, tc_id, created_at, updated_at)
		VALUES($1,$2, $3, $3)
	`

	if _, err := dbMod.DB.Query(queryNewTC, id, tcID, time.Now()); err != nil {
		dbMod.Rollback()
		log.Printf("Error mapping TC with services: %s", err)
		return err
	}
	return nil
}

//AddPricing adds the price details to the service
func AddPricing(id int64, newPriceDetails []Models.NewPriceDetails) error {
	if newPriceDetails == nil {
		return nil
	}
	queryNewPriceDetails := `
		INSERT INTO service_price_details(service_id, currency, room_type, person_type,
			pricing_type, charge_type_id, menu_type_id, breakfast_type_id, description,
			type, category_type_id, price, created_at, updated_at)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $13)
	`

	for _, price := range newPriceDetails {
		if _, err := dbMod.DB.Query(
			queryNewPriceDetails,
			id,
			dbMod.NewNullString(price.Currency),
			dbMod.NewNullID(price.RoomType),
			dbMod.NewNullID(price.PersonType),
			dbMod.NewNullString(price.Pricing),
			dbMod.NewNullID(price.Charge),
			dbMod.NewNullID(price.MealType),
			dbMod.NewNullID(price.BreakfastType),
			dbMod.NewNullString(price.Description),
			dbMod.NewNullString(price.Type),
			dbMod.NewNullID(price.Category),
			price.Price,
			time.Now(),
		); err != nil {
			dbMod.Rollback()
			log.Printf("Error mapping tc with services: %s", err)
			return err
		}
	}
	return nil
}

//Overall is the model for overall data
type Overall struct {
	Services  int64 `json:"services"`
	Countries int64 `json:"countries"`
	Cities    int64 `json:"cities"`
}

//OverallData fetches the count of services, countries, cities
func OverallData() (Overall, error) {
	var overall Overall

	if err := dbMod.DB.QueryRow(`SELECT COUNT(*) FROM services`).Scan(&overall.Services); err != nil {
		log.Printf("Error counting services: %s", err)
		return Overall{}, err
	}

	if err := dbMod.DB.QueryRow(`SELECT COUNT(*) FROM countries`).Scan(&overall.Countries); err != nil {
		log.Printf("Error counting countries: %s", err)
		return Overall{}, err
	}

	if err := dbMod.DB.QueryRow(`SELECT COUNT(*) FROM cities`).Scan(&overall.Cities); err != nil {
		log.Printf("Error counting cities: %s", err)
		return Overall{}, err
	}
	return overall, nil
}

//Random fetches all the services for a random country
func Random() ([]Models.Service, error) {
	var randomCountry int64

	if err := dbMod.DB.QueryRow(`SELECT country_id from services order by random() limit 1`).Scan(&randomCountry); err != nil {
		log.Printf("Error finding random record countries: %s", err)
		return nil, err
	}

	queryServices := `
		SELECT service.id AS id, service.name AS name, service.code AS code,
		service.country_id AS country_id, service.city_id AS city_id,
		country.name AS country_name, city.name AS city_name,
		service.service_type_id AS service_type_id, st.name AS service_type_name
		FROM services AS service
		LEFT JOIN service_types AS st ON st.id = service.service_type_id
		LEFT JOIN countries AS country ON country.id = service.country_id
		LEFT JOIN cities AS city ON city.id = service.city_id
		WHERE service.status = 'ACTIVE' and service.country_id = $1
	`
	var services []Models.Service
	rows, err := dbMod.DB.Query(queryServices, randomCountry)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		var service Models.Service

		if err = rows.Scan(
			&service.ID, &service.Name, &service.Code, &service.CountryID,
			&service.CityID, &service.CountryName, &service.CityName,
			&service.ServiceTypeID, &service.ServiceTypeName,
		); err != nil {
			log.Println(err)
			return nil, err
		}

		services = append(services, service)
	}

	return services, nil

}
