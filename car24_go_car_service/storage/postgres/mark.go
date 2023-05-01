package postgres

import (
	"database/sql"
	"fmt"

	cs "gitlab.udevs.io/car24/car24_go_car_service/modules/car24/car24_car_service"
	"gitlab.udevs.io/car24/car24_go_car_service/pkg/helper"
	"gitlab.udevs.io/car24/car24_go_car_service/storage/repo"

	"github.com/jmoiron/sqlx"
)

type markRepo struct {
	db *sqlx.DB
}

func NewMark(db *sqlx.DB) repo.MarkI {
	return &markRepo{
		db: db,
	}
}

func (br *markRepo) Create(mark *cs.CreateMarkModel) (string, error) {
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
		mark(
			id,
			brand_id,
			name,
			updated_at
		)
	VALUES ($1, $2, $3, now())`

	_, err = tsx.Exec(
		query,
		mark.ID,
		mark.BrandID,
		mark.Name,
	)

	if err != nil {
		return "", err
	}

	return mark.ID, nil
}

func (br *markRepo) Get(id string) (*cs.MarkModel, error) {
	var (
		name      sql.NullString
		brandID   sql.NullString
		brandName sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	query := `
		SELECT
			m.name,
			b.id,
			b.name,
			m.created_at,
			m.updated_at
		FROM mark AS m
		JOIN brand AS b ON b.id = m.brand_id
		WHERE m.id = $1
	`

	row := br.db.QueryRow(query, id)
	err := row.Scan(
		&name,
		&brandID,
		&brandName,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &cs.MarkModel{
		ID:        id,
		Name:      name.String,
		BrandID:   brandID.String,
		BrandName: brandName.String,
		UpdatedAt: updatedAt.String,
		CreatedAt: createdAt.String,
	}, nil
}

func (br *markRepo) GetAll(queryParam cs.MarkQueryParamModel) (res cs.MarkListModel, err error) {

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
			m.id,
			m.name,
			b.id,
			b.name,
			m.updated_at,
			m.created_at
		FROM
			mark AS m
		JOIN brand AS b ON b.id = m.brand_id
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
			brandId   sql.NullString
			brandName sql.NullString
			updatedAt sql.NullString
			createdAt sql.NullString
		)

		err := rows.Scan(
			&res.Count,
			&id,
			&name,
			&brandId,
			&brandName,
			&updatedAt,
			&createdAt,
		)
		if err != nil {
			return res, err
		}

		res.Marks = append(res.Marks, cs.MarkModel{
			ID:        id.String,
			Name:      name.String,
			BrandID:   brandId.String,
			BrandName: brandName.String,
			UpdatedAt: updatedAt.String,
			CreatedAt: createdAt.String,
		})
	}

	return res, nil
}

func (br *markRepo) Update(mark *cs.UpdateMarkModel) (err error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
			UPDATE 
			    mark
			SET
				brand_id = :brand_id,
				name = :name,
				updated_at = now()
			WHERE
				id = :id`
	params = map[string]interface{}{
		"id":       mark.ID,
		"brand_id": mark.BrandID,
		"name":     mark.Name,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	_, err = br.db.Exec(query, args...)
	if err != nil {
		return
	}

	return err
}

func (br *markRepo) Delete(id string) error {
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

	query := `DELETE FROM mark WHERE id=$1`
	_, err = tsx.Exec(query, id)

	if err != nil {
		return err
	}

	return err
}
