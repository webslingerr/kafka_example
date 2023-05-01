package postgres

import (
	"database/sql"
	"fmt"

	cs "gitlab.udevs.io/car24/car24_go_car_service/modules/car24/car24_car_service"
	"gitlab.udevs.io/car24/car24_go_car_service/pkg/helper"
	"gitlab.udevs.io/car24/car24_go_car_service/storage/repo"

	"github.com/jmoiron/sqlx"
)

type brandRepo struct {
	db *sqlx.DB
}

func NewBrand(db *sqlx.DB) repo.BrandI {
	return &brandRepo{
		db: db,
	}
}

func (br *brandRepo) Create(brand *cs.CreateBrandModel) (string, error) {
	tsx, err := br.db.Begin()
	if err != nil {
		return "", err
	}
	defer func() {
		if err != nil {
			_ = tsx.Rollback()
		} else {
			err := tsx.Commit()
			if err != nil {
				fmt.Println("While committing transaction ", err)
			}
		}
	}()

	query := `
	INSERT INTO 
		brand(
			id,
			name,
			updated_at
		)
	VALUES ($1, $2, now())`

	_, err = tsx.Exec(
		query,
		brand.ID,
		brand.Name,
	)

	if err != nil {
		return "", err
	}

	return brand.ID, nil
}

func (br *brandRepo) Get(id string) (*cs.BrandModel, error) {
	var (
		name      sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	query := `
		SELECT
			name,
			created_at,
			updated_at
		FROM brand
		WHERE id = $1
	`

	row := br.db.QueryRow(query, id)
	err := row.Scan(
		&name,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &cs.BrandModel{
		ID:        id,
		Name:      name.String,
		UpdatedAt: updatedAt.String,
		CreatedAt: createdAt.String,
	}, nil
}

func (br *brandRepo) GetAll(queryParam cs.BrandQueryParamModel) (res cs.BrandListModel, err error) {

	var (
		filter = "WHERE 1=1"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
		params = map[string]interface{}{}
		query  string
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			name,
			updated_at,
			created_at
		FROM
			brand
	`

	if queryParam.Offset > 0 {
		offset = " OFFSET :offset"
		params["offset"] = queryParam.Offset
	}

	if queryParam.Limit > 0 {
		limit = " LIMIT :limit"
		params["limit"] = queryParam.Limit
	}

	query += filter + offset + limit
	query, args := helper.ReplaceQueryParams(query, params)
	rows, err := br.db.Query(query, args...)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var (
			id        sql.NullString
			name      sql.NullString
			updatedAt sql.NullString
			createdAt sql.NullString
		)

		err := rows.Scan(
			&res.Count,
			&id,
			&name,
			&updatedAt,
			&createdAt,
		)
		if err != nil {
			return res, err
		}

		res.Brands = append(res.Brands, cs.BrandModel{
			ID:        id.String,
			Name:      name.String,
			UpdatedAt: updatedAt.String,
			CreatedAt: createdAt.String,
		})
	}

	return res, nil
}

func (br *brandRepo) Update(brand *cs.UpdateBrandModel) (err error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
			UPDATE 
			    brand
			SET
				name = :name,
				updated_at = now()
			WHERE
				id = :id`
	params = map[string]interface{}{
		"id":   brand.ID,
		"name": brand.Name,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	_, err = br.db.Exec(query, args...)
	if err != nil {
		return
	}

	return err
}

func (br *brandRepo) Delete(id string) error {
	tsx, err := br.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tsx.Rollback()
		} else {
			err := tsx.Commit()
			if err != nil {
				fmt.Println("While committing transaction ", err)
			}
		}
	}()

	query := `DELETE FROM brand WHERE id=$1`
	_, err = tsx.Exec(query, id)

	if err != nil {
		return err
	}

	return err
}
