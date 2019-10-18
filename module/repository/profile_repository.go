package repository

import (
	"database/sql"
	"fmt"

	"github.com/febrielven/go_crud_posgres/config"
	"github.com/febrielven/go_crud_posgres/module/model"
)

type profileRepository struct {
	db *sql.DB
}

func NewProfileRepository() *profileRepository {
	db, err := config.GetPostgersDB()
	if err != nil {
		fmt.Println(err)
	}

	return &profileRepository{db}
}

func (r *profileRepository) FindAll() (model.Profiles, error) {

	query := `SELECT * FROM "profile"`

	var profiles model.Profiles

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var profile model.Profile

		err = rows.Scan(&profile.ID, &profile.FirstName, &profile.LastName, &profile.Email, &profile.Password, &profile.CreatedAt, &profile.UpdatedAt)

		if err != nil {
			return nil, err
		}

		profiles = append(profiles, profile)

	}

	return profiles, nil

}

func (r *profileRepository) FindById(id int) (*model.Profile, error) {
	query := `SELECT * FROM "profile" WHERE "id"=$1`

	var profile model.Profile

	statement, err := r.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	err = statement.QueryRow(id).Scan(&profile.ID, &profile.FirstName, &profile.LastName, &profile.Email, &profile.Password, &profile.CreatedAt, &profile.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &profile, nil
}

func (r *profileRepository) Save(profile *model.Profile) error {

	query := `INSERT INTO "profile"("id", "first_name", "last_name", "email", "password", "created_at", "updated_at")
		VALUES($1, $2, $3, $4, $5, $6, $7)`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(profile.ID,
		profile.FirstName,
		profile.LastName,
		profile.Email,
		profile.Password,
		profile.CreatedAt,
		profile.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil

}

func (r *profileRepository) Update(id int, profile *model.Profile) error {

	query := `UPDATE "profile" SET "first_name"=$1, "last_name"=$2, "email"=$3, "password"=$4, "updated_at"=$5 WHERE "id"=$6`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(profile.FirstName, profile.LastName, profile.Email, profile.Password, profile.UpdatedAt, id)

	if err != nil {
		return err
	}

	return nil
}

func (r *profileRepository) Delete(id int) error {
	query := `DELETE FROM "profile" WHERE "id"=$1`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(id)

	if err != nil {
		return err
	}

	return nil
}
