package access

import (
	"log"

	db "touristapp.com/db"
	Models "touristapp.com/models"
)

//GetAll ftches all the access available int the database
func GetAll() ([]Models.Access, error) {
	var accesses []Models.Access

	rows, err := db.DB.Query(`SELECT id, role_id, access_id from roles_accesses_mappings`)
	if err != nil {
		log.Printf("Access query error: %s", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var access Models.Access
		if err := rows.Scan(
			&access.ID,
			&access.RoleID,
			&access.AccessID,
		); err != nil {
			log.Printf("Access row scan error: %s", err)
			return nil, err
		}

		accesses = append(accesses, access)
	}

	return accesses, nil
}
