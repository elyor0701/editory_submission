package submission

import (
	"context"
	pb "editory_submission/genproto/submission_service"
	"editory_submission/pkg/helper"
	"editory_submission/pkg/util"
	"editory_submission/storage"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"strings"
)

type FileRepo struct {
	db *pgxpool.Pool
}

func NewFileRepo(db *pgxpool.Pool) storage.FileRepoI {
	return &FileRepo{
		db: db,
	}
}

func (s FileRepo) Create(ctx context.Context, req *pb.AddFilesReq) (res *pb.AddFilesRes, err error) {
	res = &pb.AddFilesRes{}

	query := `INSERT INTO "file" (
		id,                            
    	url,
        type,
        draft_id
	) VALUES (
		$1,
		$2,
		$3,
		$4
	) RETURNING 
	    id,                            
    	url,
        type,
        draft_id`

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	err = s.db.QueryRow(ctx, query,
		id.String(),
		req.Url,
		req.Type,
		req.ArticleId,
	).Scan(
		&res.Id,
		&res.Url,
		&res.Type,
		&res.ArticleId,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s FileRepo) GetList(ctx context.Context, req *pb.GetFilesReq) (res *pb.GetFilesRes, err error) {
	res = &pb.GetFilesRes{}
	params := make(map[string]interface{})
	var arr []interface{}

	query := `SELECT
		id,                            
    	url,
        type,
        draft_id
	FROM
		"file"`
	filter := " WHERE 1=1"

	offset := " OFFSET 0"

	limit := " LIMIT 50"

	if req.Type != "" {
		params["type"] = req.Type
		filter += " AND type = :type"
	}

	if util.IsValidUUID(req.ArticleId) {
		params["draft_id"] = req.ArticleId
		filter += " AND draft_id = :draft_id"
	}

	cQ := `SELECT count(1) FROM "file"` + filter

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
		obj := &pb.File{}

		err = rows.Scan(
			&obj.Id,
			&obj.Url,
			&obj.Type,
			&obj.ArticleId,
		)
		if err != nil {
			return res, err
		}

		res.Files = append(res.Files, obj)
	}

	return res, nil
}

func (s FileRepo) Delete(ctx context.Context, req *pb.DeleteFilesReq) (rowsAffected int64, err error) {
	query := `DELETE FROM "subject" WHERE id = $1`

	ids := strings.Split(req.Ids, ",")

	for _, id := range ids {
		result, err := s.db.Exec(ctx, query, id)
		if err != nil {
			return 0, err
		}

		rowsAffected += result.RowsAffected()
	}

	return rowsAffected, err
}
