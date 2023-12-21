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

type NotifyRepo struct {
	db *pgxpool.Pool
}

func NewNotifyRepo(db *pgxpool.Pool) storage.NotifyRepoI {
	return &NotifyRepo{
		db: db,
	}
}

func (s *NotifyRepo) Create(ctx context.Context, req *pb.CreateNotificationReq) (res *pb.CreateNotificationRes, err error) {

	res = &pb.CreateNotificationRes{}

	query := `INSERT INTO "notification" (
		id,                            
    	subject,
        text,
        email,
        status
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5
	) RETURNING 
	    id, 
	    subject,
	    text,
	    email,
	    status,
	    TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at,
		TO_CHAR(updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at`

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	err = s.db.QueryRow(ctx, query,
		id.String(),
		req.GetSubject(),
		req.GetText(),
		req.GetEmail(),
		req.GetStatus(),
	).Scan(
		&res.Id,
		&res.Subject,
		&res.Text,
		&res.Email,
		&res.Status,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *NotifyRepo) Get(ctx context.Context, req *pb.GetNotificationReq) (res *pb.GetNotificationRes, err error) {
	res = &pb.GetNotificationRes{}

	query := `SELECT
		id, 
	    subject,
	    text,
	    email,
	    status,
	    TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at,
		TO_CHAR(updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at
	FROM
		"notification"
	WHERE
		id = $1`

	err = s.db.QueryRow(
		ctx,
		query,
		req.GetId()).Scan(
		&res.Id,
		&res.Subject,
		&res.Text,
		&res.Email,
		&res.Status,
		&res.CreatedAt,
		&res.UpdatedAt,
	)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *NotifyRepo) GetList(ctx context.Context, req *pb.GetNotificationListReq) (res *pb.GetNotificationListRes, err error) {
	res = &pb.GetNotificationListRes{}
	params := make(map[string]interface{})
	var arr []interface{}

	query := `SELECT
		id, 
	    subject,
	    text,
	    email,
	    status,
	    TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at,
		TO_CHAR(updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at
	FROM
		"notification"`
	filter := " WHERE 1=1"

	offset := " OFFSET 0"

	limit := " LIMIT 10"

	if len(req.Search) > 0 {
		params["search"] = req.Search
		filter += ` AND (subject ILIKE '%' || :search || '%')`
	}

	if req.Offset > 0 {
		params["offset"] = req.Offset
		offset = " OFFSET :offset"
	}

	if req.Limit > 0 {
		params["limit"] = req.Limit
		limit = " LIMIT :limit"
	}

	cQ := `SELECT count(1) FROM "notification"` + filter

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
		obj := &pb.Notification{}

		err = rows.Scan(
			&obj.Id,
			&obj.Subject,
			&obj.Text,
			&obj.Email,
			&obj.Status,
			&obj.CreatedAt,
			&obj.UpdatedAt,
		)
		if err != nil {
			return res, err
		}

		res.Notifications = append(res.Notifications, obj)
	}

	return res, nil
}

func (s *NotifyRepo) Update(ctx context.Context, req *pb.UpdateNotificationReq) (res *pb.UpdateNotificationRes, err error) {
	res = &pb.UpdateNotificationRes{}

	query := `UPDATE "notification" SET                
    	status = :status
	WHERE
		id = :id
	RETURNING 
		id, 
	    subject,
	    text,
	    email,
	    status,
	    TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at,
		TO_CHAR(updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at`

	params := map[string]interface{}{
		"id":     req.GetId(),
		"status": req.GetStatus(),
	}

	q, arr := helper.ReplaceQueryParams(query, params)
	err = s.db.QueryRow(ctx, q, arr...).Scan(
		&res.Id,
		&res.Subject,
		&res.Text,
		&res.Email,
		&res.Status,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func (s *NotifyRepo) Delete(ctx context.Context, req *pb.DeleteNotificationReq) (rowsAffected int64, err error) {
	query := `DELETE FROM "notification" WHERE id = $1`

	result, err := s.db.Exec(ctx, query, req.GetId())
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}
