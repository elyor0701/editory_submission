package notification

import (
	"context"
	"editory_submission/config"
	pb "editory_submission/genproto/notification_service"
	"editory_submission/pkg/helper"
	"editory_submission/storage"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type EmailTmpRepo struct {
	db *pgxpool.Pool
}

func NewEmailTmpRepo(db *pgxpool.Pool) storage.EmailTemplateRepoI {
	return &EmailTmpRepo{
		db: db,
	}
}

func (s *EmailTmpRepo) Create(ctx context.Context, req *pb.CreateEmailTmpReq) (res *pb.CreateEmailTmpRes, err error) {

	res = &pb.CreateEmailTmpRes{}

	query := `INSERT INTO "email_template" (
		id,                            
    	title,
        description,
        type,
        text
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5
	) RETURNING 
	    id, 
	    title,
	    description,
        type,
        text,
        TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at`

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	err = s.db.QueryRow(ctx, query,
		id.String(),
		req.GetTitle(),
		req.GetDescription(),
		req.GetType(),
		req.GetText(),
	).Scan(
		&res.Id,
		&res.Title,
		&res.Description,
		&res.Type,
		&res.Text,
		&res.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *EmailTmpRepo) Get(ctx context.Context, req *pb.GetEmailTmpReq) (res *pb.GetEmailTmpRes, err error) {
	res = &pb.GetEmailTmpRes{}

	query := `SELECT
		id, 
	    title,
	    description,
        type,
        text,
        TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at
	FROM
		"email_template"
	WHERE
		id = $1`

	err = s.db.QueryRow(
		ctx,
		query,
		req.GetId()).Scan(
		&res.Id,
		&res.Title,
		&res.Description,
		&res.Type,
		&res.Text,
		&res.CreatedAt,
	)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *EmailTmpRepo) GetList(ctx context.Context, req *pb.GetEmailTmpListReq) (res *pb.GetEmailTmpListRes, err error) {
	res = &pb.GetEmailTmpListRes{}
	params := make(map[string]interface{})
	var arr []interface{}
	validMailNoType := map[string]bool{
		config.REGISTRATION:          true,
		config.NEW_ARTICLE_TO_REVIEW: true,
		config.RESET_PASSWORD:        true,
		config.ACCOUNT_DEACTIVATION:  true,
		config.NEW_JOURNAL_USER:      true,
	}

	query := `SELECT
		id, 
	    title,
	    description,
        type,
        text,
        TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at
	FROM
		"email_template"`
	filter := " WHERE 1=1"

	offset := " OFFSET 0"

	limit := " LIMIT 10"

	if len(req.Search) > 0 {
		params["search"] = req.Search
		filter += ` AND (title ILIKE '%' || :search || '%')`
	}

	if _, ok := validMailNoType[req.GetType()]; ok {
		params["type"] = req.Type
		filter += ` AND type = :type`
	}

	if req.Offset > 0 {
		params["offset"] = req.Offset
		offset = " OFFSET :offset"
	}

	if req.Limit > 0 {
		params["limit"] = req.Limit
		limit = " LIMIT :limit"
	}

	cQ := `SELECT count(1) FROM "email_template"` + filter

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
		obj := &pb.EmailTmp{}

		err = rows.Scan(
			&obj.Id,
			&obj.Title,
			&obj.Description,
			&obj.Type,
			&obj.Text,
			&obj.CreatedAt,
		)
		if err != nil {
			return res, err
		}

		res.EmailTmps = append(res.EmailTmps, obj)
	}

	return res, nil
}

func (s *EmailTmpRepo) Update(ctx context.Context, req *pb.UpdateEmailTmpReq) (res *pb.UpdateEmailTmpRes, err error) {
	res = &pb.UpdateEmailTmpRes{}

	query := `UPDATE "email_template" SET                
    	title = :title,
    	description = :description,
    	text = :text
	WHERE
		id = :id
	RETURNING 
		id,
		title,
		description,
	    type,
		text,
		TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at`

	params := map[string]interface{}{
		"id":          req.GetId(),
		"title":       req.GetTitle(),
		"description": req.GetDescription(),
		"text":        req.GetText(),
	}

	q, arr := helper.ReplaceQueryParams(query, params)
	err = s.db.QueryRow(ctx, q, arr...).Scan(
		&res.Id,
		&res.Title,
		&res.Description,
		&res.Type,
		&res.Text,
		&res.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func (s *EmailTmpRepo) Delete(ctx context.Context, req *pb.DeleteEmailTmpReq) (rowsAffected int64, err error) {
	query := `DELETE FROM "email_template" WHERE id = $1`

	result, err := s.db.Exec(ctx, query, req.GetId())
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}
