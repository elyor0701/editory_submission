package auth

import (
	"context"
	"editory_submission/config"
	pb "editory_submission/genproto/auth_service"
	"editory_submission/pkg/helper"
	"editory_submission/pkg/util"
	"editory_submission/storage"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type sessionRepo struct {
	db *pgxpool.Pool
}

func NewSessionRepo(db *pgxpool.Pool) storage.SessionRepoI {
	return &sessionRepo{
		db: db,
	}
}

// log.Printf("--->STRG: CreateSessionRequest: %+v", entity)
// log.Printf("---STRG->DeleteExpiredUserSessions---> %s", userID)

func (r *sessionRepo) Create(ctx context.Context, in *pb.CreateSessionReq) (res *pb.Session, err error) {

	query := `INSERT INTO "session" (
			id,
			user_id,
            role_id,
			ip,
			data,
			expires_at
		) VALUES (
				$1,
				$2,
				$3,
				$4,
				$5,
		        $6
		)`

	random, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	_, err = r.db.Exec(ctx, query,
		random.String(),
		in.GetUserId(),
		util.NewNullString(in.GetRoleId()),
		util.NewNullString(in.GetIp()),
		util.NewNullString(in.GetData()),
		in.GetExpiresAt(),
	)
	if err != nil {
		return nil, err
	}

	res = &pb.Session{
		Id:        random.String(),
		UserId:    in.GetUserId(),
		RoleId:    in.GetRoleId(),
		Ip:        in.GetIp(),
		Data:      in.GetData(),
		ExpiresAt: in.GetExpiresAt(),
	}

	return res, err
}

func (r *sessionRepo) GetByPK(ctx context.Context, in *pb.SessionPrimaryKey) (res *pb.Session, err error) {

	res = &pb.Session{}

	query := `SELECT
		id,
		user_id,
		coalesce(role_id::varchar, ''),
		TEXT(ip) AS ip,
		data,
		COALESCE(TO_CHAR(expires_at, ` + config.DatabaseQueryTimeLayout + `)::TEXT, '') AS expires_at,
		COALESCE(TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `)::TEXT, '') AS created_at,
		COALESCE(TO_CHAR(updated_at, ` + config.DatabaseQueryTimeLayout + `)::TEXT, '') AS updated_at
	FROM
		"session"
	WHERE
		id = $1`

	err = r.db.QueryRow(ctx, query, in.GetId()).Scan(
		&res.Id,
		&res.UserId,
		&res.RoleId,
		&res.Ip,
		&res.Data,
		&res.ExpiresAt,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *sessionRepo) GetList(ctx context.Context, in *pb.SessionGetList) (res *pb.GetSessionListRes, err error) {
	res = &pb.GetSessionListRes{}
	params := make(map[string]interface{})
	var arr []interface{}
	query := `SELECT
		id,
		user_id,
		role_id,
		TEXT(ip) AS ip,
		data,
		TO_CHAR(expires_at, ` + config.DatabaseQueryTimeLayout + `) AS expires_at,
		TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at,
		TO_CHAR(updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at
	FROM
		"session"`
	filter := " WHERE 1=1"
	offset := " OFFSET 0"
	limit := " LIMIT 10"

	if len(in.Search) > 0 {
		params["search"] = in.Search
		filter += " AND ((ip) ILIKE ('%' || :search || '%'))"
	}

	if in.Offset > 0 {
		params["offset"] = in.Offset
		offset = " OFFSET :offset"
	}

	if in.Limit > 0 {
		params["limit"] = in.Limit
		limit = " LIMIT :limit"
	}

	cQ := `SELECT count(1) FROM "session"` + filter
	cQ, arr = helper.ReplaceQueryParams(cQ, params)
	err = r.db.QueryRow(ctx, cQ, arr...).Scan(
		&res.Count,
	)
	if err != nil {
		return res, err
	}

	q := query + filter + offset + limit
	q, arr = helper.ReplaceQueryParams(q, params)

	rows, err := r.db.Query(ctx, q, arr...)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		obj := &pb.Session{}
		err = rows.Scan(
			&obj.Id,
			&obj.UserId,
			&obj.RoleId,
			&obj.Ip,
			&obj.Data,
			&obj.ExpiresAt,
			&obj.CreatedAt,
			&obj.UpdatedAt,
		)

		if err != nil {
			return res, err
		}

		res.Sessions = append(res.Sessions, obj)
	}

	return res, nil
}

func (r *sessionRepo) Update(ctx context.Context, in *pb.UpdateSessionReq) (res *pb.Session, err error) {

	params := make(map[string]interface{})
	queryInitial := `UPDATE "session" SET
        ip = :ip,
        expires_at = :expires_at,
		updated_at = CURRENT_TIMESTAMP`

	filter := ` WHERE id = :id`
	params["ip"] = in.Ip
	params["expires_at"] = in.ExpiresAt
	params["id"] = in.Id

	if util.IsValidUUID(in.RoleId) {
		params["role_id"] = in.RoleId
		queryInitial += `, role_id = :role_id`
	}

	if in.Data != "" {
		params["data"] = in.Data
		queryInitial += `, data = :data`
	}

	query := queryInitial + filter

	cQuery, arr := helper.ReplaceQueryParams(query, params)

	_, err = r.db.Exec(ctx, cQuery, arr...)
	if err != nil {
		return nil, err
	}

	res = &pb.Session{
		Id:        in.GetId(),
		UserId:    in.GetUserId(),
		RoleId:    in.GetRoleId(),
		Ip:        in.GetIp(),
		Data:      in.GetData(),
		ExpiresAt: in.GetExpiresAt(),
	}

	return res, nil
}

func (r *sessionRepo) Delete(ctx context.Context, in *pb.SessionPrimaryKey) (rowsAffected int64, err error) {
	query := `DELETE FROM "session" WHERE id = $1`

	result, err := r.db.Exec(ctx, query, in.GetId())
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}

func (r *sessionRepo) DeleteExpiredUserSessions(ctx context.Context, userID string) (rowsAffected int64, err error) {

	query := `DELETE FROM "session" WHERE user_id = $1 AND expires_at < CURRENT_TIMESTAMP`

	result, err := r.db.Exec(ctx, query, userID)
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}

func (r *sessionRepo) GetSessionListByUserID(ctx context.Context, userID string) (res *pb.GetSessionListRes, err error) {
	res = &pb.GetSessionListRes{}

	query := `SELECT
		id,
		user_id,
		coalesce(role_id::text, ''),
		TEXT(ip) AS ip,
		data,
		TO_CHAR(expires_at, ` + config.DatabaseQueryTimeLayout + `) AS expires_at,
		TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at,
		TO_CHAR(updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at
	FROM
		"session"
	WHERE user_id = $1`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		obj := &pb.Session{}
		err = rows.Scan(
			&obj.Id,
			&obj.UserId,
			&obj.RoleId,
			&obj.Ip,
			&obj.Data,
			&obj.ExpiresAt,
			&obj.CreatedAt,
			&obj.UpdatedAt,
		)

		if err != nil {
			return res, err
		}

		res.Sessions = append(res.Sessions, obj)
	}

	return res, nil
}
