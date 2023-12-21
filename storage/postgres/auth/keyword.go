package auth

import (
	"context"
	pb "editory_submission/genproto/auth_service"
	"editory_submission/pkg/helper"
	"editory_submission/storage"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type KeywordRepo struct {
	db *pgxpool.Pool
}

func NewKeywordRepo(db *pgxpool.Pool) storage.KeywordRepoI {
	return &KeywordRepo{
		db: db,
	}
}

func (s *KeywordRepo) Create(ctx context.Context, req *pb.CreateKeywordReq) (res *pb.CreateKeywordRes, err error) {

	res = &pb.CreateKeywordRes{}

	query := `INSERT INTO "keyword" (
		id,                            
    	word     
	) VALUES (
		$1,
		$2
	) RETURNING 
	    id, 
	    word`

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	err = s.db.QueryRow(ctx, query,
		id.String(),
		req.GetTitle(),
	).Scan(
		&res.Id,
		&res.Title,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *KeywordRepo) Get(ctx context.Context, req *pb.GetKeywordReq) (res *pb.GetKeywordRes, err error) {
	res = &pb.GetKeywordRes{}

	query := `SELECT
		id,                            
    	word
	FROM
		"keyword"
	WHERE
		id = $1`

	err = s.db.QueryRow(
		ctx,
		query,
		req.GetId()).Scan(
		&res.Id,
		&res.Title,
	)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *KeywordRepo) GetList(ctx context.Context, req *pb.GetKeywordListReq) (res *pb.GetKeywordListRes, err error) {
	res = &pb.GetKeywordListRes{}
	params := make(map[string]interface{})
	var arr []interface{}

	query := `SELECT
		id,           
    	word
	FROM
		"keyword"`
	filter := " WHERE 1=1"

	offset := " OFFSET 0"

	limit := " LIMIT 10"

	if len(req.Search) > 0 {
		params["search"] = req.Search
		filter += ` AND (word ILIKE '%' || :search || '%')`
	}

	if req.Offset > 0 {
		params["offset"] = req.Offset
		offset = " OFFSET :offset"
	}

	if req.Limit > 0 {
		params["limit"] = req.Limit
		limit = " LIMIT :limit"
	}

	cQ := `SELECT count(1) FROM "keyword"` + filter

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
		obj := &pb.Keyword{}

		err = rows.Scan(
			&obj.Id,
			&obj.Title,
		)
		if err != nil {
			return res, err
		}

		res.Keywords = append(res.Keywords, obj)
	}

	return res, nil
}

func (s *KeywordRepo) Update(ctx context.Context, req *pb.UpdateKeywordReq) (res *pb.UpdateKeywordRes, err error) {
	res = &pb.UpdateKeywordRes{}

	query := `UPDATE "keyword" SET                
    	word = :word
	WHERE
		id = :id
	RETURNING 
		id,
		word`

	params := map[string]interface{}{
		"id":   req.GetId(),
		"word": req.GetTitle(),
	}

	q, arr := helper.ReplaceQueryParams(query, params)
	err = s.db.QueryRow(ctx, q, arr...).Scan(
		&res.Id,
		&res.Title,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func (s *KeywordRepo) Delete(ctx context.Context, req *pb.DeleteKeywordReq) (rowsAffected int64, err error) {
	query := `DELETE FROM "keyword" WHERE id = $1`

	result, err := s.db.Exec(ctx, query, req.GetId())
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}
