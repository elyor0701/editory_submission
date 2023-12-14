package content

import (
	"context"
	pb "editory_submission/genproto/content_service"
	"editory_submission/pkg/helper"
	"editory_submission/storage"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UniversityRepo struct {
	db *pgxpool.Pool
}

func NewUniversityRepo(db *pgxpool.Pool) storage.UniversityRepoI {
	return &UniversityRepo{
		db: db,
	}
}

func (s *UniversityRepo) Create(ctx context.Context, req *pb.CreateUniversityReq) (res *pb.CreateUniversityRes, err error) {

	res = &pb.CreateUniversityRes{}

	query := `INSERT INTO "university" (
		id,                            
    	title,         
		logo
	) VALUES (
		$1,
		$2,
		$3
	) RETURNING 
	    id, 
	    title,
	    logo`

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	err = s.db.QueryRow(ctx, query,
		id.String(),
		req.GetTitle(),
		req.GetLogo(),
	).Scan(
		&res.Id,
		&res.Title,
		&res.Logo,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *UniversityRepo) Get(ctx context.Context, req *pb.GetUniversityReq) (res *pb.GetUniversityRes, err error) {
	res = &pb.GetUniversityRes{}

	query := `SELECT
		id,                            
    	title,
    	logo
	FROM
		"university"
	WHERE
		id = $1`

	err = s.db.QueryRow(
		ctx,
		query,
		req.GetId()).Scan(
		&res.Id,
		&res.Title,
		&res.Logo,
	)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *UniversityRepo) GetList(ctx context.Context, req *pb.GetUniversityListReq) (res *pb.GetUniversityListRes, err error) {
	res = &pb.GetUniversityListRes{}
	params := make(map[string]interface{})
	var arr []interface{}

	query := `SELECT
		id,           
    	title,         
    	logo
	FROM
		"university"`
	filter := " WHERE 1=1"

	offset := " OFFSET 0"

	limit := " LIMIT 10"

	if len(req.Search) > 0 {
		params["search"] = req.Search
		filter += ` AND (title ILIKE '%' || :search || '%')`
	}

	if req.Offset > 0 {
		params["offset"] = req.Offset
		offset = " OFFSET :offset"
	}

	if req.Limit > 0 {
		params["limit"] = req.Limit
		limit = " LIMIT :limit"
	}

	cQ := `SELECT count(1) FROM "university"` + filter

	cQ, arr = helper.ReplaceQueryParams(cQ, params)

	err = s.db.QueryRow(ctx, cQ, arr...).Scan(
		&res.Count,
	)
	if err != nil {
		return res, err
	}

	q := query + filter + offset + limit

	q, arr = helper.ReplaceQueryParams(q, params)
	rows, err := s.db.Query(ctx, q, arr...)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		obj := &pb.University{}

		err = rows.Scan(
			&obj.Id,
			&obj.Title,
			&obj.Logo,
		)
		if err != nil {
			return res, err
		}

		res.Universities = append(res.Universities, obj)
	}

	return res, nil
}

func (s *UniversityRepo) Update(ctx context.Context, req *pb.UpdateUniversityReq) (res *pb.UpdateUniversityRes, err error) {
	res = &pb.UpdateUniversityRes{}

	query := `UPDATE "university" SET                
    	title = :title,           
    	logo = :logo
	WHERE
		id = :id
	RETURNING 
		id,
		title,
		logo`

	params := map[string]interface{}{
		"id":    req.GetId(),
		"title": req.GetTitle(),
		"logo":  req.GetLogo(),
	}

	q, arr := helper.ReplaceQueryParams(query, params)
	err = s.db.QueryRow(ctx, q, arr...).Scan(
		&res.Id,
		&res.Title,
		&res.Logo,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func (s *UniversityRepo) Delete(ctx context.Context, req *pb.DeleteUniversityReq) (rowsAffected int64, err error) {
	query := `DELETE FROM "university" WHERE id = $1`

	result, err := s.db.Exec(ctx, query, req.GetId())
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}
