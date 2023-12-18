package content

import (
	"context"
	pb "editory_submission/genproto/content_service"
	"editory_submission/pkg/helper"
	"editory_submission/storage"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type SubjectRepo struct {
	db *pgxpool.Pool
}

func NewSubjectRepo(db *pgxpool.Pool) storage.SubjectRepoI {
	return &SubjectRepo{
		db: db,
	}
}

func (s *SubjectRepo) Create(ctx context.Context, req *pb.CreateSubjectReq) (res *pb.CreateSubjectRes, err error) {

	res = &pb.CreateSubjectRes{}

	query := `INSERT INTO "subject" (
		id,                            
    	title     
	) VALUES (
		$1,
		$2
	) RETURNING 
	    id, 
	    title`

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

func (s *SubjectRepo) Get(ctx context.Context, req *pb.GetSubjectReq) (res *pb.GetSubjectRes, err error) {
	res = &pb.GetSubjectRes{}

	query := `SELECT
		id,                            
    	title
	FROM
		"subject"
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

func (s *SubjectRepo) GetList(ctx context.Context, req *pb.GetSubjectListReq) (res *pb.GetSubjectListRes, err error) {
	res = &pb.GetSubjectListRes{}
	params := make(map[string]interface{})
	var arr []interface{}

	query := `SELECT
		id,           
    	title
	FROM
		"subject"`
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

	cQ := `SELECT count(1) FROM "subject"` + filter

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
		obj := &pb.Subject{}

		err = rows.Scan(
			&obj.Id,
			&obj.Title,
		)
		if err != nil {
			return res, err
		}

		res.Subjects = append(res.Subjects, obj)
	}

	return res, nil
}

func (s *SubjectRepo) Update(ctx context.Context, req *pb.UpdateSubjectReq) (res *pb.UpdateSubjectRes, err error) {
	res = &pb.UpdateSubjectRes{}

	query := `UPDATE "subject" SET                
    	title = :title
	WHERE
		id = :id
	RETURNING 
		id,
		title`

	params := map[string]interface{}{
		"id":    req.GetId(),
		"title": req.GetTitle(),
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
func (s *SubjectRepo) Delete(ctx context.Context, req *pb.DeleteSubjectReq) (rowsAffected int64, err error) {
	query := `DELETE FROM "subject" WHERE id = $1`

	result, err := s.db.Exec(ctx, query, req.GetId())
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}
