package agent

import (
	"log"

	Controller "touristapp.com/controllers"
	dbMod "touristapp.com/db"
	Models "touristapp.com/models"
)

//GetAll fetches all the agents or sub agents in the database
func GetAll(query map[string]int64) ([]Models.Agent, error) {
	var agents []Models.Agent

	queryAgent := `
		SELECT agent.id, agent.code AS code, agent.name AS name, 
		contact.city_id AS city_id, contact.country_id AS country_id, 
		city.name AS city_name, country.name AS country_name
		FROM companies AS agent
		LEFT JOIN contacts AS contact ON contact.id = agent.contact_id
		LEFT JOIN cities AS city ON city.id = contact.city_id
		LEFT JOIN countries AS country ON country.id = contact.country_id
		WHERE agent.status = 'ACTIVE' and agent.company_type_id = $1
	`
	rows, err := dbMod.DB.Query(queryAgent, query["type"])
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for rows.Next() {
		var agent Models.Agent

		if err = rows.Scan(
			&agent.ID, &agent.Code, &agent.Name, &agent.CityID,
			&agent.CountryID, &agent.CityName, &agent.CountryName,
		); err != nil {
			log.Println(err)
			return nil, err
		}

		agents = append(agents, agent)
	}

	return agents, nil
}

//Get fetches one row of the specified agent from the database
func Get(id int64, query map[string]int64) (Models.Agent, error) {
	var agent Models.Agent

	queryAgent := `
		SELECT agent.id, agent.code AS code, agent.name AS name, 
		contact.city_id AS city_id, contact.country_id AS country_id, 
		city.name AS city_name, country.name AS country_name
		FROM companies AS agent
		LEFT JOIN contacts AS contact ON contact.id = agent.contact_id
		LEFT JOIN cities AS city ON city.id = contact.city_id
		LEFT JOIN countries AS country ON country.id = contact.country_id
		WHERE agent.status = 'ACTIVE' and agent.id = $1 and agent.company_type_id = $2
	`
	if err := dbMod.DB.QueryRow(queryAgent, id, query["type"]).Scan(
		&agent.ID, &agent.Code, &agent.Name, &agent.CityID,
		&agent.CountryID, &agent.CityName, &agent.CountryName,
	); err != nil {
		log.Println(err)
		return Models.Agent{}, err
	}

	return agent, nil
}

//Add inserts a new agent to the database
func Add(newAgent *Models.NewAgent, companyType int64) (Models.Agent, error) {
	contactID, err := Controller.AddContact(newAgent.Contact)
	if err != nil {
		log.Printf("Cannot create contact: %s", err)
		return Models.Agent{}, err
	}

	newAgent.Agent.ContactID = contactID
	newAgent.Agent.CompanyTypeID = companyType

	agentID, err := Controller.AddCompany(newAgent.Agent)
	if err != nil {
		log.Printf("Cannot create agent: %s", err)
		return Models.Agent{}, nil
	}

	return Get(agentID, map[string]int64{
		"type": companyType,
	})
}
