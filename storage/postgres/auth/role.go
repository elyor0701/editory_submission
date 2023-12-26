package auth

import (
	"context"
	pb "editory_submission/genproto/auth_service"
	"editory_submission/pkg/helper"
	"editory_submission/pkg/util"
	"editory_submission/storage"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type RoleRepo struct {
	db *pgxpool.Pool
}

func NewRoleRepo(db *pgxpool.Pool) storage.RoleRepoI {
	return &RoleRepo{
		db: db,
	}
}

func (s *RoleRepo) Create(ctx context.Context, req *pb.Role) (res *pb.Role, err error) {
	fmt.Println("req:::::", req)
	query := `INSERT INTO "role" (
		id,                 
    	user_id,
        role_type,
        journal_id
	) VALUES (
		$1,
		$2,
		$3,
		$4
	)`

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	_, err = s.db.Exec(ctx, query,
		id.String(),
		req.GetUserId(),
		req.GetRoleType(),
		util.NewNullString(req.GetJournalId()),
	)

	res = &pb.Role{
		Id:        id.String(),
		UserId:    req.GetUserId(),
		RoleType:  req.GetRoleType(),
		JournalId: req.GetJournalId(),
	}

	return res, err
}

func (s *RoleRepo) Get(ctx context.Context, req *pb.GetRoleReq) (res *pb.Role, err error) {
	res = &pb.Role{}

	query := `SELECT
		id,                 
    	user_id,
    	role_type,
    	coalesce(journal_id::VARCHAR, '')
	FROM
		"role"
	WHERE
		id = $1`

	err = s.db.QueryRow(
		ctx,
		query,
		req.GetId()).Scan(
		&res.Id,
		&res.UserId,
		&res.RoleType,
		&res.JournalId,
	)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *RoleRepo) GetList(ctx context.Context, req *pb.GetRoleListReq) (res *pb.GetRoleListRes, err error) {
	res = &pb.GetRoleListRes{}
	params := make(map[string]interface{})
	var arr []interface{}
	fmt.Println("req", req)

	query := `SELECT
		id,                 
    	user_id,
    	role_type,
    	coalesce(journal_id::VARCHAR, '')
	FROM
		"role"`
	filter := " WHERE 1=1"

	offset := " OFFSET 0"

	limit := " LIMIT 10"

	if len(req.Search) > 0 {
		params["search"] = req.Search
		filter += ` AND (username ILIKE '%' || :search || '%')`
	}

	if util.IsValidUUID(req.GetUserId()) {
		params["user_id"] = req.GetUserId()
		filter += ` AND user_id = :user_id`
	}

	if util.IsValidUUID(req.GetJournalId()) {
		params["journal_id"] = req.GetJournalId()
		filter += ` AND journal_id = :journal_id`
	}

	if len(req.GetRoleTypes()) != 0 {
		params["role_types"] = req.GetRoleTypes()
		filter += ` AND role_type = ANY(:role_types)`
	}

	if req.Offset > 0 {
		params["offset"] = req.Offset
		offset = " OFFSET :offset"
	}

	if req.Limit > 0 {
		params["limit"] = req.Limit
		limit = " LIMIT :limit"
	}

	cQ := `SELECT count(1) FROM "role"` + filter

	cQ, arr = helper.ReplaceQueryParams(cQ, params)

	fmt.Println("cQ", cQ)

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
		obj := &pb.Role{}

		err = rows.Scan(
			&obj.Id,
			&obj.UserId,
			&obj.RoleType,
			&obj.JournalId,
		)
		if err != nil {
			return res, err
		}

		res.Roles = append(res.Roles, obj)
	}

	return res, nil
}
func (s *RoleRepo) Update(ctx context.Context, req *pb.Role) (rowsAffected int64, err error) {
	query := `UPDATE "role" SET                
    	role_type = :role_type
	WHERE
		id = :id`

	params := map[string]interface{}{
		"id":        req.GetId(),
		"role_type": req.GetRoleType(),
	}

	q, arr := helper.ReplaceQueryParams(query, params)
	result, err := s.db.Exec(ctx, q, arr...)
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}
func (s *RoleRepo) Delete(ctx context.Context, req *pb.DeleteRoleReq) (rowsAffected int64, err error) {
	query := `DELETE FROM "role" WHERE id = $1`

	result, err := s.db.Exec(ctx, query, req.GetId())
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}
