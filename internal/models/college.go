package models

import "database/sql"

var DB *sql.DB

type College struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	//Code       string `json:"code"`
	LocationId int `json:"location_id"`
}

type Location struct {
	ID           int    `json:"id"`
	Street       string `json:"street"`
	Number       int    `json:"number"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
	Cep          int    `json:"cep"`
}

//	type CollegeModel struct {
//		DB *sql.DB
//	}
func (c *College) AllUniversity() ([]College, error) {
	rows, err := DB.Query("SELECT * FROM college")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var colleges []College

	for rows.Next() {
		var college College
		err := rows.Scan(&college.ID, &college.Name, &college.LocationId)
		if err != nil {
			return nil, err
		}
		colleges = append(colleges, college)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return colleges, nil
}
