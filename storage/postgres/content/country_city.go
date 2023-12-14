package content

import (
	"context"
	"editory_submission/genproto/content_service"
	"editory_submission/pkg/helper"
	"editory_submission/pkg/util"
	"editory_submission/storage"
	"github.com/jackc/pgx/v4/pgxpool"
)

type CountryAndCityRepo struct {
	db *pgxpool.Pool
}

func NewCountryAndCityRepo(db *pgxpool.Pool) storage.CountryAndCityRepoI {
	return &CountryAndCityRepo{
		db: db,
	}
}

func (c *CountryAndCityRepo) GetCountyList(ctx context.Context, in *content_service.GetCountryListReq) (res *content_service.GetCountryListRes, err error) {
	res = &content_service.GetCountryListRes{}
	params := make(map[string]interface{})
	var arr []interface{}

	query := `SELECT
		id,                 
    	title,
    	title_uz,
    	title_ru
	FROM
		"country"`
	filter := " WHERE 1=1"

	offset := " OFFSET 0"

	limit := " LIMIT 10"

	if len(in.Search) > 0 {
		params["search"] = in.Search
		filter += ` AND ((title ILIKE '%' || :search || '%')
					OR (title_uz ILIKE '%' || :search || '%')
					OR (title_ru ILIKE '%' || :search || '%'))`
	}

	if in.Offset > 0 {
		params["offset"] = in.Offset
		offset = " OFFSET :offset"
	}

	if in.Limit > 0 {
		params["limit"] = in.Limit
		limit = " LIMIT :limit"
	}

	cQ := `SELECT count(1) FROM "country"` + filter

	cQ, arr = helper.ReplaceQueryParams(cQ, params)

	err = c.db.QueryRow(ctx, cQ, arr...).Scan(
		&res.Count,
	)
	if err != nil {
		return res, err
	}

	q := query + filter + offset + limit

	q, arr = helper.ReplaceQueryParams(q, params)
	rows, err := c.db.Query(ctx, q, arr...)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		obj := &content_service.Country{}

		err = rows.Scan(
			&obj.Id,
			&obj.Title,
			&obj.TitleUz,
			&obj.TitleRu,
		)
		if err != nil {
			return res, err
		}

		res.Countries = append(res.Countries, obj)
	}

	return res, nil
}

func (c *CountryAndCityRepo) GetCityList(ctx context.Context, in *content_service.GetCityListReq) (res *content_service.GetCityListRes, err error) {
	res = &content_service.GetCityListRes{}
	params := make(map[string]interface{})
	var arr []interface{}

	query := `SELECT
		id,                 
    	title,
    	title_uz,
    	title_ru,
    	country_id
	FROM
		"city"`
	filter := " WHERE 1=1"

	offset := " OFFSET 0"

	limit := " LIMIT 10"

	if len(in.Search) > 0 {
		params["search"] = in.Search
		filter += ` AND ((title ILIKE '%' || :search || '%')
					OR (title_uz ILIKE '%' || :search || '%')
					OR (title_ru ILIKE '%' || :search || '%'))`
	}

	if in.Offset > 0 {
		params["offset"] = in.Offset
		offset = " OFFSET :offset"
	}

	if in.Limit > 0 {
		params["limit"] = in.Limit
		limit = " LIMIT :limit"
	}

	if util.IsValidUUID(in.GetCountryId()) {
		params["country_id"] = in.CountryId
		filter += ` AND country_id = :country_id`
	}

	cQ := `SELECT count(1) FROM "city"` + filter

	cQ, arr = helper.ReplaceQueryParams(cQ, params)

	err = c.db.QueryRow(ctx, cQ, arr...).Scan(
		&res.Count,
	)
	if err != nil {
		return res, err
	}

	q := query + filter + offset + limit

	q, arr = helper.ReplaceQueryParams(q, params)
	rows, err := c.db.Query(ctx, q, arr...)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		obj := &content_service.City{}

		err = rows.Scan(
			&obj.Id,
			&obj.Title,
			&obj.TitleUz,
			&obj.TitleRu,
			&obj.CountryId,
		)
		if err != nil {
			return res, err
		}

		res.Cities = append(res.Cities, obj)
	}

	return res, nil
}
